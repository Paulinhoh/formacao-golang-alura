package main

import (
	"fmt"
	"go-orientacao-objetos/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valorSaque float64) string
}

func main() {
	contaDoGuilherme := contas.ContaCorrente{}
	contaDaAna := contas.ContaPoupanca{}

	contaDaAna.Depositar(200)
	PagarBoleto(&contaDaAna, 120)

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaAna)

}
