package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Javlopez/volume-assessment/calculate"
	"github.com/gorilla/mux"
)

const (
	Port = "8080"
	Addr = "localhost"
)

type Api struct {
	Router *mux.Router
}

func main() {
	api := NewApi()
	serveURL := fmt.Sprintf("%s:%s", Addr, Port)
	log.Print("Api is running.....")
	log.Fatal(http.ListenAndServe(serveURL, api.Router))
}

func NewApi() *Api {
	r := mux.NewRouter()
	r.HandleFunc("/calculate", calculate.Handler).Methods(http.MethodPost)
	return &Api{Router: r}
}
