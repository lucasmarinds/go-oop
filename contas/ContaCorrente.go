package contas

import (
	"fmt"

	"github.com/marin/go-pacotes/clientes"
)

/*
Quando deixamos o nome dos atributos do objeto(struct) em minusculo colocamos como se fosse "private" na visibilidade e o mesmo serve para Métodos
*/
type ContaCorrente struct {
	Titular        clientes.Titular
	Agencia, Conta int
	saldo          float64 // Deixei o saldo com primeira letra minuscula pq estou encapsulando ele, logo ele é private e só pode ser acessado por get e setters que sao meus métodos da conta que mexem com saldo
}

/*
Esses c *ContaCorrente, estamos informando que
a conta que chamar esse método e só variaveis de referencia para contacorrente
podem chamar, a que chamar vai executar ela usando os valores dela,
seria o this.saldo em java, estamos dizendo que é o OBJETO que está chamando
para alterar o atributo apenas dele.
*/
func (c *ContaCorrente) Sacar(valor float64) {
	if c.saldo < valor || valor < 0 {
		fmt.Println("Não é possivel sacar este valor")
	} else {
		c.saldo -= valor
	}
}

func (c *ContaCorrente) Depositar(valor float64) {
	if valor < 0 {
		fmt.Println("Não é possivel depositar um valor negativo")
	} else {
		c.saldo += valor
		fmt.Println("Valor depositado com sucesso!", c.saldo)
	}
}

func (c *ContaCorrente) Transferir(valor float64, destino *ContaCorrente) (string, float64) {
	if c.saldo >= valor && valor > 0 {
		c.saldo -= valor
		destino.saldo += valor
		return "Transferencia realizada com sucesso no valor de", valor
	} else {
		return "Impossivel realizar essa transferencia no valor de", valor
	}
}

func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}
