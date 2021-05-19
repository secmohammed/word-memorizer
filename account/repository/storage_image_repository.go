package repository

import (
    "cloud.google.com/go/storage"
    "github.com/secmohammed/word-memorizer/account/model"
)

type storageImageRepository struct {
    Storage    *storage.Client
    BucketName string
}

func NewImageRepository(s *storage.Client, bucketName string) model.ImageRepository {
    return &storageImageRepository{
        Storage:    s,
        BucketName: bucketName,
    }
}
