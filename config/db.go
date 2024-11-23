package config

import (
	"fmt"
	"os"

	"github.com/saipulmuiz/mnc-test-tahap2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname =%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		panic(err)
	}

	err = db.Exec("DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_type') THEN CREATE TYPE transaction_type AS ENUM ('credit', 'debit'); END IF; END $$;").Error
	if err != nil {
		panic(err)
	}

	err = db.Exec("DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_status') THEN CREATE TYPE transaction_status AS ENUM ('pending', 'success', 'failed'); END IF; END $$;").Error
	if err != nil {
		panic(err)
	}

	err = db.Exec("DO $$ BEGIN IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transfer_status') THEN CREATE TYPE transfer_status AS ENUM ('pending', 'success', 'failed'); END IF; END $$;").Error
	if err != nil {
		panic(err)
	}

	err = db.Debug().AutoMigrate(
		models.User{},
		models.Transaction{},
		models.Transfer{},
		models.TopUp{},
		models.Payment{},
	)

	// Seed data for `users` table
	// if err == nil && db.Migrator().HasTable(&models.User{}) {
	// 	if err := db.First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
	// 		users := []models.User{
	// 			{Name: "Agus", Email: "agus@gmail.com", Password: "agus123"},
	// 			{Name: "Fahri", Email: "fahri@gmail.com", Password: "fahri123"},
	// 			{Name: "Ujang", Email: "ujang@gmail.com", Password: "ujang123"},
	// 		}
	// 		if err := db.Create(&users).Error; err != nil {
	// 			log.Printf("Error seeding users: %s", err)
	// 		} else {
	// 			log.Println("Users seeded successfully")
	// 		}
	// 	}
	// }

	return db
}
