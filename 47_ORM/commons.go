package main

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Blog struct {
	DBConn *pg.DB
	Users  []User
	Posts  []Post
}

func (b *Blog) CreateTable4Model(pModel interface{}) error {
	return b.DBConn.CreateTable(pModel, &orm.CreateTableOptions{Temp: false, IfNotExists: true})
}

func NewBlog(pDB *pg.DB) (*Blog, error) {
	result := new(Blog)
	result.DBConn = pDB

	errUsers := result.CreateTable4Model(interface{}(&User{}))
	if errUsers != nil {
		return nil, errUsers
	}
	errPosts := result.CreateTable4Model(interface{}(&Post{}))
	if errPosts != nil {
		return nil, errPosts
	}
	return result, nil
}
