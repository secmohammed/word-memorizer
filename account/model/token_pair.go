package model

//User defines ddomain modal and its json and db representations
type TokenPair struct {
    IDToken      string `db:"id_token" json:"id_token"`
    RefreshToken string `db:"refresh_token" json:"refresh_token"`
}
