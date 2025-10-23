package handlers

import (
	"plant_project/repository"
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
