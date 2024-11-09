package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type IImageUsecase interface {
	UploadImage(image model.Image) (model.ImageResponse, error)
}

type imageUsecase struct {
	ir repository.IImageRepository
}

func NewImageUsecase(ir repository.IImageRepository) IImageUsecase {
	return &imageUsecase{ir}
}

func (iu *imageUsecase) UploadImage(image model.Image) (model.ImageResponse, error) {
	uploadResult, err := iu.ir.UploadImage(&image)
	if err != nil {
		return model.ImageResponse{}, err
	}
	resImageUrl := model.ImageResponse{
		ImageUrl: uploadResult.URL,
	}
	return resImageUrl, nil
}
