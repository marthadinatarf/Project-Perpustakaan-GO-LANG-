package book

import (
	"net/http"
	"perpusgo/delivery/views"
	req "perpusgo/delivery/views/request"
	"perpusgo/delivery/views/responses"
	"perpusgo/entity"
	bookRepo "perpusgo/repository/book"
	"strconv"

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

// Controller untuk menampilkan seluruh data buku
func (bc *BookControl) SelectBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := bc.Repo.SelectBook()
		if err != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, views.InternalServerError())
		}
		var arrBook []responses.BookResponse
		for _, x := range res {
			book := responses.BookResponse{
				ID:       int(x.ID),
				Judul:    x.Judul,
				Author:   x.Author,
				Penerbit: x.Penerbit,
			}
			arrBook = append(arrBook, book)
		}
		log.Info("berhasil menampilkan buku")
		return c.JSON(http.StatusOK, views.StatusOK("Success Get Data", arrBook))
	}
}

// Controller untuk mengupdate data buku sesuai ID
func (bc *BookControl) UpdateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmpUpdate req.UpdateBookRequest

		if err := c.Bind(&tmpUpdate); err != nil {
			return c.JSON(http.StatusUnsupportedMediaType, views.StatusBindData())
		}

		if err := bc.Valid.Struct(tmpUpdate); err != nil {
			log.Warn(err.Error())
			return c.JSON(http.StatusNotAcceptable, views.StatusValidate())
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotAcceptable, views.StatusIdConversion())
		}

		updateBook := entity.Book{
			Judul:    tmpUpdate.Judul,
			Author:   tmpUpdate.Author,
			Penerbit: tmpUpdate.Penerbit,
		}

		book, err := bc.Repo.UpdateBook(int(id), updateBook)

		if err != nil {
			log.Warn(err)
			//notFound := "data tidak ditemukan"
			if err != nil {
				return c.JSON(http.StatusNotFound, views.StatusNotFound("data tidak ditemukan"))
			}
			return c.JSON(http.StatusInternalServerError, views.InternalServerError())
		}
		resp := responses.BookResponse{
			ID:       int(book.ID),
			Judul:    book.Judul,
			Author:   book.Author,
			Penerbit: book.Penerbit,
		}
		return c.JSON(http.StatusOK, views.StatusUpdate(resp))
	}
}

// Controller untuk menghapus data buku sesuai ID
func (bc *BookControl) DeleteBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		idConv, err := strconv.Atoi(id)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusNotAcceptable, views.StatusIdConversion())
		}
		deleted, err := bc.Repo.SelectBookById(idConv)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, views.StatusNotFound("data tidak ditemukan"))
		}

		_, error := bc.Repo.DeleteBook(int(deleted.ID))
		if error != nil {
			log.Warn("masalah pada server")
			return c.JSON(http.StatusInternalServerError, views.InternalServerError())
		}
		return c.JSON(http.StatusOK, views.StatusDelete())
	}
}

// Controller untuk menampilkan data buku berdasarkan ID
func (bc *BookControl) SelectBookById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		idConv, err := strconv.Atoi(id)
		if err != nil {
			log.Error(err)
			return c.JSON(http.StatusNotAcceptable, views.StatusIdConversion())
		}

		book, err := bc.Repo.SelectBookById(int(idConv))
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, views.StatusNotFound("data tidak ditemukan"))
		}
		resp := responses.BookResponse{
			ID:       int(book.ID),
			Judul:    book.Judul,
			Author:   book.Author,
			Penerbit: book.Penerbit,
		}
		return c.JSON(http.StatusOK, views.StatusGetDatIdOK(resp))
	}
}

// Controller untuk menampilkan data buku berdasarkan penerbit
func (bc *BookControl) SelectBookByPenerbit() echo.HandlerFunc {
	return func(c echo.Context) error {
		Penerbit := c.Param("penerbit")

		book, err := bc.Repo.SelectBookByPenerbit(Penerbit)
		if err != nil {
			log.Warn(err)
			return c.JSON(http.StatusNotFound, views.StatusNotFound("data tidak ditemukan"))
		}
		var arrBook []responses.BookResponse
		for _, x := range book {
			book := responses.BookResponse{
				ID:       int(x.ID),
				Judul:    x.Judul,
				Author:   x.Author,
				Penerbit: x.Penerbit,
			}
			arrBook = append(arrBook, book)
		}
		log.Info("berhasil menampilkan buku berdasarkan penerbit")
		return c.JSON(http.StatusOK, views.StatusOK("Success Get Data", arrBook))
	}
}
