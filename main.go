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
	contaBruna2 := ContaCorrente{"Bruna", 0112, 661234, 300}
	fmt.Println(contaGuilherme)
	fmt.Println(contaBruna)
	fmt.Println(contaBruna == contaBruna2) // Aqui é true porque quando não sao intanciadas com o New, o if apenas checa os valores dos objetos

	var contaMarin *ContaCorrente
	contaMarin = new(ContaCorrente)
	contaMarin.titular = "Marin"
	contaMarin.saldo = 30.30
	var contaMarin2 *ContaCorrente
	contaMarin = new(ContaCorrente)
	contaMarin.titular = "Marin"
	contaMarin.saldo = 30.30
	fmt.Println(*contaMarin)               // &{valores} < o & só aparece porque ele está dizendo "Dentro dssa memoria tem esse valor" se colocarmos um ponteiro na conta marin apontando para o tipo dele, mostrara da maneira correta sem o &
	fmt.Println(contaMarin == contaMarin2) // aqui é false, porque como sao 2 variaveis que foram instanciadas apartir de um ponteiro apontando para conta corrente
}
