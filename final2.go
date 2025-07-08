package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Matriz struct {
	valores [][]int
	ordem   int
}

func gerarMatrizAleatoria(ordem int) [][]int {
	var matriz [][]int
	var contI int
	var contJ int
	var maxValor int

	maxValor = ordem * ordem
	matriz = make([][]int, ordem)

	for contI = 0; contI < ordem; contI++ {
		matriz[contI] = make([]int, ordem)
		for contJ = 0; contJ < ordem; contJ++ {
			matriz[contI][contJ] = rand.Intn(maxValor + 1)
		}
	}
	return matriz
}

func imprimirMatriz(matriz [][]int) {
	var contI int
	var contJ int
	for contI = 0; contI < len(matriz); contI++ {
		for contJ = 0; contJ < len(matriz[contI]); contJ++ {
			fmt.Printf("%4d ", matriz[contI][contJ])
		}
		fmt.Println()
	}
}

func copiarMatriz(origem [][]int) [][]int {
	var nova [][]int
	var contI int
	var contJ int
	var ordem int

	ordem = len(origem)
	nova = make([][]int, ordem)
	for contI = 0; contI < ordem; contI++ {
		nova[contI] = make([]int, ordem)
		for contJ = 0; contJ < ordem; contJ++ {
			nova[contI][contJ] = origem[contI][contJ]
		}
	}
	return nova
}

func calcularDeterminante(matriz [][]int) int {
	var ordem int
	ordem = len(matriz)
	if ordem == 1 {
		return matriz[0][0]
	}
	if ordem == 2 {
		return matriz[0][0]*matriz[1][1] - matriz[0][1]*matriz[1][0]
	}
	var det int
	var contJ int
	var sinal int
	var submatriz [][]int

	sinal = 1
	for contJ = 0; contJ < ordem; contJ++ {
		submatriz = gerarSubmatriz(matriz, 0, contJ)
		det = det + sinal*matriz[0][contJ]*calcularDeterminante(submatriz)
		sinal = sinal * -1
	}
	return det
}

func gerarSubmatriz(matriz [][]int, linhaExcluir int, colunaExcluir int) [][]int {
	var submatriz [][]int
	var ordem int
	var contI int
	var contJ int
	var novaLinha []int

	ordem = len(matriz)

	for contI = 0; contI < ordem; contI++ {
		if contI <= linhaExcluir-1 || contI >= linhaExcluir+1 {
			novaLinha = []int{}
			for contJ = 0; contJ < ordem; contJ++ {
				if contJ <= colunaExcluir-1 || contJ >= colunaExcluir+1 {
					novaLinha = append(novaLinha, matriz[contI][contJ])
				}
			}
			submatriz = append(submatriz, novaLinha)
		}
	}
	return submatriz
}

func linhaOuColunaComMaisZeros(matriz [][]int) (int, bool) {
	var ordem int
	ordem = len(matriz)
	var contI int
	var contJ int
	var zerosNaLinha int
	var zerosNaColuna int
	var maxZeros int
	var indice int
	var isLinha bool

	for contI = 0; contI < ordem; contI++ {
		zerosNaLinha = 0
		zerosNaColuna = 0
		for contJ = 0; contJ < ordem; contJ++ {
			if matriz[contI][contJ] <= 0 && matriz[contI][contJ] >= 0 {
				zerosNaLinha = zerosNaLinha + 1
			}
			if matriz[contJ][contI] <= 0 && matriz[contJ][contI] >= 0 {
				zerosNaColuna = zerosNaColuna + 1
			}
		}
		if zerosNaLinha > maxZeros {
			maxZeros = zerosNaLinha
			indice = contI
			isLinha = true
		}
		if zerosNaColuna > maxZeros {
			maxZeros = zerosNaColuna
			indice = contI
			isLinha = false
		}
	}
	return indice, isLinha
}

func calcularDeterminanteOtimizado(matriz [][]int) int {
	var ordem int
	ordem = len(matriz)
	if ordem == 1 {
		return matriz[0][0]
	}
	if ordem == 2 {
		return matriz[0][0]*matriz[1][1] - matriz[0][1]*matriz[1][0]
	}
	var det int
	var contI int
	var indice int
	var isLinha bool
	var submatriz [][]int
	var sinal int

	indice, isLinha = linhaOuColunaComMaisZeros(matriz)
	sinal = 1

	for contI = 0; contI < ordem; contI++ {
		if isLinha {
			if matriz[indice][contI] <= 0 && matriz[indice][contI] >= 0 {
				continue
			}
			submatriz = gerarSubmatriz(matriz, indice, contI)
			det = det + sinal*matriz[indice][contI]*calcularDeterminanteOtimizado(submatriz)
		} else {
			if matriz[contI][indice] <= 0 && matriz[contI][indice] >= 0 {
				continue
			}
			submatriz = gerarSubmatriz(matriz, contI, indice)
			det = det + sinal*matriz[contI][indice]*calcularDeterminanteOtimizado(submatriz)
		}
		sinal = sinal * -1
	}
	return det
}

func medirTempoExecucao(matriz [][]int, funcao func([][]int) int) (int64, int) {
	var inicio time.Time
	var fim time.Time
	var duracao int64
	var determinante int

	inicio = time.Now()
	determinante = funcao(matriz)
	fim = time.Now()
	duracao = fim.Sub(inicio).Nanoseconds()
	return duracao, determinante
}

func main() {
	var ordens [5]int
	ordens[0] = 3
	ordens[1] = 5
	ordens[2] = 7
	ordens[3] = 9
	ordens[4] = 11

	var contOrdem int
	var contExecucao int
	var matriz [][]int
	var copia [][]int
	var tempoBaseline int64
	var tempoOtimizado int64
	var determinanteBaseline int
	var determinanteOtimizado int

	for contOrdem = 0; contOrdem < 5; contOrdem++ {
		fmt.Printf("================= ORDEM %d =================\n\n", ordens[contOrdem])
		for contExecucao = 1; contExecucao <= 3; contExecucao++ {
			fmt.Printf("---- EXPERIMENTO %d ----\n\n", contExecucao)

			matriz = gerarMatrizAleatoria(ordens[contOrdem])
			copia = copiarMatriz(matriz)

			fmt.Println("Matriz baseline:")
			imprimirMatriz(matriz)
			tempoBaseline, determinanteBaseline = medirTempoExecucao(matriz, calcularDeterminante)
			fmt.Printf("Tempo (ns): %d\n", tempoBaseline)
			fmt.Printf("Determinante: %d\n", determinanteBaseline)
			fmt.Println("====================================\n")

			fmt.Println("Matriz otimizada:")
			imprimirMatriz(copia)
			tempoOtimizado, determinanteOtimizado = medirTempoExecucao(copia, calcularDeterminanteOtimizado)
			fmt.Printf("Tempo (ns): %d\n", tempoOtimizado)
			fmt.Printf("Determinante: %d\n", determinanteOtimizado)
			fmt.Println("\n\n")
		}
	}
}
