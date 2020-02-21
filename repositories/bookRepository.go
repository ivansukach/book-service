package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func New(db *sqlx.DB) Repository {
	return &bookRepository{db: db}
}

type bookRepository struct {
	db *sqlx.DB
}

func (br *bookRepository) Create(book *Book) error {
	_, err := br.db.NamedExec("INSERT INTO book VALUES (:id, :title, :author, :genre, :edition, :numberofpages, :year, :amount, :ispopular, :instock)", book)
	return err
}
func (br *bookRepository) Read(id string) (*Book, error) {
	u := Book{}
	err := br.db.QueryRowx("SELECT * FROM book WHERE id=$1", id).StructScan(&u)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &u, err
}
func (br *bookRepository) Update(book *Book) error {
	_, err := br.db.NamedExec("UPDATE book SET Title=:title, Author=:author, Genre=:genre, Edition=:edition, "+
		"NumberOfPages=:numberofpages, Year=:year, Amount=:amount, IsPopular=:ispopular, InStock=:instock WHERE Id=:id", book)
	return err
}
func (br *bookRepository) Delete(id string) error {
	_, err := br.db.Exec("DELETE FROM book WHERE id=$1", id)
	return err
}
func (br *bookRepository) Listing() ([]Book, error) {
	rows, err := br.db.Queryx("SELECT Id, Title, Author, Genre, Edition, NumberOfPages, Year, Amount, IsPopular, InStock FROM book")
	if err != nil {
		log.Warning(err)
		return nil, err
	}
	b := make([]Book, 0)
	for rows.Next() {
		book := new(Book)
		err = rows.StructScan(&book)
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
