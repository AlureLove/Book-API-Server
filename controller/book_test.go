package controller_test

import (
	"Book-API-Server/controller"
	"Book-API-Server/model"
	"context"
	"testing"
)

func TestGetBook(t *testing.T) {
	book := controller.NewBookController()
	ins, err := book.GetBook(context.Background(), &controller.GetBookRequest{Isbn: 1})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestCreateBook(t *testing.T) {
	book := controller.NewBookController()
	ins, err := book.CreateBook(context.Background(), &model.BookSpec{
		Title:  "Go Programming Language",
		Author: "Robert Griesemer",
		Price:  1000,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
