package main

import (
	"Book-API-Server/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type Book struct {
	IsBN uint `json:"isbn" gorm:"primaryKey;column:isbn"`
	BookSpec
}

type BookSpec struct {
	Title  string  `json:"title" gorm:"column:title;type:varchar(200)"`
	Author string  `json:"author" gorm:"column:author;type:varchar(200);index"`
	Price  float64 `json:"price" gorm:"column:price"`
	IsSale *bool   `json:"is_sale" gorm:"column:is_sale"`
}

func (t *Book) TableName() string {
	return "books"
}

func Failed(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{"code": 0, "error": err.Error()})
}

func main() {
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "config/application.yaml"
	}
	if err := config.LoadConfigFromYaml(path); err != nil {
		fmt.Printf("load config err: %s\n", err)
		os.Exit(1)
	}

	conf := config.Get()

	server := gin.Default()

	db := conf.MySQL.DB()

	book := server.Group("api/books")

	book.POST("", func(ctx *gin.Context) {
		ins := new(Book)
		if err := ctx.ShouldBindJSON(ins); err != nil {
			Failed(ctx, err)
			return
		}

		if err := db.Save(ins).Error; err != nil {
			Failed(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"code": 0, "data": ins})
	})

	book.GET("", func(ctx *gin.Context) {
		var books []Book
		if err := db.Find(&books).Error; err != nil {
			Failed(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, books)
	})

	book.GET("/:isbn", func(ctx *gin.Context) {
		var ins Book
		id := ctx.Param("isbn")
		if err := db.Where("id = ?", id).First(&ins).Error; err != nil {
			Failed(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, ins)
	})

	book.PUT("/:isbn", func(ctx *gin.Context) {
		id := ctx.Param("isbn")

		req := BookSpec{}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			Failed(ctx, err)
			return
		}

		if err := db.Where("id = ?", id).Model(&Book{}).Updates(&req).Error; err != nil {
			Failed(ctx, err)
			return
		}

		var ins Book
		if err := db.Where("id = ?", id).Take(&ins).Error; err != nil {
			Failed(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, ins)
	})

	book.DELETE("/:isbn", func(ctx *gin.Context) {
		id := ctx.Param("isbn")
		if err := db.Where("id = ?", id).Delete(&Book{}).Error; err != nil {
			Failed(ctx, err)
			return
		}
	})

	if err := server.Run(conf.App.Address()); err != nil {
		log.Println(err)
	}
}
