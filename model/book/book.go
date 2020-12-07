package book

import (
	"time"

	"github.com/lz1998/ecust_library/model"
)

type EcustBook struct {
	ID          int64     `gorm:"column:id" json:"id" form:"id"`
	Author      string    `gorm:"column:author" json:"author" form:"author"`
	Title       string    `gorm:"column:title" json:"title" form:"title"`
	Press       string    `gorm:"column:press" json:"press" form:"press"`
	Year        int32     `gorm:"column:year" json:"year" form:"year"`
	BookId      string    `gorm:"column:book_id" json:"book_id" form:"book_id"`
	Isbn        string    `gorm:"column:isbn" json:"isbn" form:"isbn"`
	Institution string    `gorm:"column:institution" json:"institution" form:"institution"`
	Status      int32     `gorm:"column:status" json:"status" form:"status"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`
}

func init() {
	if err := model.Db.AutoMigrate(&EcustBook{}); err != nil {
		panic(err)
	}
}

func CreateBook(author string, title string, press string, year int32, bookId string, isbn string, institution string) error {
	book := &EcustBook{
		Author:      author,
		Title:       title,
		Press:       press,
		Year:        year,
		BookId:      bookId,
		Isbn:        isbn,
		Institution: institution,
		Status:      0,
	}
	return model.Db.Save(book).Error
}

func ListBook(author []string, title []string, press []string, startYear int32, endYear int32, bookId []string, isbn []string, institution []string) ([]*EcustBook, error) {
	q := model.Db.Model(&EcustBook{})
	if len(author) != 0 {
		q = q.Where("author in ?", author)
	}
	if len(title) != 0 {
		q = q.Where("title in ?", title)
	}
	if len(press) != 0 {
		q = q.Where("press in ?", press)
	}
	if startYear != 0 {
		q = q.Where("year > ?", startYear)
	}
	if endYear != 0 {
		q = q.Where("year < ?", endYear)
	}
	if len(bookId) != 0 {
		q = q.Where("book_id in ?", bookId)
	}
	if len(isbn) != 0 {
		q = q.Where("isbn in ?", isbn)
	}
	if len(institution) != 0 {
		q = q.Where("institution in ?", institution)
	}
	q = q.Where("status = 0")
	var books []*EcustBook
	if err := q.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func UpdateBook(id int64, author string, title string, press string, year int32, bookId string, isbn string, institution string, status int32) error {
	var book EcustBook
	if err := model.Db.Model(&EcustBook{}).Where("id = ?", id).First(&book).Error; err != nil {
		return err
	}

	book.Author = author
	book.Title = title
	book.Press = press
	book.Year = year
	book.BookId = bookId
	book.Isbn = isbn
	book.Institution = institution
	book.Status = status

	return model.Db.Save(&book).Error
}
