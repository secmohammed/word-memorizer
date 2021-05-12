package service

import (
    "context"
    "crypto/rsa"
    "log"

    "github.com/secmohammed/word-memorizer/account/errors"
    "github.com/secmohammed/word-memorizer/account/model"
)

// tokenService used for injecting an implementation of TokenRepository
// for use in service methods along with keys and secrets for
// signing JWTs
type tokenService struct {
    // TokenRepository model.TokenRepository
    PrivKey               *rsa.PrivateKey
    PubKey                *rsa.PublicKey
    RefreshSecret         string
    IDExpirationSecs      int64
    RefreshExpirationSecs int64
}

type TokenServiceConfig struct {
    PrivKey               *rsa.PrivateKey
    PubKey                *rsa.PublicKey
    RefreshSecret         string
    IDExpirationSecs      int64
    RefreshExpirationSecs int64
}

// NewTokenService is a factory function for
// initializing a UserService with its repository layer dependencies
func NewTokenService(c *TokenServiceConfig) model.TokenService {
    return &tokenService{
        PrivKey:               c.PrivKey,
        PubKey:                c.PubKey,
        RefreshSecret:         c.RefreshSecret,
        IDExpirationSecs:      c.IDExpirationSecs,
        RefreshExpirationSecs: c.RefreshExpirationSecs,
    }
}

func (s *tokenService) NewPairFromUser(ctx context.Context, u *model.User, prevTokenID string) (*model.TokenPair, error) {

    // No need to use a repository for idToken as it is unrelated to any data source
    idToken, err := generateIDToken(u, s.PrivKey, s.IDExpirationSecs)

    if err != nil {
        log.Printf("Error generating idToken for uid: %v. Error: %v\n", u.UID, err.Error())
        return nil, errors.NewInternal()
    }

    refreshToken, err := generateRefreshToken(u.UID, s.RefreshSecret, s.RefreshExpirationSecs)

    if err != nil {
        log.Printf("Error generating refreshToken for uid: %v. Error: %v\n", u.UID, err.Error())
        return nil, errors.NewInternal()
    }

    // TODO: store refresh tokens by calling TokenRepository methods

    return &model.TokenPair{
        IDToken:      idToken,
        RefreshToken: refreshToken.SS,
    }, nil
}
