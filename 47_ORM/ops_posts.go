package main

import (
	"strconv"
)

type Post struct {
	Id       int64
	Title    string
	Contents string
	AuthorId int64
}

func (b *Blog) AddPost(pPost *Post) error {
	return b.DBConn.Insert(pPost)
}

func (b *Blog) GetPost(pID int64) (Post, error) {
	result := Post{Id: pID}
	errSelect := b.DBConn.Select(&result)
	return result, errSelect
}

func (b *Blog) UpdatePost(pPost *Post) error {
	return b.DBConn.Update(pPost)
}

// GetUserPosts fetches posts for specific user, reverse order, latest first.
func (b *Blog) GetUserPosts(pUserID int64, pNo int64) ([]Post, error) {
	var result []Post
	var limitRows string
	if pNo > 0 {
		limitRows = " limit " + strconv.FormatInt(pNo, 10)
	}
	sql := "select * from posts where author_id = ? order by 1 desc " + limitRows

	_, errSelect := b.DBConn.Query(&result, sql, pUserID)
	return result, errSelect
}

// GetLatestPosts fetches last posts from all users, reverse order, latest first.
func (b *Blog) GetLatestPosts(pNo int64) ([]Post, error) {
	var result []Post
	var limitRows string
	if pNo > 0 {
		limitRows = " limit " + strconv.FormatInt(pNo, 10)
	}
	sql := "select * from posts order by 1 desc " + limitRows

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
