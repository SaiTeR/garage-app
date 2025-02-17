package repositories

import (
	"garage-app-client-backend/app/models"
	"garage-app-client-backend/app/requests"
	"garage-app-client-backend/database"
	"gorm.io/gorm"
)

// UserRepository — это структура репозитория для работы с пользователями
type UserRepository struct {
	DB *gorm.DB
}

// NewUserRepository — конструктор для создания нового репозитория
func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: database.GetMySQLClient(),
	}
}

// CreateUser — метод для создания нового пользователя в базе данных
func (repo *UserRepository) CreateUser(request *requests.RegisterRequest) (uint, error) {
	// Создаем нового пользователя, используя данные из запроса
	user := models.User{
		Phone:    request.Phone,
		Name:     request.Name,
		Surname:  request.Surname,
		Password: request.Password,
		Email:    request.Email,
	}

	// Сохраняем пользователя в базе данных
	if err := repo.DB.Create(&user).Error; err != nil {
		return 0, err // Возвращаем ошибку, если пользователь не был создан
	}

	// Возвращаем ID созданного пользователя
	return user.ID, nil
}
