package repositories

import "github.com/no-code-api/api/internal/projects/domain/models"

type IRepository interface {
	Create(project *models.Project) bool
	FindByUser(userId uint) ([]*models.Project, bool)
	FindById(id string) (*models.Project, bool)
	Update(project *models.Project) bool
	DeleteById(id string) bool
}
