package repository

import (
	"context"
	"go-rest-api/model"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type IImageRepository interface {
	UploadImage(image *model.Image) (*uploader.UploadResult, error)
}

type imageRepository struct {
	cld *cloudinary.Cloudinary
}

func NewImageRepository(cld *cloudinary.Cloudinary) IImageRepository {
	return &imageRepository{cld}
}

func (ir *imageRepository) UploadImage(image *model.Image) (*uploader.UploadResult, error) {
	ctx := context.Background()
	// publicId := fmt.Sprintf("tweet_image_%s", image.ImageData)

	uploadResult, err := ir.cld.Upload.Upload(
		ctx,
		image.ImageData,
		uploader.UploadParams{PublicID: ""},
	)
	if err != nil {
		return &uploader.UploadResult{}, err
	}

	return uploadResult, nil
}
