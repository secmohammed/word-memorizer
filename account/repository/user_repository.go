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

// FindByEmail retrieves user row by email address
func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
    user := &model.User{}

    query := "SELECT * FROM users WHERE email=$1"

    if err := r.DB.GetContext(ctx, user, query, email); err != nil {
        log.Printf("Unable to get user with email address: %v. Err: %v\n", email, err)
        return user, errors.NewNotFound("email", email)
    }

    return user, nil
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

// Update updates a user's properties
func (r *userRepository) Update(ctx context.Context, u *model.User) error {
    query := `
        UPDATE users 
        SET name=:name, email=:email, website=:website
        WHERE uid=:uid
        RETURNING *;
    `

    nstmt, err := r.DB.PrepareNamedContext(ctx, query)

    if err != nil {
        log.Printf("Unable to prepare user update query: %v\n", err)
        return errors.NewInternal()
    }

    if err := nstmt.GetContext(ctx, u, u); err != nil {
        log.Printf("Failed to update details for user: %v\n", u)
        return errors.NewInternal()
    }

    return nil
}
func (r *userRepository) UpdateImage(ctx context.Context, uid uuid.UUID, imageURL string) (*model.User, error) {
    query := `UPDATE users SET image_url=$2 WHERE uid=$1 RETURNING *;`
    u := &model.User{}
    err := r.DB.GetContext(ctx, u, query, uid, imageURL)
    if err != nil {
        log.Printf("Error updating image_url in db: %v\n", err)
        return nil, errors.NewInternal()
    }
    return u, nil
}
