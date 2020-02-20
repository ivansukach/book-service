package server

import (
	"fmt"
	"github.com/ivansukach/book-service/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"testing"
	"time"
)

func TestDelete(t *testing.T) {
	opts := grpc.WithInsecure()
	conn, err := grpc.Dial(":1323", opts)
	if err != nil {
		log.Error(err)
		return
	}
	defer conn.Close()

	client := protocol.NewBookServiceClient(conn)
	book := new(protocol.Book)
	book.Id = fmt.Sprintf("book%d", time.Now().Unix())
	book.Title = "Captain Blood"
	book.Author = "Rafael Sabatini"
	book.Genre = "Realistic Fiction"
	book.Amount = 2500
	book.Year = 1991
	book.NumberOfPages = 320
	book.Edition = "White Flow"
	book.IsPopular = false
	book.InStock = false
	_, err = client.Add(context.Background(), &protocol.AddRequest{Book: book})
	log.Error(err)
}
