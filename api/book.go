package api

import (
	"Book-API-Server/controller"
	"Book-API-Server/model"
	"Book-API-Server/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BookApiHandler struct {
	svc *controller.BookController
}

func NewBookApiHandler() *BookApiHandler {
	return &BookApiHandler{
		svc: controller.NewBookController(),
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
	req := new(model.BookSpec)
	if err := ctx.ShouldBindJSON(req); err != nil {
		response.Failed(ctx, err)
		return
	}

	ins, err := h.svc.CreateBook(ctx.Request.Context(), req)
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, ins)
}

func (h *BookApiHandler) ListBooks(ctx *gin.Context) {
	books, err := h.svc.ListBooks(ctx.Request.Context())
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, books)
}

func (h *BookApiHandler) GetBook(ctx *gin.Context) {
	strId := ctx.Param("isbn")
	id, err := strconv.ParseUint(strId, 10, 0)
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	ins, err := h.svc.GetBook(ctx.Request.Context(), &controller.GetBookRequest{Isbn: uint(id)})
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, ins)
}

func (h *BookApiHandler) UpdateBook(ctx *gin.Context) {
	strId := ctx.Param("isbn")
	id, err := strconv.ParseUint(strId, 10, 0)
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	req := new(model.BookSpec)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Failed(ctx, err)
		return
	}

	ins, err := h.svc.UpdateBook(ctx.Request.Context(), uint(id), req)
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, ins)
}

func (h *BookApiHandler) DeleteBook(ctx *gin.Context) {
	strId := ctx.Param("isbn")
	id, err := strconv.ParseUint(strId, 10, 0)
	if err != nil {
		response.Failed(ctx, err)
		return
	}

	if err = h.svc.DeleteBook(ctx.Request.Context(), &controller.GetBookRequest{Isbn: uint(id)}); err != nil {
		response.Failed(ctx, err)
		return
	}

	response.Success(ctx, nil)
}
