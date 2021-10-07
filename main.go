package main

import (
	"log"
	"net/http"
)

func main() {
	router := makeRoute()
	if err := http.ListenAndServe(":9000", router); err != nil {
		log.Fatalln("cannot start server", err)
	}
}
