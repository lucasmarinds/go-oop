package contas

import (
	"fmt"

	"github.com/marin/go-pacotes/clientes"
)

/*
Quando deixamos o nome dos atributos do objeto(struct) em minusculo colocamos como se fosse "private" na visibilidade e o mesmo serve para Métodos
*/
type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

/*
Esses c *ContaCorrente, estamos informando que
a conta que chamar esse método e só variaveis de referencia para contacorrente
podem chamar, a que chamar vai executar ela usando os valores dela,
seria o this.Saldo em java, estamos dizendo que é o OBJETO que está chamando
para alterar o atributo apenas dele.
*/
func (c *ContaCorrente) Sacar(valor float64) {
	if c.Saldo < valor || valor < 0 {
		fmt.Println("Não é possivel sacar este valor")
	} else {
		c.Saldo -= valor
	}
}

func (c *ContaCorrente) Depositar(valor float64) {
	if valor < 0 {
		fmt.Println("Não é possivel depositar um valor negativo")
	} else {
		c.Saldo += valor
		fmt.Println("Valor depositado com sucesso!", c.Saldo)
	}
}

func (c *ContaCorrente) Transferir(valor float64, destino *ContaCorrente) (string, float64) {
	if c.Saldo >= valor && valor > 0 {
		c.Saldo -= valor
		destino.Saldo += valor
		return "Transferencia realizada com sucesso no valor de", valor
	} else {
		return "Impossivel realizar essa transferencia no valor de", valor
	}
}
