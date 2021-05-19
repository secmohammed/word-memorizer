package service

import (
    "context"
    "log"
    "mime/multipart"
    "net/url"
    "path"

    "github.com/google/uuid"
    "github.com/secmohammed/word-memorizer/account/errors"
    "github.com/secmohammed/word-memorizer/account/model"
    "github.com/secmohammed/word-memorizer/account/utils"
)

// userService acts as a struct for injecting an implementation of UserRepository
// for use in service methods
type userService struct {
    UserRepository  model.UserRepository
    ImageRepository model.ImageRepository
}

type UserServiceConfig struct {
    UserRepository  model.UserRepository
    ImageRepository model.ImageRepository
}

func NewUserService(c *UserServiceConfig) model.UserService {
    return &userService{
        UserRepository:  c.UserRepository,
        ImageRepository: c.ImageRepository,
    }
}

func (s *userService) Signup(ctx context.Context, u *model.User) error {
    password, err := utils.HashPassword(u.Password)
    if err != nil {
        return err
    }
    u.Password = password
    if err := s.UserRepository.Create(ctx, u); err != nil {
        return err
    }
    return nil
}

func (s *userService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
    u, err := s.UserRepository.FindByID(ctx, uid)

    return u, err
}

// Signin reaches our to a UserRepository check if the user exists
// and then compares the supplied password with the provided password.
// If a valid email/password combo is provided, u will hold all
// available user fields
func (s *userService) Signin(ctx context.Context, u *model.User) error {
    uFetched, err := s.UserRepository.FindByEmail(ctx, u.Email)

    // Will return NotAuthorized to client to omit details of why
    if err != nil {
        return errors.NewAuthorization("Invalid email and password combination")
    }

    // verify password - we previously created this method
    match := utils.CheckPassword(u.Password, uFetched.Password)

    if !match {
        return errors.NewAuthorization("Invalid email and password combination")
    }

    *u = *uFetched
    return nil
}
func (s *userService) UpdateDetails(ctx context.Context, u *model.User) error {
    // Update user in UserRepository
    err := s.UserRepository.Update(ctx, u)

    if err != nil {
        return err
    }

    // // Publish user updated
    // err = s.EventsBroker.PublishUserUpdated(u, false)
    // if err != nil {
    //  return apperrors.NewInternal()
    // }

    return nil
}
func (s *userService) SetProfileImage(ctx context.Context, uid uuid.UUID, imageFileHeader *multipart.FileHeader) (*model.User, error) {
    u, err := s.UserRepository.FindByID(ctx, uid)
    if err != nil {
        return nil, err
    }
    objName, err := objectNameFromURL(u.ImageURL)
    if err != nil {
        return nil, err
    }
    imageFile, err := imageFileHeader.Open()
    if err != nil {
        log.Printf("Failed to open image file: %v\n", err)
        return nil, errors.NewInternal()
    }
    imageURL, err := s.ImageRepository.UpdateProfile(ctx, objName, imageFile)
    if err != nil {
        log.Printf("Failed to upload image to cloud provider:%v \n", err)
        return nil, err
    }
    updatedUser, err := s.UserRepository.UpdateImage(ctx, u.UID, imageURL)
    if err != nil {
        log.Printf("Failed to update image url: %v\n", err)
        return nil, err
    }

    return updatedUser, nil
}
func objectNameFromURL(imageURL string) (string, error) {
    if imageURL != "" {
        objID, _ := uuid.NewRandom()
        return objID.String(), nil
    }
    urlPath, err := url.Parse(imageURL)
    if err != nil {
        log.Printf("Failed to parse objectName from imageURL: %v\n", imageURL)
        return "", errors.NewInternal()
    }
    return path.Base(urlPath.Path), nil
}
