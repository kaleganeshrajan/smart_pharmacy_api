package models

import "time"

type InvoiceProduct struct {
	Quantity            int     `json:"Quantity"`
	SupplierProductName string  `json:"SupplierProductName"`
	Rate                float64 `json:"Rate"`
	Amount              float64 `json:"Amount"`
}

type Invoices struct {
	SupplierCode      string     `json:"SupplierCode"`
	SupplierName      string     `json:"SupplierName"`
	InvoiceNumberNo   string     `json:"InvoiceNumber_No"`
	InvoiceDate       *time.Time     `json:"Invoice_Date"`
	InvoiceAmount     int        `json:"InvoiceAmount"`
	PendingAmount     int        `json:"PendingAmount"`
	AdjustedAmount    int        `json:"AdjustedAmount"`
	DueDate           *time.Time `json:"DueDate"`
	IsModiTransaction bool       `json:"IsModiTransaction"`
	IsDeletedFlag     bool       `json:"IsDeleted_flag"`
	IsOfflineFlag     bool       `json:"IsOffline_flag"`
	ProductList       []InvoiceProduct  `json:"PrdList"`
}

type InvoiceDetails struct {
	SupplierName        string
	IsModiTransaction   bool
	SupplierId          string
	SupplierProductName string
	SupplierProductCode string
	Rate                float32
	Quantity            int
	Amount              float64
	DueDate             time.Time
	InvoiceNo           string
	InvoiceDate         time.Time
	NetInvoiceAmount    int
	PendingAmount       int
	AdjustedAmount      int
	IsOffline_flag      bool
	IsDeleted_flag      bool
}
