package contaCorrente

import c "golang2/Cliente"

type ContaCorrente struct {
	Titular     c.Titular
	Agencia     int
	NumeroConta int
	saldo       float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podesacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podesacar {
		c.saldo -= valorDoSaque
		return "saque realizado com suecesso"

	} else {
		return "saldo insuficiente"
	}

}
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito é menor que 0", c.saldo
	}

}
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia > 0 && valorDaTransferencia < c.saldo {
		c.saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}
func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
