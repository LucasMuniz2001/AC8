package main

import (
	"fmt"
)

type No struct {
	Chave int
	Esquerda *No
	Direita  *No
}

func inserir(raiz *No, chave int) *No {
	if raiz == nil {
		return &No{Chave: chave, Esquerda: nil, Direita: nil}
	}

	if chave < raiz.Chave {
		raiz.Esquerda = inserir(raiz.Esquerda, chave)
	} else if chave > raiz.Chave {
		raiz.Direita = inserir(raiz.Direita, chave)
	}

	return raiz
}

func buscar(raiz *No, chave int) *No {
	if raiz == nil || raiz.Chave == chave {
		return raiz
	}

	if chave < raiz.Chave {
		return buscar(raiz.Esquerda, chave)
	}

	return buscar(raiz.Direita, chave)
}

func percorrerEmOrdem(raiz *No) {
	if raiz != nil {
		percorrerEmOrdem(raiz.Esquerda)
		fmt.Printf("%d ", raiz.Chave)
		percorrerEmOrdem(raiz.Direita)
	}
}

func main() {
	
	var raiz *No

  raiz = inserir(raiz, 50)
	inserir(raiz, 30)
	inserir(raiz, 20)
	inserir(raiz, 40)
	inserir(raiz, 70)
	inserir(raiz, 60)
	inserir(raiz, 80)

	
	fmt.Println("Árvore em ordem:")
	percorrerEmOrdem(raiz)
	fmt.Println()

	
	chaveBusca := 40
	resultado := buscar(raiz, chaveBusca)

	if resultado != nil {
		fmt.Printf("O valor %d foi encontrado na árvore.\n", chaveBusca)
	} else {
		fmt.Printf("O valor %d não foi encontrado na árvore.\n", chaveBusca)
	}
}
