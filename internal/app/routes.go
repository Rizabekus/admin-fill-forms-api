package app

import (
	"log"
	"net/http"

	"Rizabekus/admin-fill-forms-api/internal/handlers"
	"Rizabekus/admin-fill-forms-api/pkg/loggers"

	"github.com/gorilla/mux"
)

func Routes(h *handlers.Handlers) {
	r := mux.NewRouter()

	r.HandleFunc("/adminlogin", h.AdminLogin).Methods("POST")
	r.HandleFunc("/addform", h.AddForm).Methods("POST")
	// r.HandleFunc("/login", h.Login).Methods("POST")
	// r.HandleFunc("/modify", h.Modify).Methods("POST")

	loggers.InfoLog.Println("Started the server at http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
