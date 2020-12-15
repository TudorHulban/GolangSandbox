package main

import (
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.GET("/", defaultRouteHandler)
	e.POST("/login", loginRouteHandler)

	var logLevel log.Lvl
	logLevel = 2
	e.Logger.SetLevel(logLevel)
	e.HideBanner = true
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return (customGenerator(35))
		},
	}))
	e.Logger.Debug(e.Start(":1323"))
}

func customGenerator(noDigits int) string {
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	digits := math.Max(float64(len(timestamp)+1), float64(noDigits))
	token := make([]rune, int(digits)-len(timestamp))

	for i := range token {
		token[i] = letterRunes[randomNumber(len(letterRunes))]
	}
	return string(token) + timestamp
}

func randomNumber(maximum int) int {
	if maximum < 1 {
		return 0
	}
	theNumber := rand.Intn(maximum)
	return theNumber
}
