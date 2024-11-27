package responses

import "github.com/no-code-api/no-code-api/internal/projects/domain/models"

type ProjectResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (projectResponse *ProjectResponse) FromModel(project *models.Project) {
	projectResponse.Id = project.Id
	projectResponse.Name = project.Name
	projectResponse.Description = project.Description
}
