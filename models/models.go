package models

import (
	"time"
)

type Customer struct {
	CustomerID          int       `gorm:"column:customer_id;primary_key;AUTO_INCREMENT" json:"customer_id"`
	CompanyName         string    `gorm:"column:company_name" json:"company_name"`
	FirstName           string    `gorm:"column:first_name" json:"first_name"`
	LastName            string    `gorm:"column:last_name" json:"last_name"`
	BillingAddress      string    `gorm:"column:billing_address" json:"billing_address"`
	City                string    `gorm:"column:city" json:"city"`
	StateOrProvince     string    `gorm:"column:state_or_province" json:"state_or_province"`
	ZipCode             string    `gorm:"column:zip_code" json:"zip_code"`
	Email               string    `gorm:"column:email" json:"email"`
	CompanyWebsite      string    `gorm:"column:company_website" json:"company_website"`
	PhoneNumber         string    `gorm:"column:phone_number" json:"phone_number"`
	FaxNumber           string    `gorm:"column:fax_number" json:"fax_number"`
	ShipAddress         string    `gorm:"column:ship_address" json:"ship_address"`
	ShipCity            string    `gorm:"column:ship_city" json:"ship_city"`
	ShipStateOrProvince string    `gorm:"column:ship_state_or_province" json:"ship_state_or_province"`
	ShipZipCode         string    `gorm:"column:ship_zip_code" json:"ship_zip_code"`
	ShipPhoneNumber     string    `gorm:"column:ship_phone_number" json:"ship_phone_number"`
	CreatedAt           time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt           time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (Customer) TableName() string {
	return "customers"
}

type Employee struct {
	EmployeeID int       `gorm:"column:employee_id;primary_key;AUTO_INCREMENT" json:"employee_id"`
	FirstName  string    `gorm:"column:first_name" json:"first_name"`
	LastName   string    `gorm:"column:last_name" json:"last_name"`
	Title      string    `gorm:"column:title" json:"title"`
	WorkPhone  string    `gorm:"column:work_phone" json:"work_phone"`
	CreatedAt  time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt  time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (Employee) TableName() string {
	return "employees"
}

type Order struct {
	OrderID             int            `gorm:"column:order_id;primary_key;AUTO_INCREMENT" json:"order_id"`
	CustomerID          int            `gorm:"column:customer_id;NOT NULL" json:"customer_id"`
	EmployeeID          int            `gorm:"column:employee_id;NOT NULL" json:"employee_id"`
	OrderDate           time.Time      `gorm:"column:order_date" json:"order_date"`
	PurchaseOrderNumber string         `gorm:"column:purchase_order_number" json:"purchase_order_number"`
	ShipDate            time.Time      `gorm:"column:ship_date" json:"ship_date"`
	ShippingMethodID    int            `gorm:"column:shipping_method_id" json:"shipping_method_id"`
	FreightCharge       int            `gorm:"column:freight_charge" json:"freight_charge"`
	Taxes               int            `gorm:"column:taxes" json:"taxes"`
	PaymentReceived     string         `gorm:"column:payment_received" json:"payment_received"`
	Comment             string         `gorm:"column:comment" json:"comment"`
	OrderDetail         []OrderDetail  `json:"order_detail"`
	Customer            Customer       `json:"customer"`
	Employee            Employee       `json:"employee"`
	ShippingMethod      ShippingMethod `json:"shipping_method"`
	CreatedAt           time.Time      `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt           time.Time      `gorm:"column:deleted_at" json:"deleted_at"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderDetail struct {
	OrderDetailID int       `gorm:"column:order_detail_id;primary_key;AUTO_INCREMENT" json:"order_detail_id"`
	OrderID       int       `gorm:"column:order_id" json:"order_id"`
	ProductID     int       `gorm:"column:product_id" json:"product_id"`
	Quantity      int       `gorm:"column:quantity" json:"quantity"`
	UnitPrice     float64   `gorm:"column:unit_price" json:"unit_price"`
	Discount      float64   `gorm:"column:discount" json:"discount"`
	Product       Product   `json:"product"`
	CreatedAt     time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (OrderDetail) TableName() string {
	return "order_details"
}

type Product struct {
	ProductID   int       `gorm:"column:product_id;primary_key;AUTO_INCREMENT" json:"product_id"`
	ProductName string    `gorm:"column:product_name" json:"product_name"`
	UnitPrice   int       `gorm:"column:unit_price" json:"unit_price"`
	InStock     string    `gorm:"column:in_stock" json:"in_stock"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (Product) TableName() string {
	return "products"
}

type ShippingMethod struct {
	ShippingMethodID int       `gorm:"column:shipping_method_id;primary_key;AUTO_INCREMENT" json:"shipping_method_id"`
	ShippingMethod   string    `gorm:"column:shipping_method" json:"shipping_method"`
	CreatedAt        time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt        time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

func (ShippingMethod) TableName() string {
	return "shipping_methods"
}
