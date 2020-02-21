package server

import (
	"context"
	"github.com/ivansukach/book-service/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"testing"
)

func TestUpdate(t *testing.T) {
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(":1323", opts)
	if err != nil {
		log.Error(err)
		return
	}
	defer conn.Close()

	client := protocol.NewBookServiceClient(conn)
	book := new(protocol.Book)
	book.Id = "book1582286213"
	book.Title = "World and Piece"
	book.Author = "Lev Tolstoy"
	book.Genre = "Romance"
	book.Amount = 2500
	book.Year = 1954
	book.NumberOfPages = 320
	book.Edition = "Moscow-1996"
	book.IsPopular = false
	book.InStock = false
	//_, err = client.Update(context.Background(), &protocol.UpdateRequest{Book: book})
	_, err = client.Update(context.Background(), &protocol.UpdateRequest{Book: book})
	if err != nil {
		log.Error(err)
	}

}
