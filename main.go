package main

import (
	"fmt"
	"log"
	"net/http"

	"api/apis/apiserver"
)

func main() {
	http.HandleFunc("/", apiserver.Handler)
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
