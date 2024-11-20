package migrations

import "github.com/leandro-d-santos/no-code-api/pkg/postgre/utils"

type createResourcesTables struct {
}

func CreateResourcesTablesMigration() Migration {
	return createResourcesTables{}
}

func (crt createResourcesTables) GetId() string {
	return "20241111191500_CreateResourcesTables"
}

func (crt createResourcesTables) GetDescription() string {
	return "Cria tabelas de recursos"
}

func (crt createResourcesTables) Operations() []string {
	operations := make([]string, 2)
	operations[0] = createResourceTableQuery()
	operations[1] = createEndpointTableQuery()
	return operations
}

func createResourceTableQuery() string {
	return utils.NewStringBuilder().
		AppendLine("CREATE TABLE IF NOT EXISTS resources (").
		AppendLine("id VARCHAR(32) PRIMARY KEY,").
		AppendLine("projectId VARCHAR(32) NOT NULL,").
		AppendLine("path VARCHAR(50) NOT NULL,").
		AppendLine("createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,").
		AppendLine("updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,").
		AppendLine("FOREIGN KEY (projectId) REFERENCES projects(id) ON DELETE CASCADE").
		AppendLine(")").
		String()
}

func createEndpointTableQuery() string {
	return utils.NewStringBuilder().
		AppendLine("CREATE TABLE IF NOT EXISTS endpoints (").
		AppendLine("id VARCHAR(32) PRIMARY KEY,").
		AppendLine("path VARCHAR(50) NOT NULL,").
		AppendLine("method VARCHAR(10) NOT NULL,").
		AppendLine("resourceId VARCHAR(32) NOT NULL,").
		AppendLine("createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,").
		AppendLine("updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,").
		AppendLine("FOREIGN KEY (resourceId) REFERENCES resources(id) ON DELETE CASCADE").
		AppendLine(")").
		String()
}
