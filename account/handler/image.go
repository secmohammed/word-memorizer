package handler

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/secmohammed/word-memorizer/account/errors"
    "github.com/secmohammed/word-memorizer/account/model"
)

func (h *Handler) Image(c *gin.Context) {
    authUser := c.MustGet("user").(*model.User)
    c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, h.MaxBodyBytes)
    imageFileHeader, err := c.FormFile("imageFile")
    if err != nil {
        log.Printf("Unable parse multipart/form-data: %+v", err)
        if err.Error() == "http: request body too large" {
            c.JSON(http.StatusRequestEntityTooLarge, gin.H{
                "error": fmt.Sprintf("Max request body size is %v bytes\n", h.MaxBodyBytes),
            })
            return
        }
        e := errors.NewBadRequest("unable to parse multipart/form-data")
        c.JSON(e.Status(), gin.H{
            "error": e,
        })
        return
    }
    mimeType := imageFileHeader.Header.Get("Content-Type")
    if valid := isAllowedImageType(mimeType); !valid {
        log.Println("Image mime type isn't allowed")
        e := errors.NewBadRequest("Image file must be 'image/jpeg' or 'image/jpg'")
        c.JSON(e.Status(), gin.H{
            "error": e,
        })
        return
    }
    ctx := c.Request.Context()
    updatedUser, err := h.UserService.SetProfileImage(ctx, authUser.UID, imageFileHeader)
    if err != nil {
        c.JSON(errors.Status(err), gin.H{
            "error": err,
        })
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "imageUrl": updatedUser.ImageURL,
        "message":  "success",
    })
    return
}
func (h *Handler) DeleteImage(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "hello": "spacce persons",
    })
}
