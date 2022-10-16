package main

import (
	"fmt"
	"ippublico/ippublico"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto Inicial")

	app := ippublico.Run()

	// criando variável erro que recebe o app.run
	// em seguida, se error for diferente de nulo, mostra qual foi o erro
	if error := app.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
