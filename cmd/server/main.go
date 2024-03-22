package main

import (
	"log"

	"github.com/massimobiagioli/golang-serverless-terraform-demo/internal/app"
)

func main() {
	fiberApp := app.Build()

	log.Fatal(fiberApp.Listen(":3000"))
}
