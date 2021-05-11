package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func (h *Handler) Tokens(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "hello": "spacce persons",
    })
}
