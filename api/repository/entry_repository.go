package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IEntryRepository interface {
	CreateEntry(entry *model.Entry) error
}

type entryRepository struct {
	db *gorm.DB
}

func NewEntryRepository(db *gorm.DB) IEntryRepository {
	return &entryRepository{db}
}

func (er *entryRepository) CreateEntry(entry *model.Entry) error {
	if err := er.db.Create(entry).Error; err != nil {
		return err
	}
	return nil
}
