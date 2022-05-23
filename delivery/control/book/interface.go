package book

import "github.com/labstack/echo/v4"

type IBookController interface {
	InsertBook() echo.HandlerFunc
	SelectBook() echo.HandlerFunc
	UpdateBook() echo.HandlerFunc
	DeleteBook() echo.HandlerFunc
	SelectBookById() echo.HandlerFunc
	SelectBookByPenerbit() echo.HandlerFunc
}
