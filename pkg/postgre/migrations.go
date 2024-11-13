package postgre

import (
	"github.com/leandro-d-santos/no-code-api/pkg/postgre/migrations"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre/utils"
)

func RunMigrations(conn *Connection) {
	CreateMigrationsTable(conn)
	run(conn, migrations.CreateUserTableMigration())
	run(conn, migrations.CreateProjectTableMigration())
	run(conn, migrations.CreateEndpointsTableMigration())
}

func CreateMigrationsTable(conn *Connection) {
	query := `
		CREATE TABLE IF NOT EXISTS migrations (
			id VARCHAR(150) NOT NULL PRIMARY KEY,
			description VARCHAR(200) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
	`
	if err := conn.ExecuteNonQuery(query); err != nil {
		logger.Fatal(err.Error())
	}
}

func run(conn *Connection, migration migrations.Migration) {
	if existsMigration(conn, migration.GetId()) {
		return
	}
	for _, operation := range migration.Operations() {
		if err := conn.ExecuteNonQuery(operation); err != nil {
			logger.Fatal(err.Error())
		}
	}
	insertMigration(conn, migration)
}

func insertMigration(conn *Connection, migration migrations.Migration) {
	query := utils.NewStringBuilder()
	query.AppendLine("INSERT INTO migrations")
	query.AppendLine("(id, description)")
	query.AppendFormat("VALUES (%s, %s)", utils.SqlString(migration.GetId()), utils.SqlString(migration.GetDescription()))
	if err := conn.ExecuteNonQuery(query.String()); err != nil {
		logger.Fatal(err.Error())
	}
}

func existsMigration(conn *Connection, migrationId string) bool {
	query := utils.NewStringBuilder()
	query.AppendLine("SELECT COUNT(0)")
	query.AppendLine("FROM migrations")
	query.AppendFormat("WHERE id=%s", utils.SqlString(migrationId))
	value, err := conn.ExecuteSingleQuery(query.String())
	if err != nil {
		logger.Fatal(err.Error())
	}
	return value.(int64) > 0
}
