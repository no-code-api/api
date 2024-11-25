package repositories

import (
	"github.com/leandro-d-santos/no-code-api/internal/logger"
	"github.com/leandro-d-santos/no-code-api/internal/users/domain/models"
	"github.com/leandro-d-santos/no-code-api/internal/users/domain/repositories"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre/utils"
)

type userRepository struct {
	connection *postgre.Connection
	logg       *logger.Logger
}

func NewRepository(connection *postgre.Connection) repositories.IRepository {
	return &userRepository{
		connection: connection,
		logg:       logger.NewLogger("UserRepository"),
	}
}

func (r *userRepository) Create(user *models.User) (ok bool) {
	command := utils.NewStringBuilder()
	command.AppendLine("INSERT INTO users")
	command.AppendLine("(name, email, password, createdAt, updatedAt)")
	command.AppendFormat("VALUES (%s", utils.SqlString(user.Name)).AppendNewLine()
	command.AppendFormat(",%s", utils.SqlString(user.Email)).AppendNewLine()
	command.AppendFormat(",%s", utils.SqlString(user.Password)).AppendNewLine()
	command.AppendLine(",NOW()")
	command.AppendLine(",NOW())")
	if err := r.connection.ExecuteNonQuery(command.String()); err != nil {
		r.logg.ErrorF("error to insert user: %s", err.Error())
		return false
	}
	return true
}

func (r *userRepository) FindAll() (users []*models.User, ok bool) {
	return r.FindUsers(&models.UserFilter{})
}

func (r *userRepository) FindById(id uint) (user *models.User, ok bool) {
	users, ok := r.FindUsers(&models.UserFilter{Id: id})
	if !ok {
		return nil, false
	}
	user = nil
	if len(users) > 0 {
		user = users[0]
	}
	return user, true
}

func (r *userRepository) FindByEmail(email string) (*models.User, bool) {
	users, ok := r.FindUsers(&models.UserFilter{Email: email})
	if !ok {
		return nil, false
	}
	var user *models.User = nil
	if len(users) > 0 {
		user = users[0]
	}
	return user, true
}

func (r *userRepository) Update(user *models.User) bool {
	command := utils.NewStringBuilder()
	command.AppendLine("UPDATE users")
	command.AppendFormat("SET name=%s", utils.SqlString(user.Name)).AppendNewLine()
	command.AppendLine(",updatedAt=NOW()")
	command.AppendFormat("WHERE id=%d", user.Id)
	if err := r.connection.ExecuteNonQuery(command.String()); err != nil {
		r.logg.ErrorF("error to update user: %s", err.Error())
		return false
	}
	return true
}

func (r *userRepository) Delete(id uint) (ok bool) {
	command := utils.NewStringBuilder()
	command.AppendLine("DELETE FROM users")
	command.AppendFormat("WHERE id=%d", id)
	if err := r.connection.ExecuteNonQuery(command.String()); err != nil {
		r.logg.ErrorF("error to delete user: %s", err.Error())
		return false
	}
	return true
}

func (r *userRepository) FindUsers(filter *models.UserFilter) ([]*models.User, bool) {
	query := utils.NewStringBuilder()
	query.AppendLine(r.GetQuery())
	query.AppendLine(r.GetQueryFilter(filter))
	result, err := r.connection.ExecuteQuery(query.String())
	if err != nil {
		return nil, false
	}

	var users []*models.User
	for result.Next() {
		user := &models.User{
			Id:       uint(result.ReadInt("id")),
			Name:     result.ReadString("name"),
			Email:    result.ReadString("email"),
			Password: result.ReadString("password"),
		}
		users = append(users, user)
	}
	return users, true
}

func (r *userRepository) GetQuery() string {
	query := utils.NewStringBuilder()
	query.AppendLine("SELECT id")
	query.AppendLine(",name")
	query.AppendLine(",email")
	query.AppendLine(",password")
	query.AppendLine("FROM users")
	return query.String()
}

func (r *userRepository) GetQueryFilter(filter *models.UserFilter) string {
	query := utils.NewStringBuilder()
	query.AppendLine("WHERE 1=1")
	if filter.Id > 0 {
		query.AppendFormat("AND id=%d", filter.Id).AppendNewLine()
	}
	if filter.Email != "" {
		query.AppendFormat("AND email=%s", utils.SqlString(filter.Email))
	}
	return query.String()
}
