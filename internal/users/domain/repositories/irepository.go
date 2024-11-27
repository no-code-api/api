package repositories

import "github.com/no-code-api/no-code-api/internal/users/domain/models"

type IRepository interface {
	Create(user *models.User) (ok bool)
	FindAll() (users []*models.User, ok bool)
	FindById(id uint) (user *models.User, ok bool)
	FindByEmail(email string) (user *models.User, ok bool)
	Update(user *models.User) (ok bool)
	Delete(id uint) (ok bool)
}
