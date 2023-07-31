package main

import (
	"context"
	"github.com/hxcuber/hollywood-api/api/internal/api/router"
	"github.com/hxcuber/hollywood-api/api/pkg/httpserv"
)

func main() {
	r := router.Handler()
	c := context.Background()
	httpserv.NewServer(r).Start(c)
}
