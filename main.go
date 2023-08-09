package main

import {
	"fmt"
}

// Elabore uma função e_primo() que retorna se um número é primo ou não. Caso o número não seja primo, liste todos os números pelos quais ele é divisível.


func main(){

// fmt.println(e_primo)
e_primo()

}

func e_primo() {

var i, numero int

var primo bool = true

fmt.Print("Informe um número inteiro : ")
  fmt.Scanln(&numero)

  for i = 2; i <= int(numero / 2); i++ {
	if numero % i == 0 {
	  primo = false
	  break
	}
  }    
   
  if(primo){
	fmt.Print("O número informado é primo")
  }else{
	fmt.Print("O número informado não é primo")
  }
}




}