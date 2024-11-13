package users

import (
	"github.com/leandro-d-santos/no-code-api/internal/logger"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre/utils"
)

type IUserRepository interface {
	Create(user *User) (ok bool)
	FindAll() (users []*User, ok bool)
	FindById(id uint) (user *User, ok bool)
	FindByEmail(email string) (user *User, ok bool)
	Update(user *User) (ok bool)
	Delete(id uint) (ok bool)
}

type userRepository struct {
	connection *postgre.Connection
	logg       *logger.Logger
}

func NewRepository(connection *postgre.Connection) IUserRepository {
	return &userRepository{
		connection: connection,
		logg:       logger.NewLogger("UserRepository"),
	}
}

func (r *userRepository) Create(user *User) (ok bool) {
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

func (r *userRepository) FindAll() (users []*User, ok bool) {
	return r.FindUsers(&UserFilter{})
}

func (r *userRepository) FindById(id uint) (user *User, ok bool) {
	users, ok := r.FindUsers(&UserFilter{Id: id})
	if !ok {
		return nil, false
	}
	user = nil
	if len(users) > 0 {
		user = users[0]
	}
	return user, true
}

func (r *userRepository) FindByEmail(email string) (*User, bool) {
	users, ok := r.FindUsers(&UserFilter{Email: email})
	if !ok {
		return nil, false
	}
	var user *User = nil
	if len(users) > 0 {
		user = users[0]
	}
	return user, true
}

func (r *userRepository) Update(user *User) bool {
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

func (r *userRepository) FindUsers(filter *UserFilter) ([]*User, bool) {
	query := utils.NewStringBuilder()
	query.AppendLine(r.GetQuery())
	query.AppendLine(r.GetQueryFilter(filter))
	result, err := r.connection.ExecuteQuery(query.String())
	if err != nil {
		return nil, false
	}

	var users []*User
	for result.Next() {
		user := &User{
			Id:       uint(result.ReadUint("id")),
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

func (r *userRepository) GetQueryFilter(filter *UserFilter) string {
	query := utils.NewStringBuilder()
	query.AppendLine("WHERE 1=1")
	if filter.Id > 0 {
		query.AppendFormat("AND id=%d", filter.Id)
	}
	if filter.Email != "" {
		query.AppendFormat("AND email=%s", utils.SqlString(filter.Email))
	}
	return query.String()
}
