package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var cpf string

	for i := 0; i < 9; i++ {
		cpf += fmt.Sprintf("%d", rand.Intn(9))
	}

	sum := calculator(cpf, 10)
	cpf += fmt.Sprintf("%d", getDigit(sum))

	sum = calculator(cpf, 11)
	cpf += fmt.Sprintf("%d", getDigit(sum))
	fmt.Println("Sem formatação", cpf)
	fmt.Println("Com formatação", formatCPF(cpf))
	fmt.Println("CPF válido?", isValid(cpf))
}

func calculator(cpf string, first int) int {
	var sum int
	for _, value := range cpf {
		var rs = int(value-'0') * first
		sum += rs
		first--
	}
	return sum
}

func formatCPF(cpf string) string {
	return cpf[:3] + "." + cpf[3:6] + "." + cpf[6:9] + "-" + cpf[9:]
}

func removeFormat(cpf string) string {
	if len(cpf) == 11 {
		return cpf
	}

	return cpf[:3] + cpf[4:7] + cpf[8:11] + cpf[12:]
}

func getDigit(sum int) int {
	rest := sum % 11
	if rest < 2 {
		return 0
	}
	return 11 - rest
}

func isValid(cpf string) bool {
	cpf = removeFormat(cpf)
	if len(cpf) != 11 {
		return false
	}

	var sum int
	for i := 0; i < 9; i++ {
		sum += int(cpf[i]-'0') * (10 - i)
	}
	firstDigit := getDigit(sum)

	if firstDigit != int(cpf[9]-'0') {
		return false
	}

	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(cpf[i]-'0') * (11 - i)
	}
	sum = sum % 11
	secondDigit := getDigit(sum)

	if secondDigit != int(cpf[10]-'0') {
		return false
	}

	return true
}
