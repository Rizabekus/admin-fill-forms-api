package storage

import (
	"Rizabekus/admin-fill-forms-api/internal/models"
	"database/sql"
)

type Storage struct {
	FormsStorage models.FormsStorage
}

func StorageInstance(db *sql.DB) *Storage {
	return &Storage{FormsStorage: CreateFormsStorage(db)}
}
