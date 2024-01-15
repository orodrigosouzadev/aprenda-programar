package main

import "fmt"

func main() {
	// Janeiro
	salario1 := 3000.0
	aluguel1 := 1000.0
	contaAgua1 := 200.0
	contaLuz1 := 100.0
	internet1 := 100.0
	transporte1 := 300.0
	lazer1 := 300.0
	alimentacao1 := 500.0
	saldo1 := salario1 - aluguel1 - contaAgua1 - contaLuz1 - internet1 - transporte1 - lazer1 - alimentacao1
	fmt.Println("Janeiro")
	fmt.Println(saldo1)
	estaNegativo1 := saldo1 < 0
	if estaNegativo1 {
		fmt.Println("Juros de janeiro")
		juros1 := saldo1 * 0.1
		fmt.Println(juros1)
		saldo1 = saldo1 + juros1
	} else {
		fmt.Println("Rendimentos de janeiro")
		rendimentos1 := saldo1 * 0.005
		fmt.Println(rendimentos1)
		saldo1 = saldo1 + rendimentos1
	}
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
	saldo2 := salario2 - aluguel2 - contaAgua2 - contaLuz2 - internet2 - transporte2 - lazer2 - alimentacao2
	fmt.Println("Fevereiro")
	fmt.Println(saldo2)
	estaNegativo2 := saldo2 < 0
	if estaNegativo2 {
		fmt.Println("Juros de fevereiro")
		juros2 := saldo2 * 0.1
		fmt.Println(juros2)
		saldo2 = saldo2 + juros2
	} else {
		fmt.Println("Rendimentos de fevereiro")
		rendimentos2 := saldo2 * 0.005
		fmt.Println(rendimentos2)
		saldo2 = saldo2 + rendimentos2
	}
	fmt.Println(saldo2)

	acumuladoAno := saldo1 + saldo2
	fmt.Println("Ano")
	fmt.Println(acumuladoAno)
}
