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

	// Product
	productRepo := repositories.NewProductRepo(db, globalRepo)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	// Cart
	cartRepo := repositories.NewCartRepo(db, globalRepo)
	cartService := services.NewCartService(cartRepo, productRepo)
	cartController := controllers.NewCartController(cartService)

	// Order
	orderRepo := repositories.NewOrderRepo(db, globalRepo)
	orderService := services.NewOrderService(productRepo, cartRepo, orderRepo, db)
	orderController := controllers.NewOrderController(orderService)

	//Route Group
	mainRouter := route.Group("/v1")
	{
		mainRouter.POST("/register", userController.RegisterUser)
		mainRouter.POST("/login", userController.Login)

		authorized := mainRouter.Group("/")
		authorized.Use(middlewares.Auth())
		{
			// user router
			authorized.POST("/logout", userController.Logout)
			authorized.PUT("/me/update", userController.UpdateProfile)

			// product router
			authorized.GET("/products", productController.GetProducts)
			authorized.GET("/products/:productId", productController.GetProductById)
			authorized.POST("/products", productController.CreateProduct)
			authorized.PUT("/products/:productId", productController.UpdateProduct)
			authorized.DELETE("/products/:productId", productController.DeleteProduct)

			// cart router
			authorized.GET("/carts", cartController.GetCarts)
			authorized.POST("/carts", cartController.AddToCart)
			authorized.PUT("/carts/:cartId", cartController.UpdateCart)
			authorized.DELETE("/carts/:cartId", cartController.DeleteCart)

			// order router
			authorized.GET("/orders", orderController.GetOrders)
			authorized.GET("/orders/:orderId", orderController.GetOrderById)
			authorized.POST("/orders/checkout", orderController.CheckoutOrder)
		}
	}

	return route
}