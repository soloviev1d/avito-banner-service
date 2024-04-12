package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/soloviev1d/avito-banner-service/internal/database"
)

var globStr string

func main() {
	dbUsr := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	fmt.Println(dbUsr, dbPass)
	fmt.Println("hai hello")
	s, err := database.ParseData(dbUsr, dbPass)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("here")
	globStr = s
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeGet)
	if err = http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err)
	}
}

func homeGet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s</h1>", globStr)
	fmt.Println(r.Header)
}
