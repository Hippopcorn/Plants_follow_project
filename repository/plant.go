package repository

import (
	"context"

	"gorm.io/gorm"
)

type Plant struct {
	ID   int
	Name string
}

func (r *Repository) ListPlants() []Plant {
	plants, err := gorm.G[Plant](r.repo).Raw("SELECT id, name FROM plants").Find(context.Background())
	if err != nil {
		panic(err)
	}
	return plants
}

func (r *Repository) GetPlantByID(id int) Plant {
	plant, err := gorm.G[Plant](r.repo).Raw("SELECT id, name FROM plants WHERE id = ?", id).Take(context.Background())
	if err != nil {
		panic(err)
	}
	return plant
}

func (r *Repository) CreatePlant(name string) {
	row := gorm.G[Plant](r.repo).Raw("INSERT INTO plants (name) VALUES (?)", name).Row(context.Background())
	if row.Err() != nil {
		panic(row.Err())
	}
}

func (r *Repository) DeletePlant(id int) {
	row := gorm.G[Plant](r.repo).Raw("DELETE FROM plants WHERE id = ?", id).Row(context.Background())
	if row.Err() != nil {
		panic(row.Err())
	} else if row.Rows
}