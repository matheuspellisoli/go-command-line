package main

import (
	"command-line/app"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	application := app.Generate()

	application.Run(os.Args)
}
