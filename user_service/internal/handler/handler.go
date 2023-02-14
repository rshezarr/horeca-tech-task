package handler

import (
	"user_svc/internal/usecase"

	"github.com/go-chi/chi"
)

type Handler struct {
	router  *chi.Mux
	service *usecase.Usecase
}

func NewHandler(usecase *usecase.Usecase) *Handler {
	return &Handler{
		router:  chi.NewRouter(),
		service: usecase,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	h.router.Post("/create-user", nil)

	h.router.Get("/get-user/{email}", nil)

	return h.router
}
