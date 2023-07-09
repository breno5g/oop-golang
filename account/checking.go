package account

import (
	"math"

	"github.com/breno5g/oop-bank/client"
)

type Checking struct {
	Holder  client.Client
	Agency  int
	Number  int
	balance float64
}

func (c *Checking) Withdraw(value float64) string {
	if value <= c.balance && !math.Signbit(value) {
		c.balance -= value
		return "Saque realizado"
	}

	return "Saldo insuficiente"
}

func (c *Checking) Deposit(value float64) string {
	if !math.Signbit(value) {
		c.balance += value
		return "Deposito efetuado com sucesso"
	}
	return "O valor de deposito é invalido"
}

func (c *Checking) Transfer(value float64, destinyChecking *Checking) bool {
	if value < c.balance && !math.Signbit(value) {
		c.balance -= value
		destinyChecking.Deposit(value)
		return true
	}

	return false
}

func (c *Checking) GetBalance() float64 {
	return c.balance
}

func (c *Checking) PayBill(value float64) string {
	if !math.Signbit(value) && c.balance >= value {
		c.balance -= value
		return "Conta paga com sucesso"
	}

	return "Houve um erro ao pagar a conta"
}