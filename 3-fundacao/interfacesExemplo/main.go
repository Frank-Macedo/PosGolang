package main

//Em Go, uma interface define um conjunto de métodos. Qualquer tipo que implementar todos os métodos definidos na interface,
//  automaticamente satisfaz essa interface, sem a necessidade de uma declaração explícita. Essa característica é conhecida como interfaces implícitas.
//  As interfaces em Go servem como contratos que um tipo deve seguir, especificando o comportamento que ele deve ter.
//  Elas são usadas para criar abstrações e permitir polimorfismo, onde diferentes tipos podem ser tratados de forma
//  uniforme se implementarem a mesma interface

import "fmt"

func main() {

	Pessoa := Cliente{Nome: "Franklin", Tipo: "Normal", Ativo: true}

	Desativacao(Pessoa)

}

type Cliente struct {
	Nome  string
	Tipo  string
	Ativo bool
}

type Pessoa interface {
	Desativar()
}

func Desativacao(Pessoa Pessoa) {
	Pessoa.Desativar()
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Println("Cliente desativado")
}
