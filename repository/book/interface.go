package book

import "perpusgo/entity"

type IBookRepository interface {
	InsertBook(newBook entity.Book) (entity.Book, error)
	SelectBook() ([]entity.Book, error)
	UpdateBook(id int, update entity.Book) (entity.Book, error)
	DeleteBook(id int) (entity.Book, error)
	SelectBookById(id int) (entity.Book, error)
	SelectBookByPenerbit(penerbit string) ([]entity.Book, error)
}
