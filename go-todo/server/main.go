package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JakubPluta/go-projects/go-todo/router"
)

func main() {
	r := router.Router()
	fmt.Println("starting server on port: 9000")
	log.Fatal(http.ListenAndServe(":9000", r))
}
