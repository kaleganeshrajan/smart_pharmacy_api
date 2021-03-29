package models

// Product is entity
type Product struct {
	ProductCode string  `json:"productCode"`
	ProductName string  `json:"productName"`
	MRP         float32 `json:"mrp"`
	PTR         float32 `json:"ptr"`
	PTS         float32 `json:"pts"`
}

//TableName retunrs source table name
func (Product) TableName() string {
	return "dbo.BulkMapping_GCP"
}
