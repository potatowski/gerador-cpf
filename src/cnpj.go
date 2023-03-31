package src

import (
	"fmt"
	"math/rand"
	"regexp"
)

func GeneratorCNPJ() string {
	var cnpj string

	for i := 0; i < 12; i++ {
		cnpj += fmt.Sprintf("%d", rand.Intn(9))
	}

	sum := calculator(cnpj, 5)
	cnpj += fmt.Sprintf("%d", getDigit(sum))

	sum = calculator(cnpj, 6)
	cnpj += fmt.Sprintf("%d", getDigit(sum))

	return cnpj
}

func FormatCNPJ(cnpj string) string {
	return cnpj[:2] + "." + cnpj[2:5] + "." + cnpj[5:8] + "/" + cnpj[8:12] + "-" + cnpj[12:]
}

func removeFormat(doc string) string {
	return regexp.MustCompile("[^0-9]+").ReplaceAllString(doc, "")
}
