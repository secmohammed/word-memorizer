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
    TokenRepository       model.TokenRepository
    PrivKey               *rsa.PrivateKey
    PubKey                *rsa.PublicKey
    RefreshSecret         string
    IDExpirationSecs      int64
    RefreshExpirationSecs int64
}

type TokenServiceConfig struct {
    TokenRepository       model.TokenRepository
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
        TokenRepository:       c.TokenRepository,
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
    // set freshly minted refresh token to valid list
    if err := s.TokenRepository.SetRefreshToken(ctx, u.UID.String(), refreshToken.ID, refreshToken.ExpiresIn); err != nil {
        log.Printf("Error storing tokenID for uid: %v. Error: %v\n", u.UID, err.Error())
        return nil, errors.NewInternal()
    }

    // delete user's current refresh token (used when refreshing idToken)
    if prevTokenID != "" {
        if err := s.TokenRepository.DeleteRefreshToken(ctx, u.UID.String(), prevTokenID); err != nil {
            log.Printf("Could not delete previous refreshToken for uid: %v, tokenID: %v\n", u.UID.String(), prevTokenID)
        }
    }
    return &model.TokenPair{
        IDToken:      idToken,
        RefreshToken: refreshToken.SS,
    }, nil
}

// ValidateIDToken validates the id token jwt string
// It returns the user extract from the IDTokenCustomClaims
func (s *tokenService) ValidateIDToken(tokenString string) (*model.User, error) {
    claims, err := validateIDToken(tokenString, s.PubKey) // uses public RSA key

    // We'll just return unauthorized error in all instances of failing to verify user
    if err != nil {
        log.Printf("Unable to validate or parse idToken - Error: %v\n", err)
        return nil, errors.NewAuthorization("Unable to verify user from idToken")
    }

    return claims.User, nil
}
