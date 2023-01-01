package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"jatis.com/test/rest/models"
	"jatis.com/test/rest/models/responses"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	Handler struct {
		DB *gorm.DB
	}
	Response struct {
		ErrorCode int         `json:"error_code"`
		Message   string      `json:"message"`
		Data      interface{} `json:"data"`
	}
)

func (h *Handler) GetOrderDetails(c echo.Context) (err error) {
	orderId := c.Param("order_id")
	response := new(Response)
	order := models.Order{}
	order.OrderID, _ = strconv.Atoi(orderId)
	h.DB.Preload("OrderDetail.Product").Preload(clause.Associations).Find(&order)

	orderDetail := responses.OrderDetail{
		OrderId:        order.OrderID,
		CustomerName:   order.Customer.CompanyName,
		EmployeeName:   fmt.Sprintf("%s %s", order.Employee.FirstName, order.Employee.LastName),
		ShippingMethod: order.ShippingMethod.ShippingMethod,
	}
	var products []responses.Product
	for _, p := range order.OrderDetail {
		product := responses.Product{
			ProductName: p.Product.ProductName,
			Quantity:    p.Quantity,
			UnitPrice:   p.UnitPrice,
			Discount:    p.Discount,
			SubTotal:    (p.UnitPrice * float64(p.Quantity)) - p.Discount,
		}
		orderDetail.TotalPayment += product.SubTotal
		products = append(products, product)
	}

	orderDetail.Product = products

	response.Data = orderDetail
	return c.JSON(http.StatusOK, response)
}
