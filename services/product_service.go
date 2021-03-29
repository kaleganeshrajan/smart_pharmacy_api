package services

import (
	"smart_pharmacy_api/models"

	db "github.com/brkelkar/common_utils/databases"
)

func GetProductSearch(Product *[]models.Product, productName string) (err error) {
	err = db.DB["smartdb"].Select(`ProductCode,
	ProductName,
		MRP,
		PTR,
		PTS`).Where(
		" ProductName LIKE '" + productName + "%'").Find(&Product).Error
	return
}
