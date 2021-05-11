package service

import (
    "context"

    "github.com/google/uuid"
    "github.com/secmohammed/word-memorizer/account/model"
)

type UserService struct {
    UserRepository model.UserRepository
}
type UserServiceConfig struct {
    UserRepository model.UserRepository
}

func NewUserService(c *UserServiceConfig) model.UserService {
    return &UserService{
        UserRepository: c.UserRepository,
    }
}

func (s *UserService) Signup(ctx context.Context, u *model.User) error {
    panic("Method not implemented")
}

func (s *UserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
    u, err := s.UserRepository.FindByID(ctx, uid)

    return u, err
}
