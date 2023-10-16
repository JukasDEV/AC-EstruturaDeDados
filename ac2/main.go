package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	// e_primo()
	//e_bissexto()
	e_fibo(1)

}

func e_primo() {

	var i, numero int

	var primo bool = true

	fmt.Print("Informe um número inteiro : ")
	fmt.Scanln(&numero)

	for i = 2; i <= int(numero/2); i++ {
		if numero%i == 0 {
			primo = false
			break
		}
	}

	if primo {
		fmt.Print("O número informado é primo")
	} else {
		fmt.Print("O número informado não é primo")
	}
}

func e_bissexto() {
	var Ano int

	fmt.Print("Informe um Ano : ")
	fmt.Scanln(&Ano)

	if Ano%4 == 0 {
		if Ano%100 == 0 && Ano%400 != 0 {
			fmt.Print("Não é Bissexto")
		}
		fmt.Print("é bissexto")
	} else {
		fmt.Print("Não é Bissexto")

	}
}

func e_fibo(num int) {
	listaAuxiliar := []int{num}
	limite := 9
	var intAuxiliar int
	for i := 0; i <= limite; i++ {
		if len(listaAuxiliar) == 1 {
			intAuxiliar = listaAuxiliar[len(listaAuxiliar)-1]
			listaAuxiliar = append(listaAuxiliar, intAuxiliar)
		} else {
			intAuxiliar = listaAuxiliar[len(listaAuxiliar)-1] + listaAuxiliar[len(listaAuxiliar)-2]
			listaAuxiliar = append(listaAuxiliar, intAuxiliar)
		}
	}

	fmt.Println(listaAuxiliar)
}
