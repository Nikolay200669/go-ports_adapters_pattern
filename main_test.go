// Пакет application_test - тесты для кода, взаимодействующего с внешним миром по HTTP

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Nikolay200669/Ports_and_Adapters_patern/application"
	"github.com/Nikolay200669/Ports_and_Adapters_patern/domain"
	"github.com/stretchr/testify/assert"
)

// MockUserRepository реализует UserRepository для тестирования
type MockUserRepository struct {
	UserToReturn *domain.User
	ErrToReturn  error
}

// FindByID имитирует поиск пользователя по идентификатору
func (m *MockUserRepository) FindByID(id int) (*domain.User, error) {
	return m.UserToReturn, m.ErrToReturn
}

// Save имитирует сохранение пользователя
func (m *MockUserRepository) Save(user *domain.User) error {
	return m.ErrToReturn
}

// TestApp_GetUserHandler тестирует метод getUserHandler в App
func TestApp_GetUserHandler(t *testing.T) {
	// Создание мока UserRepository для тестирования
	mockRepo := &MockUserRepository{
		UserToReturn: &domain.User{
			ID:   1,
			Name: "John Doe",
		},
		ErrToReturn: nil,
	}

	// Создание мока UserService для тестирования
	mockUserService := &domain.UserService{
		UserRepository: mockRepo,
	}

	// Создание экземпляра приложения с моками
	app := &application.App{
		UserService: mockUserService,
	}

	// Настройка роутера Gin
	router := app.SetupRouter()

	// Создание фейкового HTTP-запроса
	req, err := http.NewRequest("GET", "/users/1", nil)
	assert.NoError(t, err)

	// Создание фейкового HTTP-ответа
	w := httptest.NewRecorder()

	// Запуск роутера Gin
	router.ServeHTTP(w, req)

	// Проверка статуса ответа
	assert.Equal(t, http.StatusOK, w.Code)

	// Проверка тела ответа
	expectedResponse := `{"ID":1,"Name":"John Doe"}`
	assert.Equal(t, expectedResponse, w.Body.String())
}

// TestApp_GetUserHandler_InvalidID тестирует метод getUserHandler в App с недопустимым ID
func TestApp_GetUserHandler_InvalidID(t *testing.T) {
	// Создание экземпляра приложения без моков (используется реальный UserService)
	app := &application.App{}

	// Настройка роутера Gin
	router := app.SetupRouter()

	// Создание фейкового HTTP-запроса с недопустимым ID
	req, err := http.NewRequest("GET", "/users/invalid", nil)
	assert.NoError(t, err)

	// Создание фейкового HTTP-ответа
	w := httptest.NewRecorder()

	// Запуск роутера Gin
	router.ServeHTTP(w, req)

	// Проверка статуса ответа (должен быть BadRequest)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
