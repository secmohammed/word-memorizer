package handler

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
    "github.com/secmohammed/word-memorizer/account/errors"
)

type invalidArgument struct {
    Field string `json:"field"`
    Value string `json:"value"`
    Tag   string `json:"tag"`
    Param string `json:"param"`
}

func bindData(c *gin.Context, req interface{}) bool {
    if err := c.ShouldBind(req); err != nil {
        log.Printf("Error binding data: %+v\n", err)
        if errs, ok := err.(validator.ValidationErrors); ok {
            var invalidArgs []invalidArgument
            for _, err := range errs {
                invalidArgs = append(invalidArgs, invalidArgument{
                    err.Field(),
                    err.Value().(string),
                    err.Tag(),
                    err.Param(),
                })
            }
            err := errors.NewBadRequest("Invalid request parameters. See invalidArgs")
            c.JSON(err.Status(), gin.H{
                "error":       err,
                "invalidArgs": invalidArgs,
            })
            return false

        }
        fallbackErr := errors.NewInternal()
        c.JSON(fallbackErr.Status(), gin.H{"error": fallbackErr})
        return false
    }
    return true
}
