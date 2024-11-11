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
	// var result []*User
	// if ok := r.connection.Find(&result, nil); !ok {
	// 	return nil, false
	// }
	// return result, true
	return nil, true
}

func (r *userRepository) FindById(id uint) (user *User, ok bool) {
	// result := &User{}
	// filter := &filter{Id: id}
	// if ok := r.connection.Find(result, filter); !ok {
	// 	return nil, false
	// }
	// if result.Id == 0 {
	// 	result = nil
	// }
	// return result, true
	return nil, true
}

func (r *userRepository) FindByEmail(email string) (user *User, ok bool) {
	// result := &User{}
	// filter := &filter{Email: email}
	// if ok := r.connection.Find(result, filter); !ok {
	// 	return nil, false
	// }
	// if result.Id == 0 {
	// 	result = nil
	// }
	// return result, true
	return nil, true
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
	// filter := &filter{Id: id}
	// return r.connection.Delete(&User{}, filter)
	return true
}
