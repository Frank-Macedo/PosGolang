package main

import (
	"fmt"

	tempconv "github.com/Frank-Macedo/EstudoAParte"
)

func main() {
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Println(tempconv.KToC(273.15))
	fmt.Println(tempconv.FToK(32))
	fmt.Println(tempconv.KToF(273.15))
	fmt.Println(tempconv.FToC(32))

}
