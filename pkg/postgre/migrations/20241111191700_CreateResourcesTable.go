package migrations

type createEndpointsTable struct {
}

func CreateEndpointsTableMigration() Migration {
	return createEndpointsTable{}
}

func (cut createEndpointsTable) GetId() string {
	return "20241111191500_CreateEndpointsTable"
}

func (cut createEndpointsTable) GetDescription() string {
	return "Cria tabela de endpoints"
}

func (cut createEndpointsTable) Operations() []string {
	operations := make([]string, 1)
	query := `
		CREATE TABLE IF NOT EXISTS endpoints (
			id SERIAL PRIMARY KEY,
			path VARCHAR(50) NOT NULL,
			method VARCHAR(10) NOT NULL,
			schema VARCHAR(300) NOT NULL,
			projectId VARCHAR(32) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (projectId) REFERENCES projects(id) ON DELETE CASCADE
		);
	`
	operations[0] = query
	return operations
}
