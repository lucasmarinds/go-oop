package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	var contaGuilherme ContaCorrente = ContaCorrente{titular: "Guilherme", numeroAgencia: 0111, numeroConta: 123451, saldo: 59.89}
	// só sem precisar passar nome APENAS se instanciar o objeto com todos campos preenchidos
	contaBruna := ContaCorrente{"Bruna", 0112, 661234, 300}
	fmt.Println(contaGuilherme)
	fmt.Println(contaBruna)
}
