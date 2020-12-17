package main

import (
	"strconv"

	"gorm.io/gorm"
)

// Author Structure consolidating user that writes blogs information.
type Author struct {
	gorm.Model
	Id     int64
	Name   string
	Emails []string
}

type Post struct {
	gorm.Model
	Id              int64
	AuthorId        int64
	CreatedTimeUnix int64
	LastUpdateUnix  int64
	Title           string
	Contents        string
}

type Blog struct {
	postsPerPage uint
	DBConn       *gorm.DB
}

func NewBlog(db *gorm.DB) (*Blog, error) {
	result := &Blog{
		postsPerPage: 5,
		DBConn:       db,
	}

	errAuthors := db.AutoMigrate(&Author{})
	if errAuthors != nil {
		return nil, errAuthors
	}
	errPosts := db.AutoMigrate(&Post{})
	if errPosts != nil {
		return nil, errPosts
	}
	return result, nil
}

func (b *Blog) AddAuthor(a *Author) error {
	return b.DBConn.Create(a).Error
}

func (b *Blog) GetAuthor(id int64) (Author, error) {
	result := Author{Id: id}
	return result, b.DBConn.Select(&result).Error
}

func (b *Blog) UpdateAuthor(a *Author) error {
	return b.DBConn.Model(a).Updates(*a).Error
}

func (b *Blog) GetAllAuthors() ([]Author, error) {
	var result []Author

	return result, b.DBConn.Find(&result).Error
}

func (b *Blog) GetMaxIDUsers() (int64, error) {
	var maxID struct {
		Max int64
	}

	return maxID.Max, b.DBConn.Raw("select max(id) from authors").Scan(&maxID).Error
}

func (b *Blog) AddPost(p *Post) error {
	return b.DBConn.Create(p).Error
}

func (b *Blog) GetPost(id int64) (Post, error) {
	result := Post{
		Id: id,
	}

	return result, b.DBConn.Select(&result).Error
}

func (b *Blog) UpdatePost(p *Post) error {
	return b.DBConn.Model(p).Updates(*p).Error
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

	return result, b.DBConn.Raw(sql, authorID).Scan(&result).Error
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

	return result, b.DBConn.Raw(sql, "").Scan(&result).Error
}

func (b *Blog) GetMaxIDPosts() (int64, error) {
	var maxID struct {
		Max int64
	}

	return maxID.Max, b.DBConn.Raw("select max(id) from posts", "").Scan(&maxID).Error
}
