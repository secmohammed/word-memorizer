package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"

    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "github.com/secmohammed/word-memorizer/account/handler"
    "github.com/secmohammed/word-memorizer/account/repository"
    "github.com/secmohammed/word-memorizer/account/service"
)

// will initialize a handler starting from data sources
// which inject into repository layer
// which inject into service layer
// which inject into handler layer
func inject(d *dataSources) (*gin.Engine, error) {
    log.Println("Injecting data sources")

    /*
     * repository layer
     */
    userRepository := repository.NewUserRepository(d.DB)

    /*
     * repository layer
     */
    userService := service.NewUserService(&service.UserServiceConfig{
        UserRepository: userRepository,
    })

    // load rsa keys
    privKeyFile := os.Getenv("PRIV_KEY_FILE")
    priv, err := ioutil.ReadFile(privKeyFile)

    if err != nil {
        return nil, fmt.Errorf("could not read private key pem file: %w", err)
    }

    privKey, err := jwt.ParseRSAPrivateKeyFromPEM(priv)

    if err != nil {
        return nil, fmt.Errorf("could not parse private key: %w", err)
    }

    pubKeyFile := os.Getenv("PUB_KEY_FILE")
    pub, err := ioutil.ReadFile(pubKeyFile)

    if err != nil {
        return nil, fmt.Errorf("could not read public key pem file: %w", err)
    }

    pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pub)

    if err != nil {
        return nil, fmt.Errorf("could not parse public key: %w", err)
    }

    // load refresh token secret from env variable
    refreshSecret := os.Getenv("REFRESH_SECRET")

    tokenService := service.NewTokenService(&service.TokenServiceConfig{
        PrivKey:       privKey,
        PubKey:        pubKey,
        RefreshSecret: refreshSecret,
    })

    // initialize gin.Engine
    gin.SetMode(gin.ReleaseMode)

    router := gin.Default()

    handler.NewHandler(&handler.Config{
        R:            router,
        UserService:  userService,
        TokenService: tokenService,
    })

    return router, nil
}