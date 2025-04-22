package main

import "fmt"

func main() {

	var minhaVar interface{} = "Franklin Macedo"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Println("O valor de res é %v e o resultado de ok é %v", res, ok)
	res2 := minhaVar.(int)
	fmt.Println("O valor de res é %v ", res2)

}
