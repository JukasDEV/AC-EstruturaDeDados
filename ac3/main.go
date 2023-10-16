package main

import (
	"ac3/contato"
	"ac3/operacoes"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var contatos [5]contato.Contato
	var nome, email, opcao, novoEmail string
	var indice int

	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Lista de contatos!")
	for {
		fmt.Print("Digite (1) para adicionar, (2) para remover, (3) para alterar um email, (4) exibir todos os contatos ou qualquer outra coisa para sair: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Informe o seu nome: ")
			nome, _ = leitor.ReadString('\n')

			fmt.Print("Informe o seu email: ")
			fmt.Scanln(&email)

			fmt.Println(nome, email)
			contatos = operacoes.AdicionaContato(contato.Contato{Nome: nome, Email: email}, contatos)
		case "2":
			contatos = operacoes.ExcluiContato(contatos)

		case "3":
			fmt.Println("Informe o indice a ser alterado: ")
			fmt.Scanln(&indice)
			fmt.Println("Informe o novo email: ")
			fmt.Scanln(&novoEmail)
			contatos = operacoes.EditaEmail(indice, contatos, novoEmail)

		case "4":
			operacoes.ExibirContatos(contatos)

		default:
			return
		}

		fmt.Println(contatos)
	}

}
