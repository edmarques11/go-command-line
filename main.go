package main

import (
	"fmt"
	"go-command-line/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Start point")

	application := app.Generate()
	if err := application.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
