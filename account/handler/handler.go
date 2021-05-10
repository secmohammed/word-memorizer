package handler

import (
    "net/http"
    "os"

    "github.com/gin-gonic/gin"
)

type Handler struct{}

type Config struct {
    R *gin.Engine
}

func NewHandler(c *Config) {
    h := &Handler{}

    g := c.R.Group(os.Getenv("ACCOUNT_API_URL"))
    g.GET("/me", h.Me)
    g.POST("/signup", h.Signup)
    g.POST("/signout", h.Signout)
    g.POST("/signin", h.Signin)
    g.POST("/tokens", h.Tokens)
    g.POST("/image", h.Image)
    g.DELETE("/image", h.DeleteImage)
    g.PUT("/details", h.Details)
}

func (h *Handler) Me(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "hello": "spacce persons",
    })
}
func (h *Handler) Signup(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "hello": "spacce persons",
    })
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
func (h *Handler) Tokens(c *gin.Context) {
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
