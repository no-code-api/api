package migrations

type createProjectTable struct {
}

func CreateProjectTableMigration() Migration {
	return createProjectTable{}
}

func (cut createProjectTable) GetId() string {
	return "20241111191500_CreateProjectTable"
}

func (cut createProjectTable) GetDescription() string {
	return "Cria tabela de projetos"
}

func (cut createProjectTable) Operations() []string {
	operations := make([]string, 1)
	query := `
		CREATE TABLE IF NOT EXISTS projects (
			id VARCHAR(32) PRIMARY KEY,
			userId INTEGER NOT NULL,
			name VARCHAR(30) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE
		);
	`
	operations[0] = query
	return operations
}
