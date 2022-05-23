package book

import (
	"errors"
	"perpusgo/entity"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *BookRepo {
	return &BookRepo{
		Db: db,
	}
}

type BookRepo struct {
	Db *gorm.DB
}

// Func InsertBook untuk menambah buku
func (br *BookRepo) InsertBook(newBook entity.Book) (entity.Book, error) {
	if err := br.Db.Create(&newBook).Error; err != nil {
		log.Warn(err)
		return entity.Book{}, errors.New("tidak bisa insert data")
	}
	log.Info()
	return newBook, nil
}

// Func SelectBook untuk menampilkan seluruh buku
func (br *BookRepo) SelectBook() ([]entity.Book, error) {
	arrBook := []entity.Book{}

	if err := br.Db.Find(&arrBook).Error; err != nil {
		log.Warn(err)
		return nil, errors.New("tidak bisa menampilkan data seluruh buku")
	}

	if len(arrBook) == 0 {
		log.Warn("tidak ada data")
		return nil, errors.New("tidak ada data")
	}
	log.Info()
	return arrBook, nil
}

// Func UpdateBook untuk mengupdate buku sesuai ID buku
func (br *BookRepo) UpdateBook(id int, update entity.Book) (entity.Book, error) {
	var book entity.Book

	if err := br.Db.Updates(&update).Where("id = ?", id).Find(&book).Error; err != nil {
		log.Warn(err)
		return entity.Book{}, errors.New("tidak bisa update buku")
	}
	log.Info()
	return book, nil
}

// Func DeleteBook untuk menghapus buku sesuai ID buku
func (br *BookRepo) DeleteBook(id int) (entity.Book, error) {
	var book entity.Book

	if err := br.Db.Delete(&book).Where("id = ?", id).Error; err != nil {
		log.Warn(err)
		return entity.Book{}, errors.New("tidak bisa menghapus buku")
	}

	log.Info()
	return book, nil

}
