package main

import (
	"fmt"
	"math/rand"
	"time"
)

var matriz [][]int
var matrizOtimizada [][]int
var ContI int
var ContJ int
var linha int
var coluna int

var zerosLinha []int
var zerosColuna []int

var detBase int
var detOtim int

var tempoBaseline int64
var tempoOtimizado int64

var mediaBaseline int64
var mediaOtimizado int64

var totalBaseline int64
var totalOtimizado int64

var tempoInicio time.Time
var tempoFim time.Time
var tempoTotal int64

func imprimeMatriz(mat [][]int) {
	for ContI = 0; ContI < len(mat); ContI++ {
		for ContJ = 0; ContJ < len(mat[0]); ContJ++ {
			fmt.Printf("%3d ", mat[ContI][ContJ])
		}
		fmt.Println()
	}
}

func copiaMatriz(orig [][]int) [][]int {
	var ordem int
	var nova [][]int
	var ContI int

	ordem = len(orig)
	nova = make([][]int, ordem)
	for ContI = 0; ContI < ordem; ContI++ {
		nova[ContI] = make([]int, ordem)
		copy(nova[ContI], orig[ContI])
	}
	return nova
}

func gerarMatrizAleatoria(ordem int) [][]int {
	var mat [][]int
	var ContI int
	var ContJ int
	var maxValor int

	mat = make([][]int, ordem)
	maxValor = (ordem * ordem) / 2
	for ContI = 0; ContI < ordem; ContI++ {
		mat[ContI] = make([]int, ordem)
		for ContJ = 0; ContJ < ordem; ContJ++ {
			mat[ContI][ContJ] = rand.Intn(maxValor + 1)
		}
	}
	return mat
}

func menorMatriz(mat [][]int, linha int, coluna int) [][]int {
	var ordem int
	var nova [][]int
	var i int
	var j int
	var novaI int
	var novaJ int

	ordem = len(mat)
	nova = make([][]int, ordem-1)
	for i = 0; i < ordem-1; i++ {
		nova[i] = make([]int, ordem-1)
	}

	novaI = 0
	for i = 0; i < ordem; i++ {
		if (i < linha) || (i > linha) {
			novaJ = 0
			for j = 0; j < ordem; j++ {
				if (j < coluna) || (j > coluna) {
					nova[novaI][novaJ] = mat[i][j]
					novaJ = novaJ + 1
				}
			}
			novaI = novaI + 1
		}
	}
	return nova
}

func determinanteBaseline(mat [][]int) int {
	var ordem int
	var det int
	var j int
	var sinal int

	ordem = len(mat)
	if ordem == 1 {
		return mat[0][0]
	}
	if ordem == 2 {
		return mat[0][0]*mat[1][1] - mat[0][1]*mat[1][0]
	}

	det = 0
	sinal = 1
	for j = 0; j < ordem; j++ {
		det = det + sinal*mat[0][j]*determinanteBaseline(menorMatriz(mat, 0, j))
		sinal = sinal * (-1)
	}
	return det
}

