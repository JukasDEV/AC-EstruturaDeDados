package operacoes

import (
	"ac3/contato"
	"fmt"
)

func AdicionaContato(novoContato contato.Contato, a [5]contato.Contato) [5]contato.Contato {
	for ind, c := range a {
		if (c == contato.Contato{}) {
			a[ind] = novoContato
			break
		}
	}

	return a
}

func ExcluiContato(a [5]contato.Contato) [5]contato.Contato {
	if (a[0] == contato.Contato{}) {
		fmt.Println("Lista de contatos vazia!")
		return a
	}

	for ind, c := range a {
		if (c == contato.Contato{}) {
			a[ind-1] = contato.Contato{}
		}
	}

	return a
}

// Crie uma nova função editaEmail, que recebe o índice do elemento no array e altera o e-mail do elemento indicado.
func EditaEmail(indiceContato int, array [5]contato.Contato, novoEmail string) [5]contato.Contato {
	for ind, c := range array {
		if ind == indiceContato {

			array[ind].Email = novoEmail
			fmt.Println(c)
		}
	}

	return array
}

func ExibirContatos(array [5]contato.Contato) {
	for ind, c := range array {
		fmt.Println(ind, c, "\n")
	}
}
