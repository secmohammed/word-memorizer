package service

import (
    "context"
    "crypto/rsa"
    "log"

    "github.com/secmohammed/word-memorizer/account/errors"
    "github.com/secmohammed/word-memorizer/account/model"
)

type TokenService struct {
    PrivKey       *rsa.PrivateKey
    PubKey        *rsa.PublicKey
    RefreshSecret string
}

type TokenServiceConfig struct {
    PrivKey       *rsa.PrivateKey
    PubKey        *rsa.PublicKey
    RefreshSecret string
}

func NewTokenService(c *TokenServiceConfig) model.TokenService {
    return &TokenService{
        PrivKey:       c.PrivKey,
        PubKey:        c.PubKey,
        RefreshSecret: c.RefreshSecret,
    }
}

func (s *TokenService) NewPairFromUser(ctx context.Context, u *model.User, prevTokenID string) (*model.TokenPair, error) {

    // No need to use a repository for idToken as it is unrelated to any data source
    idToken, err := generateIDToken(u, s.PrivKey)

    if err != nil {
        log.Printf("Error generating idToken for uid: %v. Error: %v\n", u.UID, err.Error())
        return nil, errors.NewInternal()
    }

    refreshToken, err := generateRefreshToken(u.UID, s.RefreshSecret)

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
