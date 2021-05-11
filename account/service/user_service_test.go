package service

import (
    "context"
    "fmt"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "github.com/secmohammed/word-memorizer/account/model"
    "github.com/secmohammed/word-memorizer/account/model/mocks"
    "github.com/stretchr/testify/mock"

    "github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
    gin.SetMode(gin.TestMode)
    t.Run("Error", func(t *testing.T) {
        uid, _ := uuid.NewRandom()

        mockUserRepository := new(mocks.MockUserRepository)
        us := NewUserService(&UserServiceConfig{
            UserRepository: mockUserRepository,
        })

        mockUserRepository.On("FindByID", mock.Anything, uid).Return(nil, fmt.Errorf("Some error down the call chain"))

        ctx := context.TODO()
        u, err := us.Get(ctx, uid)

        assert.Nil(t, u)
        assert.Error(t, err)
        mockUserRepository.AssertExpectations(t)
    })
    t.Run("Success", func(t *testing.T) {
        uid, _ := uuid.NewRandom()

        mockUserResp := &model.User{
            UID:   uid,
            Email: "bob@bob.com",
            Name:  "Bobby Bobson",
        }

        mockUserRepository := new(mocks.MockUserRepository)
        us := NewUserService(&UserServiceConfig{
            UserRepository: mockUserRepository,
        })
        mockUserRepository.On("FindByID", mock.Anything, uid).Return(mockUserResp, nil)

        ctx := context.TODO()
        u, err := us.Get(ctx, uid)

        assert.NoError(t, err)
        assert.Equal(t, u, mockUserResp)
        mockUserRepository.AssertExpectations(t)
    })
}
