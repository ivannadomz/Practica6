package main

import (
	"fmt"

	"github.com/ivannadomz/Practica6/cmd/config"
)

func main() {
	fmt.Println("Variables de entorno: ", config.USERNAME, config.PASSWORD)
}
