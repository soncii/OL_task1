package service

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"login/config"
	"login/model"
	"login/storage"
	"testing"
)

//Test suite
//Integration Test Book creation, retrieval, and deletion
func TestBookService_Create(t *testing.T) {
	//Initialization
	cfg, err := config.NewConfig()
	fmt.Println(cfg)
	fmt.Println(err)
	assert.NoError(t, err, nil)
	st, err := storage.NewStorage(cfg)
	assert.Equal(t, err, nil)
	ctx := context.Background()
	service := NewBookService(st)
	//Testing create
	reqCreate := model.BookCreateReq{Title: "test", Author: "Damir"}
	resp, err := service.Create(ctx, reqCreate)
	assert.Equal(t, err, nil)
	bid := resp.BID
	reqGet := model.BookGetByIDReq{BID: bid}

	//Testing GET and if created item is saved in DB
	book, err := service.GetByID(ctx, reqGet)
	assert.Equal(t, err, nil)
	assert.Equal(t, book.Title, reqCreate.Title)
	assert.Equal(t, book.Author, reqCreate.Author)
	assert.NotEmpty(t, book.ID)

	//Testing DELETE
	deleted, err := service.DeleteByID(ctx, reqGet)
	assert.Equal(t, err, nil)
	assert.True(t, deleted)

	//Testing if the book was deleted
	book, err = service.GetByID(ctx, reqGet)
	assert.NotEqual(t, err, nil)
}
