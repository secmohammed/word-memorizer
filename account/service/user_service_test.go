package service

import (
    "context"
    "fmt"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "github.com/secmohammed/word-memorizer/account/errors"
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

func TestSignup(t *testing.T) {
    t.Run("Success", func(t *testing.T) {
        uid, _ := uuid.NewRandom()

        mockUser := &model.User{
            Email:    "bob@bob.com",
            Password: "howdyhoneighbor!",
        }

        mockUserRepository := new(mocks.MockUserRepository)
        us := NewUserService(&UserServiceConfig{
            UserRepository: mockUserRepository,
        })

        // We can use Run method to modify the user when the Create method is called.
        //  We can then chain on a Return method to return no error
        mockUserRepository.
            On("Create", mock.AnythingOfType("*context.emptyCtx"), mockUser).
            Run(func(args mock.Arguments) {
                userArg := args.Get(1).(*model.User) // arg 0 is context, arg 1 is *User
                userArg.UID = uid
            }).Return(nil)

        ctx := context.TODO()
        err := us.Signup(ctx, mockUser)

        assert.NoError(t, err)

        // assert user now has a userID
        assert.Equal(t, uid, mockUser.UID)

        mockUserRepository.AssertExpectations(t)
    })

    t.Run("Error", func(t *testing.T) {
        mockUser := &model.User{
            Email:    "bob@bob.com",
            Password: "howdyhoneighbor!",
        }

        mockUserRepository := new(mocks.MockUserRepository)
        us := NewUserService(&UserServiceConfig{
            UserRepository: mockUserRepository,
        })

        mockErr := errors.NewConflict("email", mockUser.Email)

        // We can use Run method to modify the user when the Create method is called.
        //  We can then chain on a Return method to return no error
        mockUserRepository.
            On("Create", mock.AnythingOfType("*context.emptyCtx"), mockUser).
            Return(mockErr)

        ctx := context.TODO()
        err := us.Signup(ctx, mockUser)

        // assert error is error we response with in mock
        assert.EqualError(t, err, mockErr.Error())

        mockUserRepository.AssertExpectations(t)
    })
}
func TestUpdateDetails(t *testing.T) {
    mockUserRepository := new(mocks.MockUserRepository)
    us := NewUserService(&UserServiceConfig{
        UserRepository: mockUserRepository,
    })

    t.Run("Success", func(t *testing.T) {
        uid, _ := uuid.NewRandom()

        mockUser := &model.User{
            UID:     uid,
            Email:   "new@bob.com",
            Website: "https://jacobgoodwin.me",
            Name:    "A New Bob!",
        }

        mockArgs := mock.Arguments{
            mock.AnythingOfType("*context.emptyCtx"),
            mockUser,
        }

        mockUserRepository.
            On("Update", mockArgs...).Return(nil)

        ctx := context.TODO()
        err := us.UpdateDetails(ctx, mockUser)

        assert.NoError(t, err)
        mockUserRepository.AssertCalled(t, "Update", mockArgs...)
    })

    t.Run("Failure", func(t *testing.T) {
        uid, _ := uuid.NewRandom()

        mockUser := &model.User{
            UID: uid,
        }

        mockArgs := mock.Arguments{
            mock.AnythingOfType("*context.emptyCtx"),
            mockUser,
        }

        mockError := errors.NewInternal()

        mockUserRepository.
            On("Update", mockArgs...).Return(mockError)

        ctx := context.TODO()
        err := us.UpdateDetails(ctx, mockUser)
        assert.Error(t, err)

        apperror, ok := err.(*errors.Error)
        assert.True(t, ok)
        assert.Equal(t, errors.Internal, apperror.Type)

        mockUserRepository.AssertCalled(t, "Update", mockArgs...)
    })
}
