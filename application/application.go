// Пакет application - содержит код для взаимодействия с внешним миром по HTTP

package application

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Nikolay200669/Ports_and_Adapters_patern/domain"
	"github.com/Nikolay200669/Ports_and_Adapters_patern/infrastructure"
	"github.com/gin-gonic/gin"
)

// App представляет приложение с веб-сервером Gin и службой UserService
type App struct {
	UserService *domain.UserService
}

// NewApp создает новый экземпляр приложения
func NewApp(db *sql.DB) *App {
	// Инициализация UserRepository
	userRepository := infrastructure.NewSQLUserRepository(db)

	// Инициализация UserService
	userService := domain.NewUserService(userRepository)

	return &App{
		UserService: userService,
	}
}

// SetupRouter настраивает маршруты Gin
func (app *App) SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/users/:id", app.getUserHandler)

	return router
}

// getUserHandler обрабатывает запрос GET /users/:id
func (app *App) getUserHandler(c *gin.Context) {
	userID := c.Param("id")
	id := 0

	if _, err := fmt.Sscan(userID, &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := app.UserService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	c.JSON(http.StatusOK, user)
}
