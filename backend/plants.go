package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Plant struct {
	ID   int
	Name string `json:"name" binding:"required"`
}

func (h *Handler) GetListPlants(c *gin.Context) {
	plantsDB := h.repo.ListPlants()
	plants := make([]Plant, len(plantsDB))
	// pour controler ce qui est retourné (ex : si on rajoute un password dans la structure):
	for i := range len(plantsDB) {
		plants[i] = Plant{
			ID:   plantsDB[i].ID,
			Name: plantsDB[i].Name,
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
		ID:   plantDB.ID,
		Name: plantDB.Name,
	}

	c.JSON(http.StatusOK, plant)
}

func (h *Handler) CreatePlant(c *gin.Context) {
	var newPlant Plant
	err := c.ShouldBind(&newPlant) // parse le payload (--data '{"Name":"Cactus"}') et le met dans newPlant
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})

		return
	}
	h.repo.CreatePlant(newPlant.Name)

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
