package main

import (
	"fmt"
)

const MaxSize = 5

func main() {
	var arr [MaxSize]int
	var size = 0

	insertSorted(&arr, &size, 4)
	fmt.Println(arr)

	insertSorted(&arr, &size, 12)
	fmt.Println(arr)

	insertSorted(&arr, &size, 2)
	fmt.Println(arr)

	insertSorted(&arr, &size, 6)
	fmt.Println(arr)

	insertSorted(&arr, &size, 17)
	fmt.Println(arr)

	insertSorted(&arr, &size, 1)
	fmt.Println(arr)
}

func insertSorted(array *[MaxSize]int, size *int, newValue int) {
	if *size == MaxSize {
		fmt.Println("Overflow")
		return
	}

	i := *size - 1
	for i >= 0 && newValue < array[i] {
		array[i+1] = array[i]
		i--
	}

	array[i+1] = newValue
	*size++
	fmt.Printf("Tentando inserir %d\n", newValue)
	if i == -1 || array[i] != newValue {
		fmt.Printf("%d não encontrado, inserindo na lista\n", newValue)
	}
}
