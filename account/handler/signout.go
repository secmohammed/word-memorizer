package handler

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/secmohammed/word-memorizer/account/errors"
    "github.com/secmohammed/word-memorizer/account/model"
)

// Signout handler
func (h *Handler) Signout(c *gin.Context) {
    user := c.MustGet("user")

    ctx := c.Request.Context()
    if err := h.TokenService.Signout(ctx, user.(*model.User).UID); err != nil {
        c.JSON(errors.Status(err), gin.H{
            "error": err,
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "user signed out successfully!",
    })
}
