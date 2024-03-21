package main

import (
	"log"

	"github.com/massimobiagioli/golang-serverless-terraform-demo/internal/app"
)

func init() {
	println("main.init")
}

func main() {
	fiberApp := app.Build()

	log.Fatal(fiberApp.Listen(":3000"))
}
