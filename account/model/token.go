package model

import "github.com/google/uuid"

type RefreshToken struct {
    ID  uuid.UUID `json:"-"`
    UID uuid.UUID `json:"-"`
    SS  string    `json:"refreshToken"`
}
type IDToken struct {
    SS string `json:"idToken"`
}

//TokenPair defines ddomain modal and its json and db representations
type TokenPair struct {
    IDToken
    RefreshToken
}
