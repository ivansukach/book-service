package repositories

import "database/sql"

func New(db *sql.DB) Repository {
	return &bookRepository{db: db}
}

type bookRepository struct {
	db *sql.DB
}

func (br *bookRepository) Create(book *Book) error {
	_, err := br.db.Query("INSERT INTO books VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", book.id, book.title,
		book.author, book.genre, book.edition, book.numberOfPages, book.year, book.amount, book.isPopular, book.inStock)
	return err
}
