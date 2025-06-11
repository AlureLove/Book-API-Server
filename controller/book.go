package controller

import (
	"Book-API-Server/config"
	"Book-API-Server/exception"
	"Book-API-Server/model"
	"context"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type BookController struct {
	db     *gorm.DB
	logger *zerolog.Logger
}

type GetBookRequest struct {
	Isbn uint
}

func NewBookController() *BookController {
	return &BookController{
		db:     config.Get().MySQL.DB(),
		logger: config.Get().Log.Logger(),
	}
}

func (c *BookController) CreateBook(ctx context.Context, req *model.BookSpec) (*model.Book, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	ins := &model.Book{
		BookSpec: *req,
	}
	if err := c.db.WithContext(ctx).Save(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (c *BookController) ListBooks(ctx context.Context) ([]*model.Book, error) {
	var books []*model.Book
	if err := c.db.WithContext(ctx).Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (c *BookController) GetBook(ctx context.Context, req *GetBookRequest) (*model.Book, error) {
	ins := &model.Book{}
	if err := c.db.WithContext(ctx).Where("isbn = ?", req.Isbn).Take(ins).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.ErrNotFound("%d not found", req.Isbn)
		}
		return nil, err
	}

	return ins, nil
}

func (c *BookController) UpdateBook(ctx context.Context, isbn uint, req *model.BookSpec) (*model.Book, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if err := c.db.WithContext(ctx).Where("isbn = ?", isbn).Model(&model.Book{}).Updates(req).Error; err != nil {
		return nil, err
	}

	ins := &model.Book{}
	if err := c.db.WithContext(ctx).Where("isbn = ?", isbn).Take(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (c *BookController) DeleteBook(ctx context.Context, req *GetBookRequest) error {
	if err := c.db.WithContext(ctx).Where("isbn = ?", req.Isbn).Delete(&model.Book{}).Error; err != nil {
		return err
	}

	return nil
}
