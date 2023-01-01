package responses

type Product struct {
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	Discount    float64 `json:"discount"`
	SubTotal    float64 `json:"sub_total"`
}

type OrderDetail struct {
	OrderId        int       `json:"order_id"`
	CustomerName   string    `json:"customer_name"`
	EmployeeName   string    `json:"employee_name"`
	ShippingMethod string    `json:"shipping_method"`
	Product        []Product `json:"product"`
	TotalPayment   float64   `json:"total_payment"`
}
