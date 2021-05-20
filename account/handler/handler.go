package handler

import (
    "time"

    "github.com/gin-gonic/gin"
    "github.com/secmohammed/word-memorizer/account/errors"
    "github.com/secmohammed/word-memorizer/account/handler/middleware"
    "github.com/secmohammed/word-memorizer/account/model"
)

type Handler struct {
    UserService  model.UserService
    TokenService model.TokenService
    MaxBodyBytes int64
}

type Config struct {
    R               *gin.Engine
    UserService     model.UserService
    TokenService    model.TokenService
    TimeoutDuration time.Duration
    BaseURL         string
    MaxBodyBytes    int64
}

func NewHandler(c *Config) {
    h := &Handler{
        UserService:  c.UserService,
        TokenService: c.TokenService,
        MaxBodyBytes: c.MaxBodyBytes,
    }

    g := c.R.Group(c.BaseURL)
    if gin.Mode() != gin.TestMode {
        g.Use(middleware.Timeout(c.TimeoutDuration, errors.NewServiceUnavailable()))
        g.GET("/me", middleware.AuthUser(h.TokenService), h.Me)
        g.POST("/signout", middleware.AuthUser(h.TokenService), h.Signout)
        g.POST("/image", middleware.AuthUser(h.TokenService), h.Image)
        g.PUT("/details", middleware.AuthUser(h.TokenService), h.Details)
    } else {
        g.POST("/image", h.Image)
        g.GET("/me", h.Me)
        g.PUT("/details", h.Details)
        g.POST("/signout", h.Signout)
    }

    g.POST("/signup", h.Signup)
    g.POST("/signin", h.Signin)
    g.POST("/tokens", h.Tokens)
    g.DELETE("/image", h.DeleteImage)
}
