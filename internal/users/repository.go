package users

import (
	"github.com/leandro-d-santos/no-code-api/internal/logger"
	"github.com/leandro-d-santos/no-code-api/pkg/database"
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
	connection *database.Connection
	logg       *logger.Logger
}

func NewRepository(connection *database.Connection) IUserRepository {
	return &userRepository{
		connection: connection,
		logg:       logger.NewLogger("UserRepository"),
	}
}

func (r *userRepository) Create(user *User) (ok bool) {
	user.Id = 0
	user.SetCreatedAt()
	user.SetUpdatedAt()
	return r.connection.Save(user, false)
}

func (r *userRepository) FindAll() (users []*User, ok bool) {
	var result []*User
	if ok := r.connection.Find(&result, nil); !ok {
		return nil, false
	}
	return result, true
}

func (r *userRepository) FindById(id uint) (user *User, ok bool) {
	result := &User{}
	filter := &filter{Id: id}
	if ok := r.connection.Find(result, filter); !ok {
		return nil, false
	}
	if result.Id == 0 {
		result = nil
	}
	return result, true
}

func (r *userRepository) FindByEmail(email string) (user *User, ok bool) {
	result := &User{}
	filter := &filter{Email: email}
	if ok := r.connection.Find(result, filter); !ok {
		return nil, false
	}
	if result.Id == 0 {
		result = nil
	}
	return result, true
}

func (r *userRepository) Update(user *User) (ok bool) {
	user.SetUpdatedAt()
	return r.connection.Save(user, true)
}

func (r *userRepository) Delete(id uint) (ok bool) {
	filter := &filter{Id: id}
	return r.connection.Delete(&User{}, filter)
}
