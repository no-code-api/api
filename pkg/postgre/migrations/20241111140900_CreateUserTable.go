package migrations

type createUserTable struct {
}

func CreateUserTableMigration() Migration {
	return createUserTable{}
}

func (cut createUserTable) GetId() string {
	return "20241111140900_CreateUserTable"
}

func (cut createUserTable) GetDescription() string {
	return "Cria tabela de usuários"
}

func (cut createUserTable) Operations() []string {
	operations := make([]string, 1)
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(150) NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			password VARCHAR(60) NOT NULL,
			createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`
	operations[0] = query
	return operations
}
