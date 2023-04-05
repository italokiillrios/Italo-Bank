package main

import (
	"fmt"
	contapoupanca "golang2/ContaPoupanca"
	c "golang2/contaCorrente"
)

func PagarBoleto(conta verficarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verficarConta interface {
	Sacar(valor float64) string
}

func main() {
	// clienteItalo := Cliente.Titular{"Italo", "123.123.123-10", "Desenvolvedor"}
	// contaDoItalo := c.ContaCorrente{clienteItalo, 123, 123456, 100}

	// fmt.Println(contaDoItalo)
	contaDoWill := contapoupanca.ContapouPanca{}
	contaDoWill.Depositar(100)
	PagarBoleto(&contaDoWill, 60)

	contaDaBeca := c.ContaCorrente{}
	contaDaBeca.Depositar(500)
	PagarBoleto(&contaDaBeca, 100)
	fmt.Println(contaDoWill.ObterSaldo())
	fmt.Println(contaDaBeca.ObterSaldo())

}
