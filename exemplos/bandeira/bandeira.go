package main

import (
	"fmt"
	"os"
	"strings"
)

const indicadorA = 0x1F1E6 // REGIONAL INDICATOR SYMBOL LETTER A

func main() {
	if len(os.Args) != 2 || len(os.Args[1]) != 2 {
		fmt.Println("Informe o código do país. Ex: BR ou br")
		os.Exit(0)
	}
	indicadores := []string{}
	for _, c := range strings.ToUpper(os.Args[1]) {
		código := rune(c - 'A' + indicadorA)
		indicadores = append(indicadores, string(código))
	}
	// terminal do Ubuntu 18.04 ainda não reconhece REGIONAL INDICATORs..
	fmt.Println(strings.Join(indicadores, ""))
}
