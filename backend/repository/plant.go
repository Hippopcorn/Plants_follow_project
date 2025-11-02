package repository

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Plant struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	Comment   string    `db:"comment"`
}

func (r *Repository) ListPlants() []Plant {
	plants, err := gorm.G[Plant](r.repo).Raw("SELECT id, name, comment, created_at FROM plants").Find(context.Background())
	if err != nil {
		panic(err)
	}
	return plants
}

func (r *Repository) GetPlantByID(id int) Plant {
	plant, err := gorm.G[Plant](r.repo).Raw("SELECT id, name, comment, created_at FROM plants WHERE id = ?", id).Take(context.Background())
	if err != nil {
		panic(err)
	}
	return plant
}

type CreatePlant struct {
	Name    string `db:"name"`
	Comment string `db:"comment"`
}

func (r *Repository) CreatePlant(plant CreatePlant) {
	row := gorm.G[CreatePlant](r.repo).Raw("INSERT INTO plants (name, comment) VALUES (?, ?)", plant.Name, plant.Comment).Row(context.Background())
	if row.Err() != nil {
		panic(row.Err())
	}
}

func (r *Repository) DeletePlant(id int) {
	row := gorm.G[Plant](r.repo).Raw("DELETE FROM plants WHERE id = ?", id).Row(context.Background())
	if row.Err() != nil {
		panic(row.Err())
	}
}

type UpdatePlant struct {
	Name    string `db:"name"`
	Comment string `db:"comment"`
}

func (r *Repository) UpdatePlant(plant UpdatePlant, id int) {
	row := gorm.G[Plant](r.repo).Raw(`
	UPDATE plants 
	SET name = ?,
		comment = ?
	WHERE id = ?
	`, plant.Name, plant.Comment, id).Row(context.Background())
	if row.Err() != nil {
		panic(row.Err())
	}
}