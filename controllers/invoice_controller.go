package controllers

import (
	"log"
	"net/http"
	"smart_pharmacy_api/models"
	"smart_pharmacy_api/services"

	"github.com/gin-gonic/gin"
)

// RetailerInvoice godoc
// @Summary Get retailer invoice
// @Produce json
// @Param buyerID path string true "buyerId"
// @Param fromDate path string true "fromDate"
// @Param toDate path string true "toDate"
// @Success 200 {object} models.Product
// @Router /api/retailerinvoice/{buyerID}/{fromDate}/{toDate} [get]
func RetailerInvoice(c *gin.Context) {
	buyerId := c.Params.ByName("buyerId")
	fromDate := c.Params.ByName("fromDate")
	toDate := c.Params.ByName("toDate")
	var Invoice []models.Invoices
	err := services.GetRetailerInvoice(&Invoice, buyerId, fromDate, toDate)
	if err != nil || Invoice == nil {
		log.Printf("Not found : %v\n ", err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		log.Printf("success : %v\n ", http.StatusOK)
		c.JSON(http.StatusOK, Invoice)
	}
}
