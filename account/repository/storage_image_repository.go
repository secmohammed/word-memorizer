package repository

import (
    "context"
    "fmt"
    "io"
    "mime/multipart"
    "os"

    "cloud.google.com/go/storage"
    "github.com/secmohammed/word-memorizer/account/errors"
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
func (r *storageImageRepository) UpdateProfile(ctx context.Context, objName string, imageFile multipart.File) (string, error) {
    buck := r.Storage.Bucket(r.BucketName)
    object := buck.Object(objName)
    wc := object.NewWriter(ctx)
    wc.ObjectAttrs.CacheControl = "Cache-Control: no-cache, max-age=0"
    if _, err := io.Copy(wc, imageFile); err != nil {
        return "", errors.NewInternal()
    }
    if err := wc.Close(); err != nil {
        return "", fmt.Errorf("Writer.close: %v", err)
    }
    imageURL := fmt.Sprintf("%s/%s", os.Getenv("GCS_GCS_BUCKET_LINK"), objName)

    return imageURL, nil
}
