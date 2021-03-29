package app

import (
	"smart_pharmacy_api/controllers"
)

func mapUrls() {
	syncApis := r.Group("/api/productsearch")
	{
		syncApis.GET("/:productName", controllers.ProductSearch)
	}

	invoiceApis := r.Group("/api/retailerinvoice")
	{
		invoiceApis.GET("/:buyerId/:fromDate/:toDate", controllers.RetailerInvoice)
	}
}
