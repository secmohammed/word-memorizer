package model

import (
    "context"
    "mime/multipart"
    "time"

    "github.com/google/uuid"
)

//UserService defines methods the handler layer expects
//any service it interacts with to implement
type UserService interface {
    Get(ctx context.Context, uid uuid.UUID) (*User, error)
    Signup(ctx context.Context, u *User) error
    Signin(ctx context.Context, u *User) error
    UpdateDetails(ctx context.Context, u *User) error
    SetProfileImage(ctx context.Context, uid uuid.UUID, imageFileHeader *multipart.FileHeader) (*User, error)
}

//TokenService defines methods the handler layer expects
//any service it interacts with to implement
type TokenService interface {
    NewPairFromUser(ctx context.Context, u *User, prevTokenID string) (*TokenPair, error)
    ValidateIDToken(tokenString string) (*User, error)
    ValidateRefreshToken(refreshTokenString string) (*RefreshToken, error)
    Signout(ctx context.Context, uid uuid.UUID) error
}

//UserRepository defines methodds the service lay expects
//any repository it interacts with to implement
type UserRepository interface {
    FindByEmail(ctx context.Context, email string) (*User, error)
    FindByID(ctx context.Context, uid uuid.UUID) (*User, error)
    Create(ctx context.Context, u *User) error
    Update(ctx context.Context, u *User) error
    UpdateImage(ctx context.Context, uid uuid.UUID, imageURL string) (*User, error)
}

// TokenRepository defines methods it expects a repository
// it interacts with to implement
type TokenRepository interface {
    SetRefreshToken(ctx context.Context, userID string, tokenID string, expiresIn time.Duration) error
    DeleteRefreshToken(ctx context.Context, userID string, prevTokenID string) error
    DeleteUserRefreshTokens(ctx context.Context, userID string) error
}
type ImageRepository interface {
    UpdateProfile(ctx context.Context, objName string, imageFile multipart.File) (string, error)
}
