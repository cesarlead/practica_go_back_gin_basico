package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/cesarlead/practica_go_back_gin_basico/internal/handler"
	"github.com/cesarlead/practica_go_back_gin_basico/internal/repository"
	"github.com/cesarlead/practica_go_back_gin_basico/internal/usecase"
)

// Setup monta el API v1 en el engine, usando el pool de conexiones.
func Setup(r *gin.Engine, pool *pgxpool.Pool) {
	userRepo := repository.NewPostgresUserRepo(pool)
	userUC := usecase.NewUserUseCase(userRepo)
	userH := handler.NewUserHandler(userUC)

	api := r.Group("/api/v1")
	users := api.Group("/users")
	{
		users.POST("", userH.CreateUser)
		users.GET("", userH.GetAllUsers)
		users.GET("/:id", userH.GetUserByID)
		users.PUT("/:id", userH.UpdateUser)
		users.DELETE("/:id", userH.DeleteUser)
	}
}
