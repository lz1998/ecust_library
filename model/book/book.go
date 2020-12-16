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

func ListBook(offset int, count int, authors []string, titles []string, presses []string, startYear int32, endYear int32, bookIds []string, isbns []string, institutions []string) ([]*EcustBook, int64, error) {
	q := model.Db.Model(&EcustBook{})
	if len(authors) != 0 {
		for _, author := range authors {
			q = q.Where("author like ?", "%"+author+"%")
		}
	}
	if len(titles) != 0 {
		for _, title := range titles {
			q = q.Where("title like ?", "%"+title+"%")
		}
	}
	if len(presses) != 0 {
		for _, press := range presses {
			q = q.Where("press like ?", "%"+press+"%")
		}
	}
	if len(bookIds) != 0 {
		for _, bookId := range bookIds {
			q = q.Where("book_id like ?", "%"+bookId+"%")
		}
	}
	if len(isbns) != 0 {
		for _, isbn := range isbns {
			q = q.Where("isbn like ?", "%"+isbn+"%")
		}
	}
	if len(institutions) != 0 {
		for _, institution := range institutions {
			q = q.Where("institution like ?", "%"+institution+"%")
		}
	}
	if startYear != 0 {
		q = q.Where("year > ?", startYear)
	}
	if endYear != 0 {
		q = q.Where("year < ?", endYear)
	}
	q = q.Where("status = 0")

	// 计算总数
	var total int64
	q.Count(&total)

	q = q.Order("id").Offset(offset).Limit(count)
	var books []*EcustBook
	if err := q.Find(&books).Error; err != nil {
		return nil, 0, err
	}
	return books, total, nil
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
