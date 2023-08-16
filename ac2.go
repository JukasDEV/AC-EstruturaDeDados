package main

import (
	"bufio"
	"fmt"
	"os"
)

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) AlterarEmail(novoEmail string) {
	c.Email = novoEmail
}

func adicionaContato(contato Contato, arrayContatos []Contato) {
	for i := 0; i < len(arrayContatos); i++ {
		if arrayContatos[i].Nome == "" && arrayContatos[i].Email == "" {
			arrayContatos[i] = contato
			return
		}
	}
	fmt.Println("array está cheio")
}

func excluiContato(arrayContatos []Contato) {
	for i := len(arrayContatos) - 1; i >= 0; i-- {
		if arrayContatos[i].Nome != "" && arrayContatos[i].Email != "" {
			arrayContatos[i] = Contato{}
			return
		}
	}
	fmt.Println("Sem Contatos para excluir.")
}

func main() {
	contatos := make([]Contato, 5)

	for {
		fmt.Println("\nOpções:")
		fmt.Println("1. Adicionar Contato")
		fmt.Println("2. Excluir Último Contato")
		fmt.Println("3. Sair")

		var escolha int
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&escolha)

		switch escolha {
		case 1:
			fmt.Println("Digite o nome do contato:")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			nome := scanner.Text()

			fmt.Println("Digite o email do contato:")
			scanner.Scan()
			email := scanner.Text()

			adicionaContato(Contato{Nome: nome, Email: email}, contatos)
			fmt.Println("Contato adicionado com sucesso!")

		case 2:
			excluiContato(contatos)
			fmt.Println("Último contato excluído com sucesso!")

		case 3:
			fmt.Println("Encerrando o programa.")
			return

		default:
			fmt.Println("Opção inválida. Escolha uma opção válida.")
		}

		fmt.Println("\nContatos:")
		for _, contato := range contatos {
			fmt.Printf("Nome: %s, E-mail: %s\n", contato.Nome, contato.Email)
		}
	}
}
