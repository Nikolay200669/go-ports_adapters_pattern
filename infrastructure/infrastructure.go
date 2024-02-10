// Пакет infrastructure содержит адаптеры для взаимодействия с внешним миром.
// В данном случае, это адаптер для работы с базой данных.
// Он реализует порт UserRepository, предоставляя методы для работы с пользователями в базе данных.

package infrastructure

import (
	"database/sql"

	"github.com/Nikolay200669/Ports_and_Adapters_patern/domain"
)

// SQLUserRepository реализует порт UserRepository для работы с базой данных
type SQLUserRepository struct {
	db *sql.DB
}

// NewSQLUserRepository создает новый экземпляр SQLUserRepository
func NewSQLUserRepository(db *sql.DB) *SQLUserRepository {
	return &SQLUserRepository{
		db: db,
	}
}

// FindByID ищет пользователя по его идентификатору в базе данных
func (r *SQLUserRepository) FindByID(id int) (*domain.User, error) {
	// Реализация запроса к базе данных
	// ...
	return &domain.User{}, nil
}

// Save сохраняет пользователя в базе данных
func (r *SQLUserRepository) Save(user *domain.User) error {
	// Реализация сохранения в базе данных
	// ...
	return nil
}
