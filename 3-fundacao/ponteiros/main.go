package main

import "fmt"

type Conta struct {
	nome  string
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func main() {

	c1 := NewConta()

	c1.Deposita(50)

	println("\n")

	println(c1.saldo)

}

func (c *Conta) Deposita(a int) {
	c.saldo += a
	fmt.Println(c.saldo)
}
