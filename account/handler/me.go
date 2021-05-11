package handler

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.secom/secmohammed/word-memorizer/errors"
    "github.secom/secmohammed/word-memorizer/model"
)

type authHeader struct {
    IDToken string `header:"Authorization"`
}

// Me handler fetches user from ID token
// so that user can be verified by the server and returned
func (h *Handler) Me(c *gin.Context) {
    // A *model.User will eventually be added to context in middleware
    user, exists := c.Get("user")

    // This shouldn't happen, as our middleware ought to throw an error.
    // This is an extra safety measure
    // We'll extract this logic later as it will be common to all handler
    // methods which require a valid user
    if !exists {
        log.Printf("Unable to extract user from request context for unknown reason: %v\n", c)
        err := errors.NewInternal()
        c.JSON(err.Status(), gin.H{
            "error": err,
        })

        return
    }

    uid := user.(*model.User).UID
    u, err := h.UserService.Get(c, uid)
    if err != nil {
        log.Printf("Unable to find user: %v\n%v", uid, err)
        e := errors.NewNotFound("user", uid.String())

        c.JSON(e.Status(), gin.H{
            "error": e,
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "user": u,
    })
}
