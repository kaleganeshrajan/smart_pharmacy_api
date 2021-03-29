package services

import (
	"fmt"
	"smart_pharmacy_api/models"

	db "github.com/brkelkar/common_utils/databases"
)

func GetRetailerInvoice(Invoice *[]models.Invoices, buyerID, fromDate, toDate string) (err error) {
	var invoicedetails []*models.InvoiceDetails
	err = db.DB["smartdb"].Raw("exec SP_RETAILER_INVOICE  ?, ?, ?", buyerID, fromDate, toDate).Scan(&invoicedetails).Error
	if err != nil {
		fmt.Printf("Buyer Dashboard: %V\n", err)
	}

	orderMap := make(map[string]models.Invoices, len(invoicedetails))
	for _, val := range invoicedetails {
		if _, ok := orderMap[val.SupplierId+"_"+val.InvoiceNo]; !ok {
			//fill order details
			var tempInvoice models.Invoices
			tempInvoice.SupplierCode = val.SupplierId
			tempInvoice.SupplierName = val.SupplierName
			tempInvoice.InvoiceNumberNo = val.InvoiceNo
			tempInvoice.AdjustedAmount = val.AdjustedAmount
			tempInvoice.DueDate = &val.DueDate
			tempInvoice.InvoiceAmount = val.NetInvoiceAmount
			tempInvoice.InvoiceDate = &val.InvoiceDate
			tempInvoice.IsDeletedFlag = val.IsDeleted_flag
			tempInvoice.IsModiTransaction = val.IsModiTransaction
			tempInvoice.IsOfflineFlag = val.IsOffline_flag

			orderMap[val.SupplierId+"_"+val.InvoiceNo] = tempInvoice
		}

		o := orderMap[val.SupplierId+"_"+val.InvoiceNo]
		var product models.InvoiceProduct
		product.Amount = val.Amount
		product.Quantity = val.Quantity
		product.Rate = float64(val.Rate)
		product.SupplierProductName = val.SupplierName
		o.ProductList = append(o.ProductList, product)
		orderMap[val.SupplierId+"_"+val.InvoiceNo] = o
	}

	for _, mapVal := range orderMap {
		temp := mapVal
		*Invoice = append(*Invoice, temp)
	}

	return
}
