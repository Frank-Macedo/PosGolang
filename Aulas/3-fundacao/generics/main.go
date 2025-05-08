package main

import "moduloInicial/Aulas/3-fundacao/pacotes/matematica"

type MyNumber int

type Number interface {
	~int | float64
}

func main() {

	m := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Wesley": 100.10, "João": 200.20, "Maria": 300.30}
	m3 := map[string]MyNumber{"Wesley": 200, "João": 300, "Maria": 400}

	println(matematica.SomaMatematica(m["Wesley"], m["João"]))
	println(Soma(m2))
	println(Soma(m3))

	println(Compara(10.0, 20))

}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}
