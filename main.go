package main

import (
	"fmt"
	"log"
	"net/http"

	"api/apis/apiserver"
)

func main() {
	dir := http.FileServer(http.Dir("./apis/apiserver/css"))
	http.Handle("/css/", http.StripPrefix("/css/", dir))
	http.HandleFunc("/", apiserver.Handler)
	http.HandleFunc("/artist", apiserver.HandleArtist)
	http.HandleFunc("/locs/", apiserver.HandleLocations)
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
