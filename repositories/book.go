package repositories

type Book struct {
	id            string
	title         string
	author        string
	genre         string
	edition       string
	numberOfPages int32
	year          int32
	amount        int32
	isPopular     bool
	inStock       bool
}
type Repository interface {
	Create(book *Book) error
}
