package services

import (
	"Rizabekus/admin-fill-forms-api/internal/models"
	"encoding/json"
	"net/http"
	"os"
)

type FormsService struct {
	storage models.FormsStorage
}

func CreateFormsService(storage models.FormsStorage) *FormsService {
	return &FormsService{storage: storage}
}
func (fs *FormsService) SendResponse(response models.ResponseStructure, w http.ResponseWriter, statusCode int) {
	responseJSON, err := json.Marshal(response)
	if err != nil {

		internalError := models.ResponseStructure{
			Field: "Internal Server Error",
			Error: "Failed to marshal JSON response",
		}
		internalErrorJSON, _ := json.Marshal(internalError)

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, string(internalErrorJSON), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
}

func (fs *FormsService) CheckCredentials(AdminLogin models.AdminLogin) bool {
	if AdminLogin.Login == os.Getenv("ADMIN_USERNAME") && AdminLogin.Password == os.Getenv("ADMIN_PASSWORD") {
		return true
	}
	return false
}
func (fs *FormsService) CreateSession(uuid string) error {
	return fs.storage.CreateSession(uuid)
}

func (fs *FormsService) CheckSession(uuid string) (bool, error) {
	return fs.storage.CheckSession(uuid)
}
func (fs *FormsService) AddForm(NewForm models.AddForm) error {
	return fs.storage.AddForm(NewForm)
}
