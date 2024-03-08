package services

import (
	"Rizabekus/admin-fill-forms-api/internal/models"
	"Rizabekus/admin-fill-forms-api/internal/storage"
)

type Services struct {
	FormsService models.FormsService
}

func ServiceInstance(storage *storage.Storage) *Services {
	return &Services{
		FormsService: CreateFormsService(storage.FormsStorage),
	}
}
