package main

import (
	"github.com/ivansukach/book-service/protocol"
	"github.com/ivansukach/book-service/repositories"
	"github.com/ivansukach/book-service/server"
	"github.com/ivansukach/book-service/service"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	db, err := sqlx.Connect("postgres", "user=su password=su dbname=books sslmode=disable")
	if err != nil {
		log.Error(err)
		return
	}
	rps := repositories.New(db)
	bs := service.New(rps)
	srv := server.New(bs)
	s := grpc.NewServer()
	protocol.RegisterBookServiceServer(s, srv)
	listener, err := net.Listen("tcp", ":1323")
	if err != nil {
		log.Error(err)
		return
	}
	s.Serve(listener)
}
