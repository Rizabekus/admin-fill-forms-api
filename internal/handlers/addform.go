package handlers

import (
	"Rizabekus/admin-fill-forms-api/internal/models"
	"Rizabekus/admin-fill-forms-api/pkg/loggers"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
)

func (handler *Handlers) AddForm(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("mobydev_api_admin_session")
	if err != nil {
		if err == http.ErrNoCookie {
			response := models.ResponseStructure{
				Field: "You are not logged in",
				Error: "No permission to modify",
			}
			handler.Service.FormsService.SendResponse(response, w, http.StatusInternalServerError)
			return

		} else {
			response := models.ResponseStructure{
				Field: "Internal Server Error",
				Error: err.Error(),
			}
			handler.Service.FormsService.SendResponse(response, w, http.StatusInternalServerError)

			return
		}
	}
	exists, err := handler.Service.FormsService.CheckSession(cookie.Value)
	if err != nil {
		response := models.ResponseStructure{
			Field: "Internal Server Error",
			Error: err.Error(),
		}
		handler.Service.FormsService.SendResponse(response, w, http.StatusInternalServerError)
		return

	}
	if !exists {
		response := models.ResponseStructure{
			Field: "You are not logged in",
			Error: "No permission to modify",
		}
		handler.Service.FormsService.SendResponse(response, w, http.StatusInternalServerError)
		return
	}
	var NewForm models.AddForm
	err = json.NewDecoder(r.Body).Decode(&NewForm)
	if err != nil {
		response := models.ResponseStructure{
			Field: "Failed to decode JSON",
			Error: err.Error(),
		}
		handler.Service.FormsService.SendResponse(response, w, http.StatusBadRequest)

		loggers.InfoLog.Println("Failed to decode JSON")
		return
	}
	loggers.DebugLog.Println("Received data in JSON format")
	validate := validator.New()
	validate.RegisterValidation("cyrillicOrLatinAndAscii", validators.validateCyrillicOrLatinAndAscii)
	validate.Struct(NewForm)
}
