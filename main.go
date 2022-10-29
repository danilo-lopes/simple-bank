package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/danilo-lopes/simple-bank/src/config"
	"github.com/danilo-lopes/simple-bank/src/router"
)

func main() {
	r := router.Generate()
	config.Load()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.APIPort),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("Simple Bank API Serving on Port %d\n", config.APIPort)
	log.Fatal(s.ListenAndServe())
}
