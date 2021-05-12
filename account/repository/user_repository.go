package repository

import (
    "context"
    "log"

    "github.com/google/uuid"
    "github.com/lib/pq"

    "github.com/jmoiron/sqlx"
    "github.com/secmohammed/word-memorizer/account/errors"
    "github.com/secmohammed/word-memorizer/account/model"
)

type userRepository struct {
    DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) model.UserRepository {
    return &userRepository{
        DB: db,
    }
}
func (ur *userRepository) FindByID(ctx context.Context, uid uuid.UUID) (*model.User, error) {
    user := &model.User{}

    query := "SELECT * FROM users WHERE uid=$1"

    // we need to actually check errors as it could be something other than not found
    if err := ur.DB.GetContext(ctx, user, query, uid); err != nil {
        return user, errors.NewNotFound("uid", uid.String())
    }

    return user, nil
}
func (ur *userRepository) Create(ctx context.Context, u *model.User) error {
    query := "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING *"
    if err := ur.DB.GetContext(ctx, u, query, u.Email, u.Password); err != nil {
        if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
            log.Printf("Couldn't create a user with email: %v. Reason: %v\n", u.Email, err.Code.Name())
            return errors.NewConflict("email", u.Email)

        }
        log.Printf("cannot create  a user with email: %v, Reason: %v\n", u.Email, err)
        return errors.NewInternal()
    }

    return nil
}
