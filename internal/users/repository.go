package users

import (
	"github.com/leandro-d-santos/no-code-api/internal/logger"
	"github.com/leandro-d-santos/no-code-api/pkg/database"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *User) error
	FindAll() ([]*User, error)
	FindById(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	Update(user *User) error
	Delete(id uint) error
}

type userRepository struct {
	db   *gorm.DB
	logg *logger.Logger
}

func NewRepository() IUserRepository {
	return &userRepository{
		db:   database.GetDb(),
		logg: logger.NewLogger("UserRepository"),
	}
}

func (r *userRepository) Create(user *User) error {
	user.Id = 0
	user.setCreatedAt()
	user.setUpdatedAt()
	result := r.db.Create(user)
	if result.Error != nil {
		r.logg.ErrorF("Error creating user: %v", result.Error.Error())
		return result.Error
	}
	return nil
}

func (r *userRepository) FindAll() ([]*User, error) {
	var users []*User
	result := r.db.Find(&users)
	if result.Error != nil {
		r.logg.ErrorF("Error find users: %v", result.Error.Error())
		return nil, result.Error
	}
	return users, nil
}

func (r *userRepository) FindById(id uint) (*User, error) {
	user := &User{}
	filter := &filter{Id: id}
	result := r.db.Find(user, filter)
	if result.Error != nil {
		r.logg.ErrorF("Error find user (%d): %v", id, result.Error.Error())
		return nil, result.Error
	}
	if user.Id == 0 {
		user = nil
	}
	return user, nil
}

func (r *userRepository) FindByEmail(email string) (*User, error) {
	user := &User{}
	filter := &filter{Email: email}
	result := r.db.Find(user, filter)
	if result.Error != nil {
		r.logg.ErrorF("Error find user by email(%s): %v", email, result.Error.Error())
		return nil, result.Error
	}
	if user.Id == 0 {
		user = nil
	}
	return user, nil
}

func (r *userRepository) Update(user *User) error {
	user.setUpdatedAt()
	result := r.db.Save(user)
	if result.Error != nil {
		r.logg.ErrorF("Error update user: %v", result.Error.Error())
		return result.Error
	}
	return nil
}

func (r *userRepository) Delete(id uint) error {
	filter := &filter{Id: id}
	result := r.db.Delete(&User{}, filter)
	if result.Error != nil {
		r.logg.ErrorF("Error delete user: %v", result.Error.Error())
		return result.Error
	}
	return nil
}
