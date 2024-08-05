package main

import (
	"fmt"

	ct "github.com/marin/go-pacotes/clientes"
	c "github.com/marin/go-pacotes/contas"
)

func pagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64)
}

func main() {
	marin := ct.Titular{Nome: "Lucas", CPF: "123-321-222-33", Profissao: "Eng. Software"}
	contaMarin := c.ContaCorrente{Titular: marin, Agencia: 1234, Conta: 331001}
	fmt.Println(contaMarin)
	contaMarin.Depositar(100)
	fmt.Println(contaMarin.GetSaldo())
	pagarBoleto(&contaMarin, 50)
	fmt.Println("Boleto pago e sobrou na conta do marin", contaMarin.GetSaldo())
	fmt.Println("----------------------------")
	andrade := ct.Titular{Nome: "Andrade", CPF: "123-321-222-33", Profissao: "Eng. Software"}
	contaAndrade := c.ContaPoupanca{Titular: andrade, Agencia: 6070, Conta: 123345}
	contaAndrade.Depositar(100)
	fmt.Println(contaAndrade)
	pagarBoleto(&contaAndrade, 50)
	fmt.Println(contaAndrade.GetSaldo())
}

func somaNumeros(numeros ...int) {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	fmt.Println("Soma dos numeros é", total)
}
