package repository

import "gorm.io/gorm"

// le truc (layer/couche) qui permet d'acceder a la BD
type Repository struct {
	repo *gorm.DB
}

func NewRepository(r *gorm.DB) *Repository {
	return &Repository{
		repo: r,
	}
}
