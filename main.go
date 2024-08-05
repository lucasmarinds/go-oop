package main

import (
	"fmt"

	ct "github.com/marin/go-pacotes/clientes"
	c "github.com/marin/go-pacotes/contas"
)

func main() {
	marin := ct.Titular{Nome: "Lucas", CPF: "123-321-222-33", Profissao: "Eng. Software"}
	contaMarin := c.ContaCorrente{marin, 12345, 22223, 4000}
	fmt.Println(contaMarin)
}

func somaNumeros(numeros ...int) {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	fmt.Println("Soma dos numeros é", total)
}
