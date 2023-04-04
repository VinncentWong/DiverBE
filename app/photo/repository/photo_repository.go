package repository

import (
	"context"
	"mime/multipart"

	"github.com/VinncentWong/DiverBE/infrastructure"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type IPhotoRepository interface {
	UploadPhoto(file multipart.File) (string, string, error)
	DeletePhoto(publicId string) error
}

type PhotoRepository struct {
	cloudinary *cloudinary.Cloudinary
}

func NewPhotoRepository() IPhotoRepository {
	return &PhotoRepository{
		cloudinary: infrastructure.GetStorage(),
	}
}

func (r *PhotoRepository) UploadPhoto(file multipart.File) (string, string, error) {
	result, err := r.cloudinary.Upload.Upload(context.Background(), file, uploader.UploadParams{
		Folder: "jointscamp",
	})
	if err != nil {
		return "", "", err
	}
	return result.SecureURL, result.PublicID, nil
}

func (r *PhotoRepository) DeletePhoto(publicId string) error {
	_, err := r.cloudinary.Upload.Destroy(context.Background(), uploader.DestroyParams{
		PublicID: publicId,
	})
	return err
}
