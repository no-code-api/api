package users

import (
	"errors"

	"github.com/leandro-d-santos/no-code-api/internal/jwt"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre"
)

type UserService struct {
	userRepository IUserRepository
}

func NewService() UserService {
	return UserService{
		userRepository: NewRepository(postgre.GetConnection()),
	}
}

func (s UserService) Login(request *loginRequest) (*loginResponse, error) {
	user, ok := s.userRepository.FindByEmail(request.Email)
	if !ok {
		return nil, s.getUserSearchError()
	}

	if user == nil {
		return nil, errors.New("email não existe")
	}

	if !VerifyPassword(request.Password, user.Password) {
		return nil, errors.New("senha inválida")
	}

	service := jwt.NewJwtService()
	token, err := service.GenerateJWT(user.Id)
	if err != nil {
		return nil, errors.New("erro ao gerar token")
	}

	return &loginResponse{Token: token}, nil
}

func (s UserService) FindAll() ([]UserResponse, error) {
	users, ok := s.userRepository.FindAll()
	if !ok {
		return nil, errors.New("erro ao consultar usuários")
	}
	usersReponse := make([]UserResponse, len(users))
	for index, user := range users {
		userResponse := UserResponse{}
		userResponse.FromModel(user)
		usersReponse[index] = userResponse
	}
	return usersReponse, nil
}

func (s UserService) FindById(id uint) (*UserResponse, error) {
	user, err := s.SearchUser(id)
	if err != nil {
		return nil, err
	}
	userResponse := &UserResponse{}
	userResponse.FromModel(user)
	return userResponse, nil
}

func (s UserService) Create(request *createUserRequest) error {
	existingUser, ok := s.userRepository.FindByEmail(request.Email)
	if !ok {
		return s.getUserSearchError()
	}
	if existingUser != nil {
		return errors.New("email já cadastrado")
	}

	hash, err := HashPassword(request.Password)
	if err != nil {
		return errors.New("erro ao gerar senha")
	}

	user := request.ToModel()
	user.Password = hash
	if ok := s.userRepository.Create(user); !ok {
		return errors.New("erro ao criar usuário")
	}
	return nil
}

func (s UserService) Update(request *updateUserRequest) error {
	user, err := s.SearchUser(request.Id)
	if err != nil {
		return err
	}
	user.Name = request.Name
	if ok := s.userRepository.Update(user); !ok {
		return errors.New("erro ao salvar usuário")
	}
	return nil
}

func (s UserService) Delete(id uint) error {
	if _, err := s.SearchUser(id); err != nil {
		return err
	}
	if ok := s.userRepository.Delete(uint(id)); !ok {
		return errors.New("erro ao remover usuário")
	}

	jwtService := jwt.NewJwtService()
	jwtService.RemoveStamp(uint(id))
	return nil
}

func (s UserService) getUserSearchError() error {
	return errors.New("erro ao consultar usuário")
}

func (s UserService) SearchUser(id uint) (*User, error) {
	user, ok := s.userRepository.FindById(id)
	if !ok {
		return nil, errors.New("erro ao consultar usuário")
	}
	if user == nil {
		return nil, errors.New("usuário não encontrado")
	}
	return user, nil
}
