package main

import (
	"strconv"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Author Structure consolidating user that writes blogs information.
type Author struct {
	Id     int64
	Name   string
	Emails []string
}

type Post struct {
	Id              int64
	AuthorId        int64
	CreatedTimeUnix int64
	LastUpdateUnix  int64
	Title           string
	Contents        string
}

type Blog struct {
	postsPerPage uint
	DBConn       *pg.DB
}

func NewBlog(db *pg.DB) (*Blog, error) {
	result := &Blog{
		postsPerPage: 5,
		DBConn:       db,
	}

	errUsers := result.CreateTable4Model(interface{}(&Author{}))
	if errUsers != nil {
		return nil, errUsers
	}
	errPosts := result.CreateTable4Model(interface{}(&Post{}))
	if errPosts != nil {
		return nil, errPosts
	}
	return result, nil
}

func (b *Blog) CreateTable4Model(model interface{}) error {
	return b.DBConn.CreateTable(model, &orm.CreateTableOptions{Temp: false, IfNotExists: true})
}

func (b *Blog) AddAuthor(a *Author) error {
	return b.DBConn.Insert(a)
}

func (b *Blog) GetAuthor(id int64) (Author, error) {
	result := Author{Id: id}

	errSelect := b.DBConn.Select(&result)
	return result, errSelect
}

func (b *Blog) UpdateAuthor(a *Author) error {
	return b.DBConn.Update(a)
}

func (b *Blog) GetAllAuthors() ([]Author, error) {
	var result []Author

	errSelect := b.DBConn.Model(&result).Select()
	return result, errSelect
}

func (b *Blog) GetMaxIDUsers() (int64, error) {
	var maxID struct {
		Max int64
	}

	_, errQuery := b.DBConn.QueryOne(&maxID, "select max(id) from authors")
	return maxID.Max, errQuery
}

func (b *Blog) AddPost(p *Post) error {
	return b.DBConn.Insert(p)
}

func (b *Blog) GetPost(id int64) (Post, error) {
	result := Post{
		Id: id,
	}

	errSelect := b.DBConn.Select(&result)
	return result, errSelect
}

func (b *Blog) UpdatePost(p *Post) error {
	return b.DBConn.Update(p)
}

// GetUserPosts fetches posts for specific user, reverse order, latest first.
func (b *Blog) GetPosts(authorID, noPosts int64) ([]Post, error) {
	var result []Post
	var sql string

	if noPosts > 0 {
		sql = "select * from posts where author_id = ? order by 1 desc limit " + strconv.FormatInt(noPosts, 10)
	} else {
		sql = "select * from posts where author_id = ? order by 1 desc limit " + strconv.FormatInt(int64(b.postsPerPage), 10)
	}

	_, errSelect := b.DBConn.Query(&result, sql, authorID)
	return result, errSelect
}

// GetLatestPosts fetches last posts from all authors, reverse order, latest first.
func (b *Blog) GetLatestPosts(noPosts int64) ([]Post, error) {
	var result []Post
	var sql string

	if noPosts > 0 {
		sql = "select * from posts order by 1 desc limit " + strconv.FormatInt(noPosts, 10)
	} else {
		sql = "select * from posts order by 1 desc limit " + strconv.FormatInt(int64(b.postsPerPage), 10)
	}

	_, errSelect := b.DBConn.Query(&result, sql, "")
	return result, errSelect
}

func (b *Blog) GetMaxIDPosts() (int64, error) {
	var maxID struct {
		Max int64
	}
	_, errQuery := b.DBConn.QueryOne(&maxID, "select max(id) from posts")
	return maxID.Max, errQuery
}
