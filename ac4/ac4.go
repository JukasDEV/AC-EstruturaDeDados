package main

import "fmt"

func main() {
	torreDeHanoi(1, "A", "C", "B")
}

func torreDeHanoi(numero int, inicio, destino, doMeio string) {
	if n == 1 {
		fmt.Printf(inicio + " -> " + destino + "\n")
		return
	}

	torreDeHanoi(n-1, inicio, doMeio, destino)

	fmt.Printf(origem + " -> " + destino + "\n")

	torreDeHanoi(n-1, doMeio, destino, inicio)
}
