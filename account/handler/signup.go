package handler

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/secmohammed/word-memorizer/account/errors"
    "github.com/secmohammed/word-memorizer/account/model"
)

type signupReq struct {
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,gte=6,lte=30"`
}

//Signup is used to register the singup endpoint
func (h *Handler) Signup(c *gin.Context) {
    var req signupReq
    if ok := bindData(c, &req); !ok {
        return
    }
    u := &model.User{
        Email:    req.Email,
        Password: req.Password,
    }
    err := h.UserService.Signup(c, u)
    if err != nil {
        log.Printf("Failed to signup user: %v\n", err.Error())
        c.JSON(errors.Status(err), gin.H{"error": err})
        return
    }
}
