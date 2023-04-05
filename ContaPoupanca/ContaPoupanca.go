package contapoupanca

import c "golang2/Cliente"

type ContapouPanca struct {
	Titular     c.Titular
	Agencia     int
	NumeroConta int
	Operacao    int
	saldo       float64
}

func (c *ContapouPanca) Sacar(valorDoSaque float64) string {
	podesacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podesacar {
		c.saldo -= valorDoSaque
		return "saque realizado com suecesso"

	} else {
		return "saldo insuficiente"
	}

}
func (c *ContapouPanca) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito é menor que 0", c.saldo
	}

}
func (c *ContapouPanca) ObterSaldo() float64 {
	return c.saldo
}
