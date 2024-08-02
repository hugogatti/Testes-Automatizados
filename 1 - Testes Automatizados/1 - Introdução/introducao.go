package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rua Manuel Justiniano Quintão")
	fmt.Println(tipoEndereco)

	tipoEndereco2 := enderecos.TipoDeEndereco("Jardim Maristela")
	fmt.Println(tipoEndereco2)
}