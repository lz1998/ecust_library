package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lz1998/ecust_library/dto"
	"github.com/lz1998/ecust_library/model/book"
)

func CreateBook(c *gin.Context) {
	req := &dto.CreateBookReq{}

	if err := c.Bind(req); err != nil {
		c.String(http.StatusBadRequest, "bad request, not protobuf")
		return
	}
	books := req.GetBooks()
	if books == nil {
		c.String(http.StatusBadRequest, "bad request, book is nil")
		return
	}

	for _, b := range books {
		if err := book.CreateBook(b.Author, b.Title, b.Press, b.Year, b.BookId, b.Isbn, b.Institution); err != nil {
			c.String(http.StatusInternalServerError, "db error")
			return
		}
	}
	resp := &dto.CreateBookResp{}
	Return(c, resp)
}

func ListBook(c *gin.Context) {
	req := &dto.ListBookReq{}

	if err := c.Bind(req); err != nil {
		c.String(http.StatusBadRequest, "bad request, not protobuf")
		return
	}
	books, err := book.ListBook(req.Author, req.Title, req.Press, req.StartYear, req.EndYear, req.BookId, req.Isbn, req.Institution)
	if err != nil {
		c.String(http.StatusInternalServerError, "db error")
		return
	}

	resp := &dto.ListBookResp{
		Books: convertBooksModelToProto(books),
	}
	Return(c, resp)
}

func UpdateBook(c *gin.Context) {
	req := &dto.UpdateBookReq{}

	if err := c.Bind(req); err != nil {
		c.String(http.StatusBadRequest, "bad request, not protobuf")
		return
	}

	for _, b := range req.Book {
		if err := book.UpdateBook(b.Id, b.Author, b.Title, b.Press, b.Year, b.BookId, b.Isbn, b.Institution, b.Status); err != nil {
			c.String(http.StatusInternalServerError, "db error")
			return
		}
	}
	resp := &dto.UpdateBookResp{}
	Return(c, resp)
}

func convertBookModelToProto(modelBook *book.EcustBook) *dto.EcustBook {
	return &dto.EcustBook{
		Id:          modelBook.ID,
		Author:      modelBook.Author,
		Title:       modelBook.Title,
		Press:       modelBook.Author,
		Year:        modelBook.Year,
		BookId:      modelBook.BookId,
		Isbn:        modelBook.Isbn,
		Institution: modelBook.Institution,
		Status:      modelBook.Status,
		CreatedAt:   modelBook.CreatedAt.Unix(),
		UpdatedAt:   modelBook.UpdatedAt.Unix(),
	}
}

func convertBooksModelToProto(modelBooks []*book.EcustBook) []*dto.EcustBook {
	books := make([]*dto.EcustBook, 0)
	for _, modelBook := range modelBooks {
		books = append(books, convertBookModelToProto(modelBook))
	}
	return books
}
