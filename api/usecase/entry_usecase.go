package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type IEntryUsecase interface {
	CreateEntry(entry model.Entry) error
}

type entryUsecase struct {
	er repository.IEntryRepository
}

func NewEntryUsecase(er repository.IEntryRepository) IEntryUsecase {
	return &entryUsecase{er}
}

func (eu *entryUsecase) CreateEntry(entry model.Entry) error {
	if err := eu.er.CreateEntry(&entry); err != nil {
		return err
	}
	return nil
}
