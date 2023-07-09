package account

import (
	"math"

	"github.com/breno5g/oop-bank/client"
)

type Savings struct {
	Holder    client.Client
	Agency    int
	Number    int
	Operation int
	balance   float64
}

func (s *Savings) Withdraw(value float64) string {
	if value <= s.balance && !math.Signbit(value) {
		s.balance -= value
		return "Saque efetuado com sucesso"
	}

	return "Não foi possivel efetuar o saque"
}

func (s *Savings) Deposit(value float64) (string, float64) {
	if !math.Signbit(value) {
		s.balance += value
		return "Depósito efetuado com sucesso", s.balance
	}

	return "Não foi possivel efetuar o depósito", s.balance
}

func (s *Savings) GetBalance() float64 {
	return s.balance
}

func (s *Savings) PayBill(value float64) string {
	if !math.Signbit(value) && s.balance >= value {
		s.balance -= value
		return "Conta paga com sucesso"
	}

	return "Houve um erro ao pagar a conta"
}