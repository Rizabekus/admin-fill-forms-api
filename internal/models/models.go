package models

import (
	"net/http"
)

type FormsService interface {
	SendResponse(response ResponseStructure, w http.ResponseWriter, statusCode int)
	CheckCredentials(AdminLogin AdminLogin) bool
	CreateSession(uuid string) error
	CheckSession(uuid string) (bool, error)
	AddForm(NewForm AddForm) error
}
type FormsStorage interface {
	CreateSession(uuid string) error
	CheckSession(uuid string) (bool, error)
	AddForm(NewForm AddForm) error
}

type AddForm struct {
	Project_name string `json:"project_name" validate:"omitempty,AddForm"`
	Category     string `json:"category" validate:"omitempty,AddForm"`
	Project_type string `json:"project_type" validate:"omitempty,AddForm"`
	Age_category string `json:"age_category" validate:"omitempty,AddForm"`
	Year         string `json:"year" validate:"omitempty,AddForm"`
	Timing       string `json:"timing" validate:"omitempty,AddForm"`
	Keywords     string `json:"keywords" validate:"omitempty,AddForm"`
	Summary      string `json:"summary" validate:"omitempty,AddForm"`
	Director     string `json:"director" validate:"omitempty,AddForm"`
	Producer     string `json:"producer" validate:"omitempty,AddForm"`
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
