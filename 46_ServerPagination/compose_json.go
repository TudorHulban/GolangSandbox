package main

import (
	"log"
	"strconv"
)

type PostsResponse struct {
	TotalItems   int   `json:"totalposts"`
	PageNumber   int64 `json:"pagenumber"`
	ItemsPerPage int64 `json:"itemsperpage"`
	Items        []Content
}

func ComposeResponse(pPageNumber string, pItemsPage string, pCredentials *Credentials) (PostsResponse, error) {

	dbHandle, err := Connect2RDBMS(pCredentials)
	defer dbHandle.Close()

	if err != nil {
		log.Fatal("DBHandler.sql.Open:", err)
	}

	var rPostsResponse PostsResponse
	rPostsResponse.TotalItems, err = HowManyPosts(dbHandle)
	if err != nil {
		log.Fatal("ComposeResponse.HowManyPosts:", err)
	}

	rPostsResponse.PageNumber, err = strconv.ParseInt(pPageNumber, 10, 64)
	if err != nil {
		log.Fatal("rPostsResponse.PageNumber:", err)
	}

	rPostsResponse.ItemsPerPage, err = strconv.ParseInt(pItemsPage, 10, 64)
	if err != nil {
		log.Fatal("rPostsResponse.ItemsPerPage:", err)
	}

	rPostsResponse.Items, err = QueryPosts(dbHandle, rPostsResponse.PageNumber, rPostsResponse.ItemsPerPage)

	return rPostsResponse, err
}
