package handlers

import (
	"fmt"
	"net/http"
	"plant_project/backend/repository"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Plant struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	Comment   string    `json:"comment"`
}

func (h *Handler) GetListPlants(c *gin.Context) {
	plantsDB := h.repo.ListPlants()
	plants := make([]Plant, len(plantsDB))
	// pour controler ce qui est retourné (ex : si on rajoute un password dans la structure):
	for i := range len(plantsDB) {
		plants[i] = Plant{
			ID:        plantsDB[i].ID,
			Name:      plantsDB[i].Name,
			Comment:   plantsDB[i].Comment,
			CreatedAt: plantsDB[i].CreatedAt,
		}
	}
	c.JSON(http.StatusOK, plants) //renvoi un objet sous forme de JSON
}

func (h *Handler) GetPlantByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	plantDB := h.repo.GetPlantByID(id)
	plant := Plant{
		ID:        plantDB.ID,
		Name:      plantDB.Name,
		Comment:   plantDB.Comment,
		CreatedAt: plantDB.CreatedAt,
	}

	c.JSON(http.StatusOK, plant)
}

type CreatePlant struct {
	Name    string `json:"name" binding:"required"`
	Comment string `json:"comment"`
}

func (h *Handler) CreatePlant(c *gin.Context) {
	var newPlant CreatePlant
	err := c.ShouldBind(&newPlant) // parse le payload (--data '{"Name":"Cactus"}') et le met dans newPlant
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})

		return
	}
	// creer un type createPlantInput avec juste les champs configurables par l'utilisateur
	createPlantInput := repository.CreatePlant{
		Name:    newPlant.Name,
		Comment: newPlant.Comment,
	}
	h.repo.CreatePlant(createPlantInput)

	c.Status(http.StatusCreated)
}

func (h *Handler) DeletePlant(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	h.repo.DeletePlant(id)

	c.Status(http.StatusOK)
}
