package services

import (
	"errors"

	"github.com/leandro-d-santos/no-code-api/internal/jwt"
	"github.com/leandro-d-santos/no-code-api/internal/users/application/requests"
	"github.com/leandro-d-santos/no-code-api/internal/users/application/responses"
	dataRep "github.com/leandro-d-santos/no-code-api/internal/users/data/repositories"
	userCore "github.com/leandro-d-santos/no-code-api/internal/users/domain/core"
	"github.com/leandro-d-santos/no-code-api/internal/users/domain/models"
	domainRep "github.com/leandro-d-santos/no-code-api/internal/users/domain/repositories"
	"github.com/leandro-d-santos/no-code-api/pkg/postgre"
)

type userService struct {
	userRepository domainRep.IRepository
}

func NewService(connection *postgre.Connection) IService {
	return userService{
		userRepository: dataRep.NewRepository(connection),
	}
}

func (s userService) Login(request *requests.LoginRequest) (*responses.LoginResponse, error) {
	user, ok := s.userRepository.FindByEmail(request.Email)
	if !ok {
		return nil, s.getUserSearchError()
	}

	if user == nil {
		return nil, errors.New("email não existe")
	}

	if !userCore.VerifyPassword(request.Password, user.Password) {
		return nil, errors.New("senha inválida")
	}

	service := jwt.NewJwtService()
	token, err := service.GenerateJWT(user.Id)
	if err != nil {
		return nil, errors.New("erro ao gerar token")
	}

	return &responses.LoginResponse{Token: token}, nil
}

func (s userService) FindAll() ([]responses.UserResponse, error) {
	users, ok := s.userRepository.FindAll()
	if !ok {
		return nil, errors.New("erro ao consultar usuários")
	}
	usersReponse := make([]responses.UserResponse, len(users))
	for index, user := range users {
		userResponse := responses.UserResponse{}
		userResponse.FromModel(user)
		usersReponse[index] = userResponse
	}
	return usersReponse, nil
}

func (s userService) FindById(id uint) (*responses.UserResponse, error) {
	user, err := s.SearchUser(id)
	if err != nil {
		return nil, err
	}
	userResponse := &responses.UserResponse{}
	userResponse.FromModel(user)
	return userResponse, nil
}

func (s userService) Create(request *requests.CreateUserRequest) error {
	existingUser, ok := s.userRepository.FindByEmail(request.Email)
	if !ok {
		return s.getUserSearchError()
	}
	if existingUser != nil {
		return errors.New("email já cadastrado")
	}

	hash, err := userCore.HashPassword(request.Password)
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

func (s userService) Update(request *requests.UpdateUserRequest) error {
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

func (s userService) DeleteById(id uint) error {
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

func (s userService) getUserSearchError() error {
	return errors.New("erro ao consultar usuário")
}

func (s userService) SearchUser(id uint) (*models.User, error) {
	user, ok := s.userRepository.FindById(id)
	if !ok {
		return nil, errors.New("erro ao consultar usuário")
	}
	if user == nil {
		return nil, errors.New("usuário não encontrado")
	}
	return user, nil
}
