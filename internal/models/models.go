package models

import (
	"net/http"
)

type FormsService interface {
	SendResponse(response ResponseStructure, w http.ResponseWriter, statusCode int)
	CheckCredentials(AdminLogin AdminLogin) bool
	CreateSession(uuid string) error
	CheckSession(uuid string) (bool, error)
}
type FormsStorage interface {
	CreateSession(uuid string) error
	CheckSession(uuid string) (bool, error)
}

type AddForm struct {
	Project_name string `json:"project_name" validate:""`
	Category     string `json:"category"`
	Project_type string `json:"project_type"`
	Age_category string `json:"age_category"`
	Year         string `json:"year"`
	Timing       string `json:"timing"`
	Keywords     string `json:"keywords"`
	Summary      string `json:"summary"`
	Director     string `json:"director"`
	Producer     string `json:"producer"`
}

type AdminLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Form struct {
}
type ResponseStructure struct {
	Field string `json:"field"`
	Error string `json:"error"`
}
