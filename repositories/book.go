package repositories

type Book struct {
	Id            string
	Title         string
	Author        string
	Genre         string
	Edition       string
	NumberOfPages int32
	Year          int32
	Amount        int32
	IsPopular     bool
	InStock       bool
}
type Repository interface {
	Create(book *Book) error
	Read(id string) (*Book, error)
	Update(book *Book) error
	Delete(id string) error
	Listing() ([]Book, error)
}
