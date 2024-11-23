package repositories

import (
	"github.com/saipulmuiz/mnc-test-tahap2/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo interface {
	RegisterUser(user *models.User) (*models.User, error)
	FindById(userId int) (*models.User, error)
	CheckUserByPhoneNumber(phoneNumber string) (*models.User, error)
	CheckUserByID(id int, user *models.User) (*models.User, error)
	UpdateUser(userId string, userUpdate *models.User) (*models.User, error)
}

type userRepo struct {
	db         *gorm.DB
	globalRepo GlobalRepo
}

func NewUserRepo(db *gorm.DB, globalRepo GlobalRepo) UserRepo {
	return &userRepo{db, globalRepo}
}

func (u *userRepo) RegisterUser(user *models.User) (*models.User, error) {
	return user, u.db.Create(&user).Error
}

func (u *userRepo) FindById(userId int) (*models.User, error) {
	var user models.User
	err := u.db.Where("id = ?", userId).First(&user).Error
	return &user, err
}

func (u *userRepo) CheckUserByPhoneNumber(phoneNumber string) (*models.User, error) {
	var user *models.User
	return user, u.db.Where("phone_number=?", phoneNumber).Take(&user).Error
}

func (u *userRepo) CheckUserByID(id int, user *models.User) (*models.User, error) {
	return user, u.db.Where("id", id).Take(&user).Error
}

func (u *userRepo) UpdateUser(userId string, userUpdate *models.User) (*models.User, error) {
	var user models.User
	result := u.db.Model(&user).Clauses(clause.Returning{}).Where("user_id", userId).Updates(userUpdate)
	return &user, result.Error
}
