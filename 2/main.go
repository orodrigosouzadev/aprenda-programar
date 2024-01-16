package main

import (
	"fmt"
	"math"
)

func arredondar(valor float64) float64 {
	return math.Round(valor*100) / 100
}

func calcularJuros(valor float64) float64 {
	return arredondar(valor * 0.01)
}

func calcularRendimentos(valor float64) float64 {
	return arredondar(valor * 0.005)
}

func calcularSaldo(mes string, saldoInicial, salario, aluguel, contaAgua, contaLuz, internet, transporte, lazer, alimentacao float64) float64 {
	fmt.Println(mes)
	saldo := saldoInicial + salario - aluguel - contaAgua - contaLuz - internet - transporte - lazer - alimentacao
	estaNegativo := saldo < 0
	if estaNegativo {
		juros := calcularJuros(saldo)
		saldo = arredondar(saldo + juros)
	} else {
		rendimentos := calcularRendimentos(saldo)
		saldo = arredondar(saldo + rendimentos)
	}
	return saldo
}

func main() {
	saldoInicial := 0.0
	// Janeiro
	salario1 := 3000.0
	aluguel1 := 1000.0
	contaAgua1 := 200.0
	contaLuz1 := 100.0
	internet1 := 100.0
	transporte1 := 300.0
	lazer1 := 300.0
	alimentacao1 := 500.0
	saldo1 := calcularSaldo("Janeiro", saldoInicial, salario1, aluguel1, contaAgua1, contaLuz1, internet1, transporte1, lazer1, alimentacao1)
	fmt.Println(saldo1)

	// Fevereiro
	salario2 := 3000.0
	aluguel2 := 1200.0
	contaAgua2 := 250.0
	contaLuz2 := 100.0
	internet2 := 100.0
	transporte2 := 500.0
	lazer2 := 0.0
	alimentacao2 := 1000.0
	saldo2 := calcularSaldo("Fevereiro", saldo1, salario2, aluguel2, contaAgua2, contaLuz2, internet2, transporte2, lazer2, alimentacao2)
	fmt.Println(saldo2)

	// Março
	salario3 := 4000.0
	aluguel3 := 1200.0
	contaAgua3 := 200.0
	contaLuz3 := 100.0
	internet3 := 200.0
	transporte3 := 500.0
	lazer3 := 800.0
	alimentacao3 := 1000.0
	saldo3 := calcularSaldo("Março", saldo2, salario3, aluguel3, contaAgua3, contaLuz3, internet3, transporte3, lazer3, alimentacao3)
	fmt.Println(saldo3)

	acumuladoAno := saldo3
	fmt.Println("Ano")
	fmt.Println(acumuladoAno)
}
