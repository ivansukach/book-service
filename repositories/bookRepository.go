package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func New(db *sql.DB) Repository {
	return &bookRepository{db: db}
}

type bookRepository struct {
	db *sql.DB
}

func (br *bookRepository) Create(book *Book) error {
	_, err := br.db.Query("INSERT INTO book VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", book.Id, book.Title,
		book.Author, book.Genre, book.Edition, book.NumberOfPages, book.Year, book.Amount, book.IsPopular, book.InStock)
	return err
}
func (br *bookRepository) Read(id string) (*Book, error) {
	rows, err := br.db.Query("SELECT * FROM book WHERE id=$1", id)
	if err != nil {
		log.Warning(err)
		return nil, err
	}
	u := Book{}
	for rows.Next() {
		err = rows.Scan(&u.Id, &u.Title, &u.Author, &u.Genre, &u.Edition, &u.NumberOfPages, &u.Year, &u.Amount, &u.IsPopular, &u.InStock)
		if err != nil {
			log.Error(err)
			return nil, err
		}
	}
	return &u, err
}
func (br *bookRepository) Update(book *Book) error {
	_, err := br.db.Query("UPDATE book SET title=$1, author=$2, genre=$3, edition=$4, numberofpages=$5, year=$6, amount=$7, ispopular=$8, instock=$9 WHERE id=$10",
		book.Title,
		book.Author, book.Genre, book.Edition, book.NumberOfPages, book.Year, book.Amount, book.IsPopular, book.InStock, book.Id)
	return err
}
func (br *bookRepository) Delete(id string) error {
	_, err := br.db.Query("DELETE FROM book WHERE id=$1", id)
	return err
}
func (br *bookRepository) Listing() ([]Book, error) {
	rows, err := br.db.Query("SELECT Id, Title, Author, Genre, Edition, NumberOfPages, Year, Amount, IsPopular, InStock FROM book")
	if err != nil {
		log.Warning(err)
		return nil, err
	}
	b := make([]Book, 0)
	for i := 0; rows.Next(); i++ {
		book := new(Book)
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Genre, &book.Edition, &book.NumberOfPages,
			&book.Year, &book.Amount, &book.IsPopular, &book.InStock)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		fmt.Println("ID book: ", book.Id)
		b = append(b, *book)
		fmt.Println("ID b:", b[len(b)-1].Id)
	}
	return b, err
}
