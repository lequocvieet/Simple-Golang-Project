package main

import (
	"fmt"
	"log"
	"net/http"

	"golang_standard/middlewares"
	"golang_standard/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	router := mux.NewRouter()

	router.Handle("/auth/login", http.HandlerFunc(routes.Login)).Methods("POST")
	router.Handle("/auth/register", http.HandlerFunc(routes.Register)).Methods("POST")
	router.Handle("/get-albums-onchain", http.HandlerFunc(middlewares.CheckJwt(routes.GetOnBlockChain))).Methods("GET")
	router.Handle("/albums", http.HandlerFunc(middlewares.CheckJwt(routes.GetAllAlbums))).Methods("GET")
	router.Handle("/albums", http.HandlerFunc(middlewares.CheckJwt(routes.CreateAlbum))).Methods("POST")
	//router.DELETE("/posts/:id", middlewares.CheckJwt(routes.DeletePost))

	fmt.Println("Listening to port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
