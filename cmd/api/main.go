package main

import (
	_ "be-menu-tree-system/docs"
	"be-menu-tree-system/internal/api/handler"
	"be-menu-tree-system/internal/repository"
	"be-menu-tree-system/internal/service"
	"be-menu-tree-system/pkg/database"
	"be-menu-tree-system/pkg/logger"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Menu Tree Management API
// @version 1.0
// @description Production-ready hierarchical menu management system.
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @host localhost:8080
// @BasePath /
func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Init logger
	logger.Init()
	sugar := logger.Get().Sugar()
	defer sugar.Sync()

	// Connect to DB
	db, err := database.InitDB()
	if err != nil {
		sugar.Fatalf("Failed to connect to database: %v", err)
	}

	// Setup components
	menuRepo := repository.NewMenuRepository(db)
	menuService := service.NewMenuService(menuRepo)
	menuHandler := handler.NewMenuHandler(menuService)

	// Setup Gin
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	// CORS Middleware (Basic)
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Routes
	api := r.Group("/api")
	{
		menus := api.Group("/menus")
		{
			menus.POST("", menuHandler.CreateMenu)
			menus.GET("", menuHandler.GetMenuTree)
			menus.GET("/:id", menuHandler.GetMenuByID)
			menus.PUT("/:id", menuHandler.UpdateMenu)
			menus.DELETE("/:id", menuHandler.DeleteMenu)
			menus.PATCH("/:id/move", menuHandler.MoveMenu)
			menus.PATCH("/:id/reorder", menuHandler.ReorderMenu)
		}
	}

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	sugar.Infof("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		sugar.Fatalf("Failed to start server: %v", err)
	}
}
