package model

import (
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/v2/tools/pretty"
)

var (
	v = validator.New()
)

type Book struct {
	IsBN uint `json:"isbn" gorm:"primaryKey;column:isbn"`
	BookSpec
}

type BookSpec struct {
	Title  string  `json:"title" gorm:"column:title;type:varchar(200)" validate:"required"`
	Author string  `json:"author" gorm:"column:author;type:varchar(200);index" validate:"required"`
	Price  float64 `json:"price" gorm:"column:price;" validate:"required"`
	IsSale *bool   `json:"is_sale" gorm:"column:is_sale"`
}

func (r *BookSpec) Validate() error {
	return v.Struct(r)
}

func (b *Book) TableName() string {
	return "books"
}

func (b *Book) String() string {
	return pretty.ToJSON(b)
}

func (s *BookSpec) String() string {
	return pretty.ToJSON(s)
}
