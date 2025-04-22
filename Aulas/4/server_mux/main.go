package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", &Blog{})
	mux2 := http.NewServeMux()
	mux2.Handle("/", &User{})

	go func() {
		if err := http.ListenAndServe(":8080", mux); err != nil {
			panic(err)
		}
	}()

	if err := http.ListenAndServe(":8081", mux2); err != nil {
		log.Fatalf("Erro ao iniciar servidor 8081: %v", err)
	}
}

// func homeHandler(w http.ResponseWriter, r *http.Request) {

// 	w.Write([]byte("Hello World!"))

// }

type Blog struct{}

func (b *Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(" Blog!"))
}

type User struct{}

func (u *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(" User!"))
}
