package main

import (
	"ippublico/ippublico"
	"log"
	"os"
)

func main() {

	app := ippublico.Run()

	// criando variável erro que recebe o app.run
	// em seguida, se error for diferente de nulo, mostra qual foi o erro
	if error := app.Run(os.Args); error != nil {
		log.Fatalln(error)
	}
}
