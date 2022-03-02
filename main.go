package main

import (
	"github.com/labstack/echo/v4"
	"github.com/onimarufl/test-q-change.git/service/question2"
)

func main() {

	question2 := question2.NewHandler(
		question2.New(),
	)

	e := echo.New()
	e.POST("/cashier", question2.Cashier)

	e.Logger.Fatal(e.Start(":1323"))
}
