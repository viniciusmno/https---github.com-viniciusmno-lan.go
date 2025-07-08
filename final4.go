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
}

func iniciaMatrizAleatoria(mat [][]int, ordem int) {
	var contI int
	var contJ int
	var maxValor int
	maxValor = ordem * ordem
	for contI = 0; contI < ordem; contI++ {
		for contJ = 0; contJ < ordem; contJ++ {
			mat[contI][contJ] = rand.Intn(maxValor + 1)
		}
	}
}

func copiaMatriz(orig [][]int) [][]int {
	var ordem int
	var contI int
	var contJ int
	ordem = len(orig)
	var copia [][]int
	copia = make([][]int, ordem)
	for contI = 0; contI < ordem; contI++ {
		copia[contI] = make([]int, ordem)
	}
	for contI = 0; contI < ordem; contI++ {
		for contJ = 0; contJ < ordem; contJ++ {
			copia[contI][contJ] = orig[contI][contJ]
		}
	}
	return copia
}

func verificaQuadradaOrdem(mat [][]int) (bool, int) {
	var numLinhas int
	var numColunas int
	var ehQuadrada bool
	numLinhas = len(mat)
	numColunas = len(mat[0])
	ehQuadrada = false
	if numLinhas == numColunas {
		ehQuadrada = true
	}
	return ehQuadrada, numLinhas
}

func calculaSinal(indiceL int, indiceC int) int {
	var sinal int
	sinal = -1
	if ((indiceL + indiceC) % 2) == 0 {
		sinal = 1
	}
	return sinal
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

func detOrdem1(mat [][]int) int {
	return mat[0][0]
}

func detOrdem2(mat [][]int) int {
	var diagonalP int
	var diagonalI int
	diagonalP = mat[0][0] * mat[1][1]
	diagonalI = mat[1][0] * mat[0][1]
	return (diagonalP - diagonalI)
}

func detOrdemN(mat [][]int) int {
	var sinal, cofator, detTemp, resposta, contL, contC, numL, numC, cont int
	var matMenor [][]int
	numL = len(mat)
	numC = len(mat[0])

	resposta = 0
	contL = 0
	for contC = 0; contC < numC; contC++ {
		cofator = mat[contL][contC]
		sinal = calculaSinal(contL, contC)
		matMenor = make([][]int, numL-1)
		for cont = 0; cont < (numL - 1); cont++ {
			matMenor[cont] = make([]int, numC-1)
		}
		copiaMatrizMaiorParaMenor(mat, matMenor, contL, contC)
		detTemp = determinante(matMenor)
		resposta = resposta + (cofator * sinal * detTemp)
	}
	return resposta
}

func determinante(mat [][]int) int {
	var ordem int
	var ehQuadrada bool
	var det int
	ehQuadrada, ordem = verificaQuadradaOrdem(mat)
	det = 0
	if ehQuadrada {
		fmt.Println("Ordem", ordem)
		switch ordem {
		case 1:
			det = detOrdem1(mat)
		case 2:
			det = detOrdem2(mat)
		default:
			det = detOrdemN(mat)
		}
		imprimeMatriz(mat)
		fmt.Println("Det", det)
	} else {
		fmt.Println("Matriz nao eh quadrada!! retornando 0")
	}
	return det
}

func linhaOuColunaMaisZeros(mat [][]int) (bool, int) {
	var contI, contJ, contZero, maxZeros, indiceMax int
	var ehLinha bool

	maxZeros = -1
	ehLinha = true
	for contI = 0; contI < len(mat); contI++ {
		contZero = 0
		for contJ = 0; contJ < len(mat[0]); contJ++ {
			if mat[contI][contJ] == 0 {
				contZero = contZero + 1
			}
		}
		if contZero > maxZeros {
			maxZeros = contZero
			indiceMax = contI
			ehLinha = true
		}
	}
	for contJ = 0; contJ < len(mat[0]); contJ++ {
		contZero = 0
		for contI = 0; contI < len(mat); contI++ {
			if mat[contI][contJ] == 0 {
				contZero = contZero + 1
			}
		}
		if contZero > maxZeros {
			maxZeros = contZero
			indiceMax = contJ
			ehLinha = false
		}
	}
	return ehLinha, indiceMax
}

func detOrdemNOtimizado(mat [][]int) int {
	var sinal, cofator, detTemp, resposta, contL, contC, numL, numC, cont int
	var matMenor [][]int
	var ehLinha bool
	var indiceMax int

	numL = len(mat)
	numC = len(mat[0])
	resposta = 0
	ehLinha, indiceMax = linhaOuColunaMaisZeros(mat)

	if numL == 1 {
		resposta = mat[0][0]
	}
	if numL == 2 {
		resposta = detOrdem2(mat)
	}
	if numL > 2 {
		if ehLinha == true {
			contL = indiceMax
			for contC = 0; contC < numC; contC++ {
				cofator = mat[contL][contC]
				sinal = calculaSinal(contL, contC)
				matMenor = make([][]int, numL-1)
				for cont = 0; cont < (numL - 1); cont++ {
					matMenor[cont] = make([]int, numC-1)
				}
				copiaMatrizMaiorParaMenor(mat, matMenor, contL, contC)
				detTemp = detOrdemNOtimizado(matMenor)
				resposta = resposta + (cofator * sinal * detTemp)
			}
		} else {
			contC = indiceMax
			for contL = 0; contL < numL; contL++ {
				cofator = mat[contL][contC]
				sinal = calculaSinal(contL, contC)
				matMenor = make([][]int, numL-1)
				for cont = 0; cont < (numL - 1); cont++ {
					matMenor[cont] = make([]int, numC-1)
				}
				copiaMatrizMaiorParaMenor(mat, matMenor, contL, contC)
				detTemp = detOrdemNOtimizado(matMenor)
				resposta = resposta + (cofator * sinal * detTemp)
			}
		}
	}
	return resposta
}

func medirTempo(mat [][]int, fn func([][]int) int) int64 {
	var inicio time.Time
	var fim time.Time
	var duracao int64
	inicio = time.Now()
	_ = fn(mat)
	fim = time.Now()
	duracao = fim.Sub(inicio).Nanoseconds()
	return duracao
}
