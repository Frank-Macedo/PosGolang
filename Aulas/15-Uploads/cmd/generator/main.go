package main

import (
	"fmt"
	"os"
)

func main() {
	i := 0

	for {

		f, err := os.Create(fmt.Sprintf("/home/franklin/repositorios/PosGolang/Aulas/15-Uploads/tmp/file%d.txt", i))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.WriteString("hello, world\n")
		i++
	}

}
