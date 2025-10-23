package main

import (
	"plant_project/handlers"
	"plant_project/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//dsn : infos de la base de donnee
	dsn := "host=localhost user=myuser password=mypassword dbname=mydatabase port=5434 sslmode=disable"
	// ouverture de la gormDB (creer une connection)
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// sqlDB sert a pouvoir close la connexion a la DB
	sqlDB, err := gormDB.DB()
	if err != nil {
		panic(err)
	}
	// pour fermer la connexion avec la DB quand le programme est terminé
	defer sqlDB.Close()

	// enregistrer la connection dans une structure Handler
	handler := handlers.NewHandler(repository.NewRepository(gormDB))

	// Create a Gin router HTTP (our server) with default middleware pour creer des routes d'API
	r := gin.Default()

	r.POST("/plants", handler.CreatePlant) //adresse de la fonction

	r.GET("/plants", handler.GetListPlants)
	r.GET("/plants/:id", handler.GetPlantByID)

	r.DELETE("/plants/:id", handler.DeletePlant)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}
