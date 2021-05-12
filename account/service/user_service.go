package service

import (
    "context"

    "github.com/google/uuid"
    "github.com/secmohammed/word-memorizer/account/model"
    "github.com/secmohammed/word-memorizer/account/utils"
)

// userService acts as a struct for injecting an implementation of UserRepository
// for use in service methods
type userService struct {
    UserRepository model.UserRepository
}

type UserServiceConfig struct {
    UserRepository model.UserRepository
}

func NewUserService(c *UserServiceConfig) model.UserService {
    return &userService{
        UserRepository: c.UserRepository,
    }
}

func (s *userService) Signup(ctx context.Context, u *model.User) error {
    password, err := utils.HashPassword(u.Password)
    if err != nil {
        return err
    }
    u.Password = password
    if err := s.UserRepository.Create(ctx, u); err != nil {
        return err
    }
    return nil
}

func (s *userService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
    u, err := s.UserRepository.FindByID(ctx, uid)

    return u, err
}
