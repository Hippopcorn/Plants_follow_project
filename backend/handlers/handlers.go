package handlers

import (
	"plant_project/backend/repository"
)

// s'occupe que du serveur HTTP
type Handler struct {
	repo *repository.Repository
}

func NewHandler(r *repository.Repository) *Handler {
	return &Handler{
		repo: r,
	}
}
