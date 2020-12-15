package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.HideBanner = true

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	//public routes
	e.Static("/", "../../public")
	e.POST("/login", handleLoginReq)
	e.GET("/", handlePublicRoutes)

	//restricted routes
	r := e.Group("/a")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", handleRestrictedRoutes)

	e.Logger.Fatal(e.Start(":3000"))
}

func handlePublicRoutes(c echo.Context) error {
	return c.String(http.StatusOK, "Public Route")
}

func handleRestrictedRoutes(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func handleLoginReq(c echo.Context) error {
	u := c.FormValue("username")
	p := c.FormValue("password")

	v, err := ValidateCredentials(u, p)
	if err != nil {
		return echo.ErrInternalServerError
	}

	if v.ID == -1 {
		return echo.ErrUnauthorized
	}

	token, err := createToken(v.FirstName+" "+v.LastName, v.IsAdmin, 3600)
	return c.JSON(http.StatusOK, token)
}

func createToken(pFullName string, pIsAdmin bool, pTimeInterval int) (map[string]string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = pFullName
	claims["admin"] = strconv.FormatBool(pIsAdmin)
	claims["exp"] = time.Now().Add(time.Second * time.Duration(pTimeInterval)).Unix()

	j, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return map[string]string{"token": j}, nil
}
