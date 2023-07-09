package main

import (
	"fmt"

	"github.com/breno5g/oop-bank/account"
	"github.com/breno5g/oop-bank/client"
)

type Payment interface {
	PayBill(value float64) string
}

func PayBill(account Payment, value float64)  string {
	res := account.PayBill(value)
	return res
}

func main() {
	// Short atribution
	// account := Account{
	// 	"Breno Santos", 123456, 654321, 42.42,
	// }

	test := account.Checking{
		Holder:  client.Client{Name: "Breno Santos", Cpf: "123.456.789-10", Job: "Developer"},
		Agency:  123456,
		Number:  654321,
	}

	// test2 := account.Checking{
	// 	Holder: client.Client{Name: "Teste", Cpf: "987.654.321-10", Job: "Tester"},
	// 	Agency: 456789,
	// 	Number: 987654,
	// }

	// fmt.Println(test2.Transfer(400, &test))

	// fmt.Println(test.GetBalance())
	// fmt.Println(test2.GetBalance())

	test.Deposit(200)

	fmt.Println(PayBill(&test, 300))
}
