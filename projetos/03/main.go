package main

import (
	"chi/books"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/hello", hello)
	println("\U0001f680 server is running...")
	http.ListenAndServe(":8080", r)

}

func hello(w http.ResponseWriter, r *http.Request) {
	book := books.Book{
		Name:     "Senhor dos Aneis",
		Quantity: 112,
		Price:    145.56,
	}
	fmt.Println(book)
	//book.Imprimir()
	//book.Bookstore("Livro1", 28, 37.67)

	bookGold := books.BookGold{
		Book: books.Book{
			Name:     "lIVRO GOLD",
			Quantity: 245,
			Price:    780.15,
		},
		Desconto: 15,
	}

	jsonBook := `{"nome":"lIVRO GOLD","quantidade":245,"preco":780.15,"desconto":15}`

	booGold2 := books.BookGold{}

	json.Unmarshal([]byte(jsonBook), &booGold2)

	booGold2.Imprimir()

	bookGold.Imprimir()

	b, _ := json.Marshal(bookGold)
	w.Write([]byte(b))
}
