package main

import (
	"log"

	"github.com/yuzuy/go-guide-after-progate/server"
)

func main() {
	s := server.New()
	if err := s.Start(); err != nil {
		log.Println(err)
	}
}
