package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"jatis.com/test/rest/config"
	"jatis.com/test/rest/handler"
)

func main() {
	db, err := config.NewDB()

	logFatal(err)

	e := echo.New()

	h := &handler.Handler{DB: db}
	e.GET("order/:order_id", h.GetOrderDetails)

	e.Logger.Fatal(e.Start(":9000"))
}

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
