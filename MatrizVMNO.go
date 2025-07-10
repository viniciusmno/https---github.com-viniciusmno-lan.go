package main

import (
	"fmt"
	"math/rand"
	"time"
)

func imprimeMatriz(mat [][]int) {
	var ContI int
	var ContJ int
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
	var ContI int
	var ContJ int
	var novaI int
	var novaJ int

	ordem = len(mat)
	nova = make([][]int, ordem-1)
	for ContI = 0; ContI < ordem-1; ContI++ {
		nova[ContI] = make([]int, ordem-1)
	}

	novaI = 0
	for ContI = 0; ContI < ordem; ContI++ {
		if (ContI < linha) || (ContI > linha) {
			novaJ = 0
			for ContJ = 0; ContJ < ordem; ContJ++ {
				if (ContJ < coluna) || (ContJ > coluna) {
					nova[novaI][novaJ] = mat[ContI][ContJ]
					novaJ = novaJ + 1
				}
			}
			novaI = novaI + 1
		}
	}
	return nova
}
func sinal(linha int, coluna int) int {
	if (linha+coluna)%2 == 0 {
		return 1
	} else {
		return -1
	}
}

func determinanteBaseline(mat [][]int) int {
	var ordem int
	var det int
	var ContJ int

	ordem = len(mat)
	if ordem == 1 {
		return mat[0][0]
	}
	if ordem == 2 {
		return mat[0][0]*mat[1][1] - mat[0][1]*mat[1][0]
	}

	det = 0
	for ContJ = 0; ContJ < ordem; ContJ++ {
		det = det + sinal(0, ContJ)*mat[0][ContJ]*determinanteBaseline(menorMatriz(mat, 0, ContJ))
	}
	return det
}
func determinanteOtimizado(mat [][]int) int {
	var ordem int
	var det int
	var ContI int
	var ContJ int
	var maiorZerosLinha int
	var maiorZerosColuna int
	var linhaEscolhida int
	var colunaEscolhida int
	var submat [][]int

	var zerosLinha []int
	var zerosColuna []int

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
	for ContI = 1; ContI < ordem; ContI++ {
		if zerosLinha[ContI] > maiorZerosLinha {
			maiorZerosLinha = zerosLinha[ContI]
			linhaEscolhida = ContI
		}
	}

	maiorZerosColuna = zerosColuna[0]
	colunaEscolhida = 0
	for ContJ = 1; ContJ < ordem; ContJ++ {
		if zerosColuna[ContJ] > maiorZerosColuna {
			maiorZerosColuna = zerosColuna[ContJ]
			colunaEscolhida = ContJ
		}
	}

	det = 0

	if maiorZerosLinha >= maiorZerosColuna {
		for ContJ = 0; ContJ < ordem; ContJ++ {
			if mat[linhaEscolhida][ContJ] > 0 || mat[linhaEscolhida][ContJ] < 0 {
				submat = menorMatriz(mat, linhaEscolhida, ContJ)
				det = det + sinal(linhaEscolhida, ContJ)*mat[linhaEscolhida][ContJ]*determinanteOtimizado(submat)
			}
		}
	} else {
		for ContI = 0; ContI < ordem; ContI++ {
			if mat[ContI][colunaEscolhida] > 0 || mat[ContI][colunaEscolhida] < 0 {
				submat = menorMatriz(mat, ContI, colunaEscolhida)
				det = det + sinal(ContI, colunaEscolhida)*mat[ContI][colunaEscolhida]*determinanteOtimizado(submat)
			}
		}
	}
	return det
}
func imprimaExperimento(ordem int) {
	var ContI int
	var matriz [][]int
	var matrizOtimizada [][]int
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

	mediaBaseline = 0
	mediaOtimizado = 0

	totalBaseline = 0
	totalOtimizado = 0

	fmt.Printf("\n===== MATRIZ ORDEM %d =====\n", ordem)

	for ContI = 0; ContI < 3; ContI++ {
		fmt.Printf("\n--- EXPERIMENTO %d ---\n", ContI+1)
		fmt.Println()
		matriz = gerarMatrizAleatoria(ordem)

		fmt.Println("Matriz baseline:")
		imprimeMatriz(matriz)

		tempoInicio = time.Now()
		detBase = determinanteBaseline(matriz)
		tempoFim = time.Now()
		tempoBaseline = tempoFim.Sub(tempoInicio).Nanoseconds()
		fmt.Printf("Tempo: %d\n", tempoBaseline)
		fmt.Printf("Determinante: %d\n", detBase)
		mediaBaseline = mediaBaseline + tempoBaseline
		totalBaseline = totalBaseline + tempoBaseline
		fmt.Println()
		matrizOtimizada = copiaMatriz(matriz)
		fmt.Println("Matriz otimizada:")
		imprimeMatriz(matrizOtimizada)

		tempoInicio = time.Now()
		detOtim = determinanteOtimizado(matrizOtimizada)
		tempoFim = time.Now()
		tempoOtimizado = tempoFim.Sub(tempoInicio).Nanoseconds()
		fmt.Printf("Tempo: %d\n", tempoOtimizado)
		fmt.Printf("Determinante: %d\n", detOtim)
		mediaOtimizado = mediaOtimizado + tempoOtimizado
		totalOtimizado = totalOtimizado + tempoOtimizado
		fmt.Println()
	}

	fmt.Printf("\nMÉDIA DO TEMPO DA MATRIZ BASELINE DE ORDEM %d: %d ns\n", ordem, mediaBaseline/3)
	fmt.Printf("MÉDIA DO TEMPO DA MATRIZ OTIMIZADA DE ORDEM %d: %d ns\n", ordem, mediaOtimizado/3)
	fmt.Println()
}
func main() {
	var ordens []int
	var cont int
	var tempoInicio time.Time
	var tempoFim time.Time
	var tempoTotal int64

	ordens = []int{3, 5, 7, 9, 11}
	tempoInicio = time.Now()
	for cont = 0; cont < len(ordens); cont++ {
		imprimaExperimento(ordens[cont])
	}
	tempoFim = time.Now()

	tempoTotal = tempoFim.Sub(tempoInicio).Nanoseconds()

	fmt.Printf("\nTempo Total: %d ns\n", tempoTotal)
}
