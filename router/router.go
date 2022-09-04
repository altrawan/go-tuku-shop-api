package router

import (
	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/controller"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/docs"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/middleware"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/repository"
	"gitlab.com/altrawan/final-project-bds-sanbercode-golang-batch-37/service"
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

	// addressRepository := repository.NewAddressRepository(db)
	authRepository := repository.NewAuthRepository(db)
	brandRepository := repository.NewBrandRepository(db)
	// cartRepository := repository.NewCartRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)
	productRepository := repository.NewProductRepository(db)
	profileRepository := repository.NewProfileRepository(db)
	storeRepository := repository.NewStoreRepository(db)

	// addressService := service.NewAddressService(addressRepository)
	authService := service.NewAuthService(authRepository)
	brandService := service.NewBrandService(brandRepository)
	// cartService := service.NewCartService(cartRepository)
	categoryService := service.NewCategoryService(categoryRepository)
	productService := service.NewProductService(productRepository)
	profileService := service.NewProfileService(profileRepository)
	storeService := service.NewStoreService(storeRepository)

	// addressController := controller.NewAddressController(addressService)
	authController := controller.NewAuthController(authService)
	brandController := controller.NewBrandController(brandService)
	// cartController := controller.NewCartController(cartService)
	categoryController := controller.NewCategoryController(categoryService)
	productController := controller.NewProductController(productService)
	profileController := controller.NewProfileController(profileService)
	storeController := controller.NewStoreController(storeService)

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"message": "Path not found"})
	})

	// base path /api/v1
	v1 := r.Group("/api/v1")

	// use ginSwagger middleware to serve the API docs
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// addressRoutes := v1.Group("/address")
	// {
	// 	addressRoutes.GET("/address", addressController.List)
	// 	addressRoutes.GET("/address/:id", addressController.Detail)
	// 	addressRoutes.POST("/address", addressController.Store)
	// 	addressRoutes.PUT("/address/:id", addressController.Update)
	// 	addressRoutes.DELETE("/address/:id", addressController.Delete)
	// }

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

	// cartRoutes := v1.Group("/cart")
	// {
	// 	cartRoutes.GET("/address", cartController.List)
	// 	cartRoutes.GET("/address/:id", cartController.Detail)
	// 	cartRoutes.POST("/address", cartController.Store)
	// 	cartRoutes.PUT("/address/:id", cartController.Update)
	// 	cartRoutes.DELETE("/address/:id", cartController.Delete)
	// }

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
		// productRoutes.GET("/:id", productController.Detail)
		// productRoutes.Use(middleware.JwtAuth()).POST("/", productController.Store)
		// productRoutes.Use(middleware.JwtAuth()).PUT("/:id", productController.Update)
		// productRoutes.Use(middleware.JwtAuth()).DELETE("/:id", productController.Delete)
	}

	profileRoutes := v1.Group("/profile")
	{
		// profileRoutes.GET("/", profileController.List)
		// profileRoutes.GET("/:id", profileController.Detail)
		profileRoutes.PUT("/:id", profileController.Update)
		// profileRoutes.DELETE("/change-password", profileController.ChangePassword)
	}

	storeRoutes := v1.Group("/store")
	{
		// storeRoutes.GET("/", storeController.List)
		// storeRoutes.GET("/:id", storeController.Detail)
		storeRoutes.PUT("/:id", storeController.Update)
		// storeRoutes.DELETE("/change-password", storeController.ChangePassword)
	}

	r.Run(":" + port)

	return r
}
