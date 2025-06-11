package api

import (
	"Book-API-Server/config"
	"Book-API-Server/model"
	"Book-API-Server/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookApiHandler struct {
	db *gorm.DB
}

func NewBookApiHandler() *BookApiHandler {
	return &BookApiHandler{
		db: config.Get().MySQL.DB(),
	}
}

func (h *BookApiHandler) Registry(r *gin.Engine) {
	book := r.Group("api/books")
	book.POST("", h.CreateBook)
	book.GET("", h.ListBooks)
	book.GET("/:isbn", h.GetBook)
	book.PUT("/:isbn", h.UpdateBook)
	book.DELETE("/:isbn", h.DeleteBook)
}

func (h *BookApiHandler) CreateBook(ctx *gin.Context) {
	ins := new(model.Book)
	if err := ctx.ShouldBindJSON(ins); err != nil {
		response.Failed(ctx, err)
		return
	}

	if err := h.db.Save(ins).Error; err != nil {
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, ins)
}

func (h *BookApiHandler) ListBooks(ctx *gin.Context) {
	var books []model.Book
	if err := h.db.Find(&books).Error; err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, books)
}

func (h *BookApiHandler) GetBook(ctx *gin.Context) {
	var ins model.Book
	id := ctx.Param("isbn")

	if err := h.db.Where("isbn = ?", id).Take(&ins).Error; err != nil {
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, ins)
}

func (h *BookApiHandler) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("isbn")

	req := model.BookSpec{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Failed(ctx, err)
		return
	}

	if err := h.db.Where("isbn = ?", id).Model(&model.Book{}).Updates(&req).Error; err != nil {
		response.Failed(ctx, err)
		return
	}

	var ins model.Book
	if err := h.db.Where("isbn = ?", id).Take(&ins).Error; err != nil {
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, ins)
}

func (h *BookApiHandler) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("isbn")
	if err := h.db.Where("isbn = ?", id).Delete(&model.Book{}).Error; err != nil {
		response.Failed(ctx, err)
		return
	}
}
