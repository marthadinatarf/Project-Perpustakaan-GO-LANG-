package book

import (
	"net/http"
	"perpusgo/delivery/views"
	req "perpusgo/delivery/views/request"
	"perpusgo/delivery/views/responses"
	"perpusgo/entity"
	bookRepo "perpusgo/repository/book"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type BookControl struct {
	Repo  bookRepo.IBookRepository
	Valid *validator.Validate
}

func New(repo bookRepo.IBookRepository, valid *validator.Validate) *BookControl {
	return &BookControl{
		Repo:  repo,
		Valid: valid,
	}
}

// Controller untuk menambah data buku
func (bc *BookControl) InsertBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpBook req.InsertBookRequest
		var resp responses.BookResponse

		if err := c.Bind(&tmpBook); err != nil {
			log.Warn("salah input")
			return c.JSON(http.StatusBadRequest, views.StatusInvalidRequest())
		}

		if err := bc.Valid.Struct(tmpBook); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotAcceptable, views.StatusValidate())
		}

		newBook := entity.Book{
			Judul:    tmpBook.Judul,
			Author:   tmpBook.Author,
			Penerbit: tmpBook.Penerbit,
		}

		data, err := bc.Repo.InsertBook(newBook)
		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, views.InternalServerError())
		}

		resp = responses.BookResponse{
			ID:       int(data.ID),
			Judul:    data.Judul,
			Author:   data.Author,
			Penerbit: data.Penerbit,
		}

		log.Info("berhasil insert data buku")
		return c.JSON(http.StatusCreated, responses.InsertBookSuccess(resp))
	}
}
