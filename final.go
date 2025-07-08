package main

import (
	"fmt"
	"math/rand"
	"time"
)

func imprimeMatriz(mat [][]int) {
	var contI int
	var contJ int
	for contI = 0; contI < len(mat); contI++ {
		for contJ = 0; contJ < len(mat[0]); contJ++ {
			fmt.Print(mat[contI][contJ], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func iniciaMatrizShuffle(mat [][]int, valorInicial int) {
	var contI int
	var contJ int

	for contI = 0; contI < len(mat); contI++ {
		for contJ = 0; contJ < len(mat[0]); contJ++ {
			mat[contI][contJ] = valorInicial
			valorInicial = valorInicial + 1
		}
	}

	for contI = 0; contI < len(mat)*len(mat); contI++ {
		troca(mat, rand.Intn(len(mat)), rand.Intn(len(mat[0])), rand.Intn(len(mat)), rand.Intn(len(mat[0])))
	}
}

func troca(mat [][]int, indAi int, indAj int, indBi int, indBj int) {
	var temp int
	temp = mat[indAi][indAj]
	mat[indAi][indAj] = mat[indBi][indBj]
	mat[indBi][indBj] = temp
}

func copiaMatrizMaiorParaMenor(maior [][]int, menor [][]int, isqn int, jsqn int) {
	var contAi int
	var contAj int
	var contBi int
	var contBj int
	var temp int
	var numL int
	var numC int

	numL = len(menor)
	numC = len(menor[0])

	contAi = 0
	for contBi = 0; contBi < numL; contBi++ {
		if contAi == isqn {
			contAi = contAi + 1
		}
		contAj = 0
		for contBj = 0; contBj < numC; contBj++ {
			if contAj == jsqn {
				contAj = contAj + 1
			}
			temp = maior[contAi][contAj]
			menor[contBi][contBj] = temp
			contAj = contAj + 1
		}
		contAi = contAi + 1
	}
}

func calculaSinal(indiceL int, indiceC int) int {
	var sinal int
	sinal = -1
	if (indiceL+indiceC)%2 == 0 {
		sinal = 1
	}
	return sinal
}

func detOrdem1(mat [][]int) int {
	return mat[0][0]
}

func detOrdem2(mat [][]int) int {
	var diagonalP int
	var diagonalI int
	diagonalP = mat[0][0] * mat[1][1]
	diagonalI = mat[1][0] * mat[0][1]
	return diagonalP - diagonalI
}

func detBaseline(mat [][]int) int {
	var ordem int
	ordem = len(mat)
	if ordem == 1 {
		return detOrdem1(mat)
	}
	if ordem == 2 {
		return detOrdem2(mat)
	}

	var resposta int
	var contC int
	var matMenor [][]int
	var cont int
	var cofator int
	var sinal int
	var detTemp int

	resposta = 0
	for contC = 0; contC < ordem; contC++ {
		cofator = mat[0][contC]
		sinal = calculaSinal(0, contC)

		matMenor = make([][]int, ordem-1)
		for cont = 0; cont < ordem-1; cont++ {
			matMenor[cont] = make([]int, ordem-1)
		}
		copiaMatrizMaiorParaMenor(mat, matMenor, 0, contC)
		detTemp = detBaseline(matMenor)
		resposta = resposta + (cofator * sinal * detTemp)
	}
	return resposta
}

func contaZerosLinha(mat [][]int, linha int) int {
	var cont int
	var contJ int
	cont = 0
	for contJ = 0; contJ < len(mat[0]); contJ++ {
		if mat[linha][contJ] == 0 {
			cont = cont + 1
		}
	}
	return cont
}

func contaZerosColuna(mat [][]int, coluna int) int {
	var cont int
	var contI int
	cont = 0
	for contI = 0; contI < len(mat); contI++ {
		if mat[contI][coluna] == 0 {
			cont = cont + 1
		}
	}
	return cont
}

func detOtimizado(mat [][]int) int {
	var ordem int
	ordem = len(mat)
	if ordem == 1 {
		return detOrdem1(mat)
	}
	if ordem == 2 {
		return detOrdem2(mat)
	}

	var maxZeros int
	var cont int
	var indice int
	var usarLinha bool
	var zerosLinha int
	var zerosColuna int

	maxZeros = -1
	usarLinha = true

	for cont = 0; cont < ordem; cont++ {
		zerosLinha = contaZerosLinha(mat, cont)
		if zerosLinha > maxZeros {
			maxZeros = zerosLinha
			indice = cont
			usarLinha = true
		}
		zerosColuna = contaZerosColuna(mat, cont)
		if zerosColuna > maxZeros {
			maxZeros = zerosColuna
			indice = cont
			usarLinha = false
		}
	}

	var resposta int
	var contVar int
	var cofator int
	var sinal int
	var matMenor [][]int
	var detTemp int
	var iLinha int
	var iColuna int

	resposta = 0
	for contVar = 0; contVar < ordem; contVar++ {
		if usarLinha {
			cofator = mat[indice][contVar]
			sinal = calculaSinal(indice, contVar)
			iLinha = indice
			iColuna = contVar
		} else {
			cofator = mat[contVar][indice]
			sinal = calculaSinal(contVar, indice)
			iLinha = contVar
			iColuna = indice
		}

		if cofator == 0 {
			continue
		}

		matMenor = make([][]int, ordem-1)
		for cont = 0; cont < ordem-1; cont++ {
			matMenor[cont] = make([]int, ordem-1)
		}
		copiaMatrizMaiorParaMenor(mat, matMenor, iLinha, iColuna)
		detTemp = detOtimizado(matMenor)
		resposta = resposta + (cofator * sinal * detTemp)
	}
	return resposta
}

func main() {
	var ordens [5]int
	ordens[0] = 3
	ordens[1] = 5
	ordens[2] = 7
	ordens[3] = 9
	ordens[4] = 11

	var ordem int
	var rodada int
	var matriz [][]int
	var cont int
	var tempoInicio time.Time
	var tempoFim time.Time
	var duracao time.Duration
	var detBase int
	var detOtim int

	for _, ordem = range ordens {
		fmt.Println("Ordem:", ordem)
		for rodada = 1; rodada <= 3; rodada++ {
			fmt.Println("Rodada:", rodada)
			matriz = make([][]int, ordem)
			for cont = 0; cont < ordem; cont++ {
				matriz[cont] = make([]int, ordem)
			}
			iniciaMatrizShuffle(matriz, 0)
			imprimeMatriz(matriz)

			tempoInicio = time.Now()
			detBase = detBaseline(matriz)
			tempoFim = time.Now()
			duracao = tempoFim.Sub(tempoInicio)
			fmt.Println("Baseline - Det:", detBase, "Tempo(ns):", duracao.Nanoseconds())

			tempoInicio = time.Now()
			detOtim = detOtimizado(matriz)
			tempoFim = time.Now()
			duracao = tempoFim.Sub(tempoInicio)
			fmt.Println("Otimizado - Det:", detOtim, "Tempo(ns):", duracao.Nanoseconds())
			fmt.Println("----------------------------")
		}
		fmt.Println("============================")
	}
}
