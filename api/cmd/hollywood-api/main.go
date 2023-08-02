package main

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/api/router"
	"github.com/hxcuber/hollywood-api/api/pkg/httpserv"
	"log"
)

func main() {
	r := router.Handler()
	c := context.Background()
	s := httpserv.NewServer(r)
	err := s.Start(c)
	if err != nil {
		log.Fatal(err)
	}
}
