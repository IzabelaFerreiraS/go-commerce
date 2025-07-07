package services

import (
	"errors"

	"go-commerce/dtos/request"
	"go-commerce/repositories"
	"go-commerce/schemas"
)

var ErrUserNotFound = errors.New("user not found")

type UserService interface {
    CreateUser(request.CreateUserRequest) (schemas.User, error)
    DeleteUser(id string) error
    ListUsers() ([]schemas.User, error)
    ShowUser(id string) (schemas.User, error)
    UpdateUser(id string, req request.UpdatedUserRequest) (schemas.User, error)
}

type userService struct {
    repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
    return &userService{repo: repo}
}

func (s *userService) CreateUser(req request.CreateUserRequest) (schemas.User, error) {
    user := schemas.User{
        Name:     req.Name,
        Email:    req.Email,
        Password: req.Password,
        Role:     req.Role,
    }
    err := s.repo.Create(&user)
    return user, err
}

func (s *userService) DeleteUser(id string) error {
    user, err := s.repo.FindByID(id)
    if err != nil {
        return err
    }
    if user.ID == 0 {
        return ErrUserNotFound
    }
    return s.repo.Delete(&user)
}

func (s *userService) ListUsers() ([]schemas.User, error) {
    return s.repo.List()
}

func (s *userService) ShowUser(id string) (schemas.User, error) {
    user, err := s.repo.FindByID(id)
    if err != nil {
        return schemas.User{}, err
    }
    if user.ID == 0 {
        return schemas.User{}, ErrUserNotFound
    }
    return user, nil
}

func (s *userService) UpdateUser(id string, req request.UpdatedUserRequest) (schemas.User, error) {
    user, err := s.repo.FindByID(id)
    if err != nil {
        return schemas.User{}, err
    }
    if user.ID == 0 {
        return schemas.User{}, ErrUserNotFound
    }

    if req.Name != "" {
        user.Name = req.Name
    }
    if req.Email != "" {
        user.Email = req.Email
    }
    if req.Password != "" {
        user.Password = req.Password
    }
    if req.Role != "" {
        user.Role = req.Role
    }

    err = s.repo.Update(&user)
    return user, err
}
