package main

import (
	"fmt"
	"gerador/src"
)

func main() {
	cnpj := src.GeneratorCNPJ()
	fmt.Println("Sem formatação", cnpj)
	fmt.Println("Com formatação", src.FormatCNPJ(cnpj))
	cpf := src.GeneratorCPF()
	fmt.Println("\nSem formatação", cpf)
	fmt.Println("Com formatação", src.FormatCPF(cpf))
	fmt.Println("CPF válido?", src.IsCPFValid(cpf))
}
