package controllers

import (
	"log"
	"net/http"
	"smart_pharmacy_api/models"
	"smart_pharmacy_api/services"

	"github.com/gin-gonic/gin"
)

// ProductSearch godoc
// @Summary Get product details
// @Produce json
// @Param productName path string true "productName"
// @Success 200 {object} models.Product
// @Router /api/productsearch/{productName} [get]
func ProductSearch(c *gin.Context) {
	productName := c.Params.ByName("productName")
	var Product []models.Product
	err := services.GetProductSearch(&Product, productName)
	if err != nil || Product == nil {
		log.Printf("Not found : %v\n ", err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		log.Printf("success : %v\n ", http.StatusOK)
		c.JSON(http.StatusOK, Product)
	}
}
