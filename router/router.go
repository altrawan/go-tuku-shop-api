package router

import (
	"os"

	"go-tuku-shop-api/controller"
	"go-tuku-shop-api/docs"
	"go-tuku-shop-api/middleware"
	"go-tuku-shop-api/repository"
	"go-tuku-shop-api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @contact.name 		API Support
// @contact.url 		http://www.swagger.io/support
// @contact.email 	support@swagger.io

// @license.name 		Apache 2.0
// @license.url 		http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService	http://swagger.io/terms/
func NewRouter(db *gorm.DB) *gin.Engine {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Tuku Shop API"
	docs.SwaggerInfo.Description = "RESTful API for E-commerce platforms"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	addressRepository := repository.NewAddressRepository(db)
	authRepository := repository.NewAuthRepository(db)
	brandRepository := repository.NewBrandRepository(db)
	cartRepository := repository.NewCartRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)
	productRepository := repository.NewProductRepository(db)
	profileRepository := repository.NewProfileRepository(db)
	storeRepository := repository.NewStoreRepository(db)
	transactionRepository := repository.NewTransactionRepository(db)

	addressService := service.NewAddressService(addressRepository)
	authService := service.NewAuthService(authRepository)
	brandService := service.NewBrandService(brandRepository)
	cartService := service.NewCartService(cartRepository)
	categoryService := service.NewCategoryService(categoryRepository)
	productService := service.NewProductService(productRepository)
	profileService := service.NewProfileService(profileRepository)
	storeService := service.NewStoreService(storeRepository)
	transactionService := service.NewTransactionService(transactionRepository)

	addressController := controller.NewAddressController(addressService)
	authController := controller.NewAuthController(authService)
	brandController := controller.NewBrandController(brandService)
	cartController := controller.NewCartController(cartService)
	categoryController := controller.NewCategoryController(categoryService)
	productController := controller.NewProductController(productService)
	profileController := controller.NewProfileController(profileService)
	storeController := controller.NewStoreController(storeService)
	transactionController := controller.NewTransactionController(transactionService)

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"message": "Path not found"})
	})

	// base path /api/v1
	v1 := r.Group("/api/v1")

	// use ginSwagger middleware to serve the API docs
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	addressRoutes := v1.Group("/address")
	{
		addressRoutes.Use(middleware.JwtAuth()).GET("/", addressController.List)
		addressRoutes.Use(middleware.JwtAuth()).GET("/:id", addressController.Detail)
		addressRoutes.Use(middleware.JwtAuth()).POST("/", addressController.Store)
		addressRoutes.Use(middleware.JwtAuth()).PUT("/:id", addressController.Update)
		addressRoutes.Use(middleware.JwtAuth()).DELETE("/:id", addressController.Delete)
	}

	authRoutes := v1.Group("/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register-seller", authController.RegisterSeller)
		authRoutes.POST("/register-buyer", authController.RegisterBuyer)
		// authRoutes.POST("/activation/:token", authController.Activation)
		// authRoutes.POST("/forgot-password", authController.ForgotPassword)
		// authRoutes.POST("/reset-password/:token", authController.ResetPassword)
	}

	brandRoutes := v1.Group("/brand")
	{
		brandRoutes.GET("/", brandController.List)
		brandRoutes.Use(middleware.JwtAuth()).POST("/", brandController.Store)
		brandRoutes.Use(middleware.JwtAuth()).PUT("/:id", brandController.Update)
		brandRoutes.Use(middleware.JwtAuth()).DELETE("/:id", brandController.Delete)
	}

	cartRoutes := v1.Group("/cart")
	{
		cartRoutes.Use(middleware.JwtAuth()).GET("/", cartController.List)
		cartRoutes.Use(middleware.JwtAuth()).GET("/:id", cartController.Detail)
		cartRoutes.Use(middleware.JwtAuth()).POST("/", cartController.Store)
		cartRoutes.Use(middleware.JwtAuth()).PUT("/:id", cartController.Update)
		cartRoutes.Use(middleware.JwtAuth()).DELETE("/:id", cartController.Delete)
	}

	categoryRoutes := v1.Group("/category")
	{
		categoryRoutes.GET("/", categoryController.List)
		categoryRoutes.Use(middleware.JwtAuth()).POST("/", categoryController.Store)
		categoryRoutes.Use(middleware.JwtAuth()).PUT("/:id", categoryController.Update)
		categoryRoutes.Use(middleware.JwtAuth()).DELETE("/:id", categoryController.Delete)
	}

	productRoutes := v1.Group("/product")
	{
		productRoutes.GET("/", productController.List)
		productRoutes.GET("/:id", productController.Detail)
		productRoutes.Use(middleware.JwtAuth()).POST("/", productController.Store)
		productRoutes.Use(middleware.JwtAuth()).PUT("/:id", productController.Update)
		productRoutes.Use(middleware.JwtAuth()).DELETE("/:id", productController.Delete)
	}

	profileRoutes := v1.Group("/profile")
	{
		profileRoutes.Use(middleware.JwtAuth()).GET("/", profileController.List)
		profileRoutes.Use(middleware.JwtAuth()).GET("/:id", profileController.Detail)
		profileRoutes.Use(middleware.JwtAuth()).PUT("/:id", profileController.Update)
		profileRoutes.Use(middleware.JwtAuth()).PUT("/change-password", profileController.ChangePassword)
	}

	storeRoutes := v1.Group("/store")
	{
		storeRoutes.Use(middleware.JwtAuth()).GET("/", storeController.List)
		storeRoutes.Use(middleware.JwtAuth()).GET("/:id", storeController.Detail)
		storeRoutes.Use(middleware.JwtAuth()).PUT("/:id", storeController.Update)
		storeRoutes.Use(middleware.JwtAuth()).PUT("/change-password", storeController.ChangePassword)
	}

	transactionRoutes := v1.Group("/transaction")
	{
		transactionRoutes.Use(middleware.JwtAuth()).GET("/", transactionController.List)
		transactionRoutes.Use(middleware.JwtAuth()).GET("/:id", transactionController.Detail)
		transactionRoutes.Use(middleware.JwtAuth()).POST("/", transactionController.Store)
		transactionRoutes.Use(middleware.JwtAuth()).PUT("/:id/address", transactionController.UpdateAddress)
		transactionRoutes.Use(middleware.JwtAuth()).PUT("/:id/payment", transactionController.UpdatePayment)
		transactionRoutes.Use(middleware.JwtAuth()).DELETE("/:id", transactionController.Delete)
	}

	r.Run(":" + port)

	return r
}
