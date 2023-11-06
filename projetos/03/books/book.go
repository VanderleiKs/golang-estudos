package books

import (
	_ "encoding/json"
	"fmt"
)

type Book struct {
	Name     string  `json:"nome"`
	Quantity int     `json:"quantidade"`
	Price    float64 `json:"preco"`
}

type BookGold struct {
	Book
	Desconto int `json:"desconto"`
}

func (b *Book) Bookstore(name string, quantity int, price float64) {
	b.Name = name
	b.Quantity = quantity
	b.Price = price
}

func (b BookGold) Imprimir() {
	fmt.Printf("%v\n", b)
}
