package main

import (
	"fmt"
	"math/rand"
	"time"
)

func imprimeMatriz(mat [][]int) {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			fmt.Printf("%3d ", mat[i][j])
		}
		fmt.Println()
	}
}

func copiaMatriz(orig [][]int) [][]int {
	ordem := len(orig)
	nova := make([][]int, ordem)
	for i := 0; i < ordem; i++ {
		nova[i] = make([]int, ordem)
		copy(nova[i], orig[i])
	}
	return nova
}

func gerarMatrizAleatoria(ordem int) [][]int {
	mat := make([][]int, ordem)
	maxValor := (ordem * ordem) / 2
	for i := 0; i < ordem; i++ {
		mat[i] = make([]int, ordem)
		for j := 0; j < ordem; j++ {
			mat[i][j] = rand.Intn(maxValor + 1)
		}
	}
	return mat
}

func menorMatriz(mat [][]int, linha int, coluna int) [][]int {
	ordem := len(mat)
	nova := make([][]int, ordem-1)
	for i := 0; i < ordem-1; i++ {
		nova[i] = make([]int, ordem-1)
	}
	novaI := 0
	for i := 0; i < ordem; i++ {
		if i != linha {
			novaJ := 0
			for j := 0; j < ordem; j++ {
				if j != coluna {
					nova[novaI][novaJ] = mat[i][j]
					novaJ++
				}
			}
			novaI++
		}
	}
	return nova
}

// Determinante baseline com expansão na primeira linha
func determinanteBaseline(mat [][]int) int {
	ordem := len(mat)
	if ordem == 1 {
		return mat[0][0]
	}
	if ordem == 2 {
		return mat[0][0]*mat[1][1] - mat[0][1]*mat[1][0]
	}
	det := 0
	sinal := 1
	for j := 0; j < ordem; j++ {
		det += sinal * mat[0][j] * determinanteBaseline(menorMatriz(mat, 0, j))
		sinal = -sinal
	}
	return det
}

// Função que escolhe a melhor linha ou coluna para expansão
func determinanteOtimizado(mat [][]int) int {
	ordem := len(mat)

	if ordem == 1 {
		return mat[0][0]
	}
	if ordem == 2 {
		return mat[0][0]*mat[1][1] - mat[0][1]*mat[1][0]
	}

	// Conta zeros por linha e coluna
	zerosLinha := make([]int, ordem)
	zerosColuna := make([]int, ordem)

	for i := 0; i < ordem; i++ {
		for j := 0; j < ordem; j++ {
			if mat[i][j] == 0 {
				zerosLinha[i]++
				zerosColuna[j]++
			}
		}
	}

	// Encontra linha e coluna com mais zeros
	maiorZerosLinha := 0
	linhaEscolhida := 0
	for i := 0; i < ordem; i++ {
		if zerosLinha[i] > maiorZerosLinha {
			maiorZerosLinha = zerosLinha[i]
			linhaEscolhida = i
		}
	}
	maiorZerosColuna := 0
	colunaEscolhida := 0
	for j := 0; j < ordem; j++ {
		if zerosColuna[j] > maiorZerosColuna {
			maiorZerosColuna = zerosColuna[j]
			colunaEscolhida = j
		}
	}

	det := 0
	sinal := 1

	if maiorZerosLinha >= maiorZerosColuna {
		// Expansão pela linhaEscolhida
		for j := 0; j < ordem; j++ {
			if mat[linhaEscolhida][j] != 0 {
				submat := menorMatriz(mat, linhaEscolhida, j)
				det += sinal * mat[linhaEscolhida][j] * determinanteOtimizado(submat)
			}
			sinal = -sinal
		}
	} else {
		// Expansão pela colunaEscolhida
		for i := 0; i < ordem; i++ {
			if mat[i][colunaEscolhida] != 0 {
				submat := menorMatriz(mat, i, colunaEscolhida)
				det += sinal * mat[i][colunaEscolhida] * determinanteOtimizado(submat)
			}
			sinal = -sinal
		}
	}

	return det
}

func executaTeste(ordem int) {
	var matriz [][]int
	var tempoBaseline int64
	var tempoOtimizado int64
	var mediaBaseline int64
	var mediaOtimizado int64
	var totalBaseline int64
	var totalOtimizado int64
	var detBase int
	var detOtim int
	var inicio time.Time
	var fim time.Time

	fmt.Printf("\n===== MATRIZ ORDEM %d =====\n\n", ordem)

	for i := 0; i < 3; i++ {
		matriz = gerarMatrizAleatoria(ordem)

		fmt.Printf("Execução %d\n", i+1)

		// Baseline
		fmt.Println("Matriz baseline:")
		imprimeMatriz(matriz)
		inicio = time.Now()
		detBase = determinanteBaseline(matriz)
		fim = time.Now()
		tempoBaseline = fim.Sub(inicio).Nanoseconds()
		fmt.Printf("Tempo(ns): %d\n", tempoBaseline)
		fmt.Printf("Determinante: %d\n\n", detBase)
		mediaBaseline += tempoBaseline
		totalBaseline += tempoBaseline

		// Otimizado
		matrizOtimizada := copiaMatriz(matriz)
		fmt.Println("Matriz otimizada:")
		imprimeMatriz(matrizOtimizada)
		inicio = time.Now()
		detOtim = determinanteOtimizado(matrizOtimizada)
		fim = time.Now()
		tempoOtimizado = fim.Sub(inicio).Nanoseconds()
		fmt.Printf("Tempo(ns): %d\n", tempoOtimizado)
		fmt.Printf("Determinante: %d\n\n", detOtim)
		mediaOtimizado += tempoOtimizado
		totalOtimizado += tempoOtimizado

		fmt.Println("====================================\n")
	}

	fmt.Printf("MÉDIA TEMPO BASELINE (ordem %d): %d ns\n", ordem, mediaBaseline/3)
	fmt.Printf("MÉDIA TEMPO OTIMIZADO (ordem %d): %d ns\n", ordem, mediaOtimizado/3)
	fmt.Println("====================================")
}

func main() {
	ordens := []int{3, 5, 7, 9, 11}
	tempoInicio := time.Now()
	for _, ordem := range ordens {
		executaTeste(ordem)
	}
	tempoFim := time.Now()
	tempoTotal := tempoFim.Sub(tempoInicio).Nanoseconds()
	fmt.Printf("\nTEMPO TOTAL DE EXECUÇÃO (baseline + otimizado): %d ns\n", tempoTotal)
}
