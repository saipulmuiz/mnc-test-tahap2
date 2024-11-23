package routers

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/saipulmuiz/mnc-test-tahap2/repositories"
	log "github.com/sirupsen/logrus"

	"github.com/gin-contrib/cors"
	limits "github.com/gin-contrib/size"
	"github.com/saipulmuiz/mnc-test-tahap2/controllers"
	"github.com/saipulmuiz/mnc-test-tahap2/middlewares"
	"github.com/saipulmuiz/mnc-test-tahap2/services"
	"gorm.io/gorm"
)

func RouterConfig(db *gorm.DB) *gin.Engine {
	var maxSize int64 = 1024 * 1024 * 3 //3 MB
	route := gin.Default()
	logger := log.New()

	gin.SetMode(gin.DebugMode)

	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	//Enable CORS
	corsconfig := cors.DefaultConfig()
	corsconfig.AllowAllOrigins = true
	corsconfig.AddAllowHeaders("Authorization")
	route.Use(cors.New(corsconfig))
	route.Use(limits.RequestSizeLimiter(maxSize))
	route.Use(middlewares.ErrorHandler(logger))

	// Global
	globalRepo := repositories.NewGlobalRepo()

	// User
	userRepo := repositories.NewUserRepo(db, globalRepo)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// Transaction
	transactionRepo := repositories.NewTransactionRepo(db, globalRepo)
	topupRepo := repositories.NewTopupRepo(db, globalRepo)
	paymentRepo := repositories.NewPaymentRepo(db, globalRepo)
	transferRepo := repositories.NewTransferRepo(db, globalRepo)
	transactionService := services.NewTransactionService(transactionRepo, topupRepo, paymentRepo, transferRepo, userRepo, db)
	transactionController := controllers.NewTransactionController(transactionService)

	//Route Group
	mainRouter := route.Group("/v1")
	{
		mainRouter.POST("/register", userController.RegisterUser)
		mainRouter.POST("/login", userController.Login)
		mainRouter.POST("/refresh-token", userController.RefreshToken)

		authorized := mainRouter.Group("/")
		authorized.Use(middlewares.Auth())
		{
			// user router
			authorized.PUT("/profile", userController.UpdateProfile)

			// transaction router
			authorized.POST("/topup", transactionController.Topup)
			authorized.POST("/pay", transactionController.Payment)
		}
	}

	return route
}