func determinanteOtimizado(mat [][]int) int {
	var ordem int
	var det int
	var i int
	var j int
	var maiorZerosLinha int
	var maiorZerosColuna int
	var linhaEscolhida int
	var colunaEscolhida int
	var sinal int
	var submat [][]int

	ordem = len(mat)
	if ordem == 1 {
		return mat[0][0]
	}
	if ordem == 2 {
		return mat[0][0]*mat[1][1] - mat[0][1]*mat[1][0]
	}

	zerosLinha = make([]int, ordem)
	zerosColuna = make([]int, ordem)

	for ContI = 0; ContI < ordem; ContI++ {
		zerosLinha[ContI] = 0
		zerosColuna[ContI] = 0
	}

	for ContI = 0; ContI < ordem; ContI++ {
		for ContJ = 0; ContJ < ordem; ContJ++ {
			if mat[ContI][ContJ] == 0 {
				zerosLinha[ContI] = zerosLinha[ContI] + 1
				zerosColuna[ContJ] = zerosColuna[ContJ] + 1
			}
		}
	}

	maiorZerosLinha = zerosLinha[0]
	linhaEscolhida = 0
	for i = 1; i < ordem; i++ {
		if zerosLinha[i] > maiorZerosLinha {
			maiorZerosLinha = zerosLinha[i]
			linhaEscolhida = i
		}
	}

	maiorZerosColuna = zerosColuna[0]
	colunaEscolhida = 0
	for j = 1; j < ordem; j++ {
		if zerosColuna[j] > maiorZerosColuna {
			maiorZerosColuna = zerosColuna[j]
			colunaEscolhida = j
		}
	}

	det = 0
	sinal = 1

	if maiorZerosLinha >= maiorZerosColuna {
		for j = 0; j < ordem; j++ {
			if mat[linhaEscolhida][j] != 0 {
				submat = menorMatriz(mat, linhaEscolhida, j)
				det = det + sinal*mat[linhaEscolhida][j]*determinanteOtimizado(submat)
			}
			sinal = sinal * (-1)
		}
	} else {
		for i = 0; i < ordem; i++ {
			if mat[i][colunaEscolhida] != 0 {
				submat = menorMatriz(mat, i, colunaEscolhida)
				det = det + sinal*mat[i][colunaEscolhida]*determinanteOtimizado(submat)
			}
			sinal = sinal * (-1)
		}
	}
	return det
}

func executaTeste(ordem int) {
	var i int
	var matriz [][]int
	var matrizOtimizada [][]int

	mediaBaseline = 0
	mediaOtimizado = 0

	totalBaseline = 0
	totalOtimizado = 0

	fmt.Printf("\n===== MATRIZ ORDEM %d =====\n", ordem)

	for i = 0; i < 3; i++ {
		fmt.Printf("\n--- Execução %d ---\n", i+1)
		matriz = gerarMatrizAleatoria(ordem)

		fmt.Println("Matriz baseline:")
		imprimeMatriz(matriz)

		tempoInicio = time.Now()
		detBase = determinanteBaseline(matriz)
		tempoFim = time.Now()
		tempoBaseline = tempoFim.Sub(tempoInicio).Nanoseconds()
		fmt.Printf("Tempo(ns): %d\n", tempoBaseline)
		fmt.Printf("Determinante: %d\n", detBase)
		mediaBaseline = mediaBaseline + tempoBaseline
		totalBaseline = totalBaseline + tempoBaseline

		matrizOtimizada = copiaMatriz(matriz)
		fmt.Println("Matriz otimizada:")
		imprimeMatriz(matrizOtimizada)

		tempoInicio = time.Now()
		detOtim = determinanteOtimizado(matrizOtimizada)
		tempoFim = time.Now()
		tempoOtimizado = tempoFim.Sub(tempoInicio).Nanoseconds()
		fmt.Printf("Tempo(ns): %d\n", tempoOtimizado)
		fmt.Printf("Determinante: %d\n", detOtim)
		mediaOtimizado = mediaOtimizado + tempoOtimizado
		totalOtimizado = totalOtimizado + tempoOtimizado

		fmt.Println("---------------------------")
	}

	fmt.Printf("\nMÉDIA TEMPO BASELINE (ordem %d): %d ns\n", ordem, mediaBaseline/3)
	fmt.Printf("MÉDIA TEMPO OTIMIZADO (ordem %d): %d ns\n", ordem, mediaOtimizado/3)
	fmt.Println("===============================")
}

func main() {
	var ordens []int
	var cont int

	ordens = []int{3, 5, 7, 9, 11}

	tempoInicio = time.Now()
	for cont = 0; cont < len(ordens); cont++ {
		executaTeste(ordens[cont])
	}
	tempoFim = time.Now()

	tempoTotal = tempoFim.Sub(tempoInicio).Nanoseconds()

	fmt.Printf("\nTEMPO TOTAL DE EXECUÇÃO (baseline + otimizado): %d ns\n", tempoTotal)
}
