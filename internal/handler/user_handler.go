package handler

import (
	"net/http"
	"strconv"

	"github.com/cesarlead/practica_go_back_gin_basico/internal/domain"
	"github.com/cesarlead/practica_go_back_gin_basico/internal/usecase"
	"github.com/gin-gonic/gin"
)

// UserHandler adapta HTTP ↔ UseCase
type UserHandler struct {
	uc usecase.UserUseCase
}

// NewUserHandler inyecta UserUseCase
func NewUserHandler(uc usecase.UserUseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

// CreateUser request body
type createUserRequest struct {
	Name  string `json:"name" binding:"required,min=5"`
	Email string `json:"email" binding:"required,email"`
}

// @Summary Crear usuario
// @Description Crea un nuevo usuario con nombre y email
// @Tags users
// @Accept json
// @Produce json
// @Param body body createUserRequest true "Datos de usuario"
// @Success 201 {object} domain.User
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /api/v1/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := h.uc.CreateUser(req.Name, req.Email)
	if err != nil {
		switch err {
		case domain.ErrInvalidName, domain.ErrInvalidEmail:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error al crear usuario"})
		}
		return
	}
	c.JSON(http.StatusCreated, u)
}

// @Summary Listar usuarios
// @Description Obtiene todos los usuarios
// @Tags users
// @Produce json
// @Success 200 {array} domain.User
// @Router /api/v1/users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.uc.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error al obtener usuarios"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// @Summary Obtener usuario
// @Description Obtiene un usuario por ID
// @Tags users
// @Produce json
// @Param id path int true "ID de usuario"
// @Success 200 {object} domain.User
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /api/v1/users/{id} [get]
func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}
	u, err := h.uc.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": u})
}

// @Summary Actualizar usuario
// @Description Modifica nombre y/o email de un usuario existente
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID de usuario"
// @Param body body createUserRequest true "Datos de actualización"
// @Success 200 {object} domain.User
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /api/v1/users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := h.uc.UpdateUser(id, req.Name, req.Email)
	if err != nil {
		switch err {
		case domain.ErrInvalidName, domain.ErrInvalidEmail:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		case domain.ErrUserNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error al actualizar usuario"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": u})

}

// @Summary Eliminar usuario
// @Description Borra un usuario por ID
// @Tags users
// @Produce json
// @Param id path int true "ID de usuario"
// @Success 204
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /api/v1/users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
		return
	}
	if err := h.uc.DeleteUser(id); err != nil {
		if err == domain.ErrUserNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error al eliminar usuario"})
		}
		return
	}
	c.Status(http.StatusNoContent)
}
