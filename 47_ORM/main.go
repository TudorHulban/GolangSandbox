package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	endpointAuthors = "/xxx"
	endpointPosts   = "/posts"
)

var _blog *Blog
var _db *gorm.DB

func main() {
	var errOpen error
	_db, errOpen = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if errOpen != nil {
		log.Fatalf("could not connect to database")
	}

	var errCreate error
	_blog, errCreate = NewBlog(_db)
	if errCreate != nil {
		log.Fatalf("could not create blog: %s", errCreate)
	}

	webServer := fiber.New()

	webServer.Static("/", "assets")
	//e.GET("/", hLanding)
	webServer.Post(endpointAuthors, saveUser)
	webServer.Get(endpointAuthors, getAuthors)
	webServer.Get(endpointAuthors+"/:id", getAuthor)
	webServer.Get(endpointAuthors+"/:id"+endpointPosts+"/:no", getPosts)

	log.Println("Starting server...")
	webServer.Listen(":8080")
}

func hLanding(c *fiber.Ctx) error {
	return c.SendString("Landing page ...")
}

func saveUser(c *fiber.Ctx) error {
	u := &Author{
		Name:   c.FormValue("name"),
		Emails: c.FormValue("email"),
	}

	errAdd := _blog.AddAuthor(u)
	if errAdd != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusOK)
}

func getPosts(c *fiber.Ctx) error {
	authorID, errParse := strconv.ParseInt(c.FormValue("id"), 10, 64)
	if errParse != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	_, errGet := _blog.GetAuthor(uint(authorID))
	if errGet != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	noPosts, errParse := strconv.ParseInt(c.FormValue("no"), 10, 64)
	if errParse != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	posts, errGetPosts := _blog.GetPosts(authorID, noPosts)
	if errGetPosts != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	var result string
	for _, v := range posts {
		result = result + "," + v.Title
	}
	return c.SendString(result[1:])
}

func getAuthors(c *fiber.Ctx) error {
	authors, errGetUsers := _blog.GetAllAuthors()
	if errGetUsers != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error:", errGetUsers)
	}
	if len(authors) == 0 {
		log.Println("xxx")
		return c.Status(http.StatusNotFound).SendString("no authors found")
	}

	var result string
	for _, v := range authors {
		result = result + "," + v.Name
	}
	return c.SendString(result[1:])
}

func getAuthor(c *fiber.Ctx) error {
	authorID, errParse := strconv.ParseInt(c.FormValue("id"), 10, 64)
	if errParse != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	author, errGet := _blog.GetAuthor(uint(authorID))
	if errGet != nil {
		return c.SendStatus(http.StatusNotFound)
	}
	return c.SendString(author.Name)
}
