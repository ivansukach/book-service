package service

import "github.com/ivansukach/book-service/repositories"

type BookService struct {
	r repositories.Repository
}

func (bs *BookService) Create(book *repositories.Book) error {
	return bs.Create(book)
}

func New(repo repositories.Repository) *BookService {
	return &BookService{r: repo}
}
