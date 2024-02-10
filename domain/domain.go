// Пакет domain - содержит бизнес-логику и порты для взаимодействия с внешним миром.

package domain

// UserRepository определяет порт для работы с хранилищем пользователей
type UserRepository interface {
	FindByID(id int) (*User, error)
	Save(user *User) error
}

// User представляет модель пользователя
type User struct {
	ID   int
	Name string
	// Другие поля
}

// UserService представляет службу, которая использует порт UserRepository
type UserService struct {
	UserRepository UserRepository
}

// NewUserService создает новый экземпляр UserService
func NewUserService(userRepository UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

// GetUserByID получает пользователя по его идентификатору, используя порт UserRepository
func (s *UserService) GetUserByID(userID int) (*User, error) {
	return s.UserRepository.FindByID(userID)
}
