package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func defaultRouteHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Aloha, World!")
}

func loginRouteHandler(c echo.Context) error {
	vals, err := c.FormParams()
	if err != nil {
		log.Fatal("loginRouteHandler:", err)
	}
	for k, v := range vals {
		fmt.Println(k, v[0])
	}
	return err
}
