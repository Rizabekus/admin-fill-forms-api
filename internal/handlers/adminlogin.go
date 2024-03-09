package handlers

import (
	"Rizabekus/admin-fill-forms-api/internal/models"
	"Rizabekus/admin-fill-forms-api/pkg/loggers"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

func (handler *Handlers) AdminLogin(w http.ResponseWriter, r *http.Request) {
	loggers.DebugLog.Println("Received a request to AdminLogin")
	var AdminLogin models.AdminLogin
	err := json.NewDecoder(r.Body).Decode(&AdminLogin)
	if err != nil {
		response := models.ResponseStructure{
			Field: "Failed to decode JSON",
			Error: "",
		}
		handler.Service.FormsService.SendResponse(response, w, http.StatusBadRequest)

		loggers.InfoLog.Println("Failed to decode JSON")
		return
	}
	loggers.DebugLog.Println("Received data from JSON")
	ok := handler.Service.FormsService.CheckCredentials(AdminLogin)
	if !ok {
		response := models.ResponseStructure{
			Field: "Wrong Credentials",
			Error: "Login/Password is wrong",
		}
		handler.Service.FormsService.SendResponse(response, w, http.StatusBadRequest)

		loggers.DebugLog.Println("Failed to authenticate")
		return
	}
	loggers.DebugLog.Println("Checked  for credentials")
	u2, err := uuid.NewV4()
	if err != nil {

		response := models.ResponseStructure{
			Field: "Internal Server Error",
			Error: "",
		}
		handler.Service.FormsService.SendResponse(response, w, http.StatusInternalServerError)

		loggers.InfoLog.Println("Failed to create UUID")
		return
	}
	loggers.DebugLog.Println("Created UUID instance")
	err = handler.Service.FormsService.CreateSession(u2.String())
	if err != nil {

		response := models.ResponseStructure{
			Field: "Internal Server Error",
			Error: "",
		}
		handler.Service.FormsService.SendResponse(response, w, http.StatusInternalServerError)

		loggers.InfoLog.Println("Failed to insert UUID into database")
		return
	}
	loggers.DebugLog.Println("Created a session")
	cookie := &http.Cookie{Name: "mobydev_api_admin_session", Value: u2.String(), Expires: time.Now().Add(365 * 24 * time.Hour)}
	http.SetCookie(w, cookie)
	loggers.DebugLog.Println("Successfuly logged in")
}
