package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("\U0001f680 Server is running...")
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Iniciou minha request")
	defer log.Println("Finalizou minha request")

	select {
	case <-time.After(time.Second * 5):
		w.Write([]byte("\U0001f44b Requisição processada com sucesso!"))
		log.Println("\U0001f44b Requisição processada com sucesso!")
	case <-ctx.Done():
		http.Error(w, "Request cancelada", http.StatusRequestTimeout)

	}
}
