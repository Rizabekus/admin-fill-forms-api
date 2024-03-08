package handlers

import "Rizabekus/admin-fill-forms-api/internal/services"

type Handlers struct {
	Service *services.Services
}

func HandlersInstance(services *services.Services) *Handlers {
	return &Handlers{Service: services}
}
