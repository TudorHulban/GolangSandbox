package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	endpointAuthors = "/authors"
	endpointPosts   = "/posts"
)

var _blog *Blog
var _db *gorm.DB

func main() {
	var errOpen error
	_db, errOpen = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if errOpen != nil {
		log.Fatalf("could not connect to database")
	}

	var errCreate error
	_blog, errCreate = NewBlog(_db)
	if errCreate != nil {
		log.Fatalf("could not create blog: %s", errCreate)
	}

	e := echo.New()
	e.HideBanner = true
	e.Static("/", "assets")
	//e.GET("/", hLanding)
	e.POST(endpointAuthors, saveUser)
	e.GET(endpointAuthors, getAuthors)
	e.GET(endpointAuthors+"/:id", getAuthor)
	e.GET(endpointAuthors+"/:id"+endpointPosts+"/:no", getPosts)
	e.Logger.Fatal(e.Start(":1323"))
}

func hLanding(c echo.Context) error {
	return c.String(http.StatusOK, "Landing page ...")
}

func saveUser(c echo.Context) error {
	u := &Author{
		Name:   c.FormValue("name"),
		Emails: strings.Split(c.FormValue("email"), ";"),
	}

	errAdd := _blog.AddAuthor(u)
	if errAdd != nil {
		return c.String(http.StatusInternalServerError, errAdd.Error())
	}

	return c.String(http.StatusOK, "OK")
}

func getPosts(c echo.Context) error {
	authorID, errParse := strconv.ParseInt(c.Param("id"), 10, 64)
	if errParse != nil {
		return c.String(http.StatusBadRequest, "Bad user ID "+c.Param("id"))
	}

	_, errGet := _blog.GetAuthor(authorID)
	if errGet != nil {
		return c.String(http.StatusNotFound, "Author ID "+c.Param("id")+" not found.")
	}

	noPosts, errParse := strconv.ParseInt(c.Param("no"), 10, 64)
	if errParse != nil {
		return c.String(http.StatusBadRequest, "Bad number of posts "+c.Param("no"))
	}

	posts, errGetPosts := _blog.GetPosts(authorID, noPosts)
	if errGetPosts != nil {
		return c.String(http.StatusInternalServerError, errGetPosts.Error())
	}

	var result string
	for _, v := range posts {
		result = result + "," + v.Title
	}
	return c.String(http.StatusOK, result[1:])
}

func getAuthors(c echo.Context) error {
	authors, errGetUsers := _blog.GetAllAuthors()
	if errGetUsers != nil {
		return c.String(http.StatusInternalServerError, errGetUsers.Error())
	}
	if len(authors) == 0 {
		return c.String(http.StatusNotFound, "No users.")
	}

	var result string
	for _, v := range authors {
		result = result + "," + v.Name
	}
	return c.String(http.StatusOK, result[1:])
}

func getAuthor(c echo.Context) error {
	authorID, errParse := strconv.ParseInt(c.Param("id"), 10, 64)
	if errParse != nil {
		return c.String(http.StatusBadRequest, "Bad user ID "+c.Param("id"))
	}

	author, errGet := _blog.GetAuthor(authorID)
	if errGet != nil {
		return c.String(http.StatusNotFound, "Author ID "+c.Param("id")+" not found.")
	}
	return c.String(http.StatusOK, author.Name)
}
