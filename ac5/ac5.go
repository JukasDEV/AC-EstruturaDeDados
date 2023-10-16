package main

import (
    "fmt"
)

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    targetSum := 7

    num1, num2 := findPairWithSum(numbers, targetSum)

    if num1 == -1 && num2 == -1 {
        fmt.Println("Nenhum par encontrado")
    } else {
        fmt.Println("Par encontrado: ", num1, num2)
    }
}

func findPairWithSum(numbers []int, targetSum int) (int, int) {
    left := 0
    right := len(numbers) - 1

    for left < right {
        currentSum := numbers[left] + numbers[right]

        if currentSum == targetSum {
            return numbers[left], numbers[right]
        } else if currentSum < targetSum {
            left++
        } else {
            right--
        }
    }

    return -1, -1
}
