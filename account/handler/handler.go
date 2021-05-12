package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/secmohammed/word-memorizer/account/model"
)

type Handler struct {
    UserService  model.UserService
    TokenService model.TokenService
}

type Config struct {
    R            *gin.Engine
    UserService  model.UserService
    TokenService model.TokenService
    BaseURL      string
}

func NewHandler(c *Config) {
    h := &Handler{
        UserService:  c.UserService,
        TokenService: c.TokenService,
    }

    g := c.R.Group(c.BaseURL)
    g.GET("/me", h.Me)
    g.POST("/signup", h.Signup)
    g.POST("/signout", h.Signout)
    g.POST("/signin", h.Signin)
    g.POST("/tokens", h.Tokens)
    g.POST("/image", h.Image)
    g.DELETE("/image", h.DeleteImage)
    g.PUT("/details", h.Details)
}

func (h *Handler) Signin(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "hello": "spacce persons",
    })
}
func (h *Handler) Signout(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "hello": "spacce persons",
    })
}
func (h *Handler) Image(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "hello": "spacce persons",
    })
}
func (h *Handler) DeleteImage(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "hello": "spacce persons",
    })
}
func (h *Handler) Details(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "hello": "spacce persons",
    })
}
