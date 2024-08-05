package contas

import (
	"fmt"

	"github.com/marin/go-pacotes/clientes"
)

type ContaPoupanca struct {
	Titular                  clientes.Titular
	Agencia, Conta, Operacao int
	saldo                    float64
}

func (c *ContaPoupanca) Sacar(valor float64) {
	if c.saldo < valor || valor < 0 {
		fmt.Println("Não é possivel sacar este valor")
	} else {
		c.saldo -= valor + 1
	}
}

func (c *ContaPoupanca) Depositar(valor float64) {
	if valor < 0 {
		fmt.Println("Não é possivel depositar um valor negativo")
	} else {
		c.saldo += valor
		fmt.Println("Valor depositado com sucesso!", c.saldo)
	}
}

func (c *ContaPoupanca) Transferir(valor float64, destino *ContaPoupanca) (string, float64) {
	if c.saldo >= valor && valor > 0 {
		c.saldo -= valor
		destino.saldo += valor
		return "Transferencia realizada com sucesso no valor de", valor
	} else {
		return "Impossivel realizar essa transferencia no valor de", valor
	}
}

func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}
