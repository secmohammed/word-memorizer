package handler

import (
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/secmohammed/word-memorizer/account/errors"
    "github.com/secmohammed/word-memorizer/account/model"
)

// omitempty must be listed first (tags evaluated sequentially, it seems)
type detailsReq struct {
    Name    string `json:"name" binding:"omitempty,max=50"`
    Email   string `json:"email" binding:"required,email"`
    Website string `json:"website" binding:"omitempty,url"`
}

// Details handler
func (h *Handler) Details(c *gin.Context) {
    authUser := c.MustGet("user").(*model.User)

    var req detailsReq

    if ok := bindData(c, &req); !ok {
        return
    }

    // Should be returned with current imageURL
    u := &model.User{
        UID:     authUser.UID,
        Name:    req.Name,
        Email:   req.Email,
        Website: req.Website,
    }

    ctx := c.Request.Context()
    err := h.UserService.UpdateDetails(ctx, u)

    if err != nil {
        log.Printf("Failed to update user: %v\n", err.Error())

        c.JSON(errors.Status(err), gin.H{
            "error": err,
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "user": u,
    })
}
