package main

import (
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
)

func test() {
	const a = 5
	const b = 6
	fmt.Println(a * b)
	fmt.Fscan(io.MultiReader())

	//const hoje = date.Brasilia

	agora := time.Now()

	id := uuid.New()

	fmt.Println(id)
	fmt.Println(agora)

}
