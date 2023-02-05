package controller

import (
	"fmt"
	"net/http"
	"shop/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateProductInput struct {
	Price           int `json:"price" binding:"required"`
	Quantity        int `json:"quantity" binding:"required"`
	CompetitorPrice int `json:"competitorPrice" binding:"required"`
}

type UpdateProductInput struct {
	id              uint `json:"id" binding:"required"`
	Price           int  `json:"price"`
	CompetitorPrice int  `json:"competitorPrice"`
	Quantity        int  `json:"quantity"`
}

func CreateProduct(c *gin.Context) {
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create product
	product := models.Product{Price: input.Price, Quantity: input.Quantity, CompetitorPrice: input.CompetitorPrice}
	models.DB.Create(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func GetProduct(c *gin.Context) {
	var prodcuts []models.Product
	id, _err := strconv.ParseUint(c.Query("id"), 10, 64)
	if _err != nil {
		fmt.Println((_err.Error()))
	}
	models.DB.Where(&models.Product{
		ID: uint(id),
	}).Order("quantity asc").Find(&prodcuts)
	c.IndentedJSON(http.StatusOK, prodcuts)
}

func DeleteProduct(c *gin.Context) {
	// Get model if exist
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// PATCH /books/:id
// Update a book
func UpdateProduct(c *gin.Context) {
	// Get model if exist
	var product models.Product
	if err := models.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": product})
}
