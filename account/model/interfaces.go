package model

import "github.com/google/uuid"

//UserService defines methods the handler layer expects
//any service it interacts with to implement
type UserService interface {
    Get(uid uuid.UUID) (*User, error)
}

//UserRepository defines methodds the service lay expects
//any repository it interacts with to implement
type UserRepository interface {
    FindById(uid uuid.UUID) (*User, error)
}
