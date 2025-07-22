package main

import (
	"log"
	"time"

	"github.com/cesarlead/practica_go_back_gin_basico/internal/config"
	"github.com/cesarlead/practica_go_back_gin_basico/internal/database"
	"github.com/cesarlead/practica_go_back_gin_basico/internal/router"
	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("⚠️  No se pudo cargar .env, usando variables del entorno")
	}

	cfg := config.Load()

	pool, err := database.NewPool(cfg)
	if err != nil {
		log.Fatalf("⛔ no se pudo inicializar pool de BD: %v", err)
	}
	defer pool.Close()

	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()

	engine.SetTrustedProxies([]string{"127.0.0.1"})

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.CORSOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	engine.Use(gin.Logger(), gin.Recovery())

	// Rutas de negocio
	router.Setup(engine, pool)

	// Swagger UI
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := engine.Run(":" + cfg.Port); err != nil {
		log.Fatalf("⛔ no se pudo levantar el servidor: %v", err)
	}
}
