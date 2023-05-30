package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)
// variaveis grobais

var execucoes int
var trocas int
type Pessoa struct {
	Codigo int
	Nome   string
}

func main() {
   tam, err := strconv.Atoi(os.Args[1])
fmt.Println("Valor computado ", tam)
  
	// Abrir o arquivo
	file, err := os.Open("pessoa 1.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	// Criar um slice de Pessoas
	pessoas := make([]Pessoa, 0)

	// Processar todas as linhas do arquivo
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) >= 2 {
			codigo, err := strconv.Atoi(fields[0])
			if err != nil {
				fmt.Println("Erro ao converter código para inteiro:", err)
				return
			}
			nome := strings.Join(fields[1:], " ")
			pessoas = append(pessoas, Pessoa{Codigo: codigo, Nome: nome})
		}
    if tam == len(pessoas) {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler as linhas do arquivo:", err)
		return
	}

	// Criar o arquivo para salvar os resultados
	resultFile, err := os.Create("resultados.txt")
	if err != nil {
		fmt.Println("Erro ao criar o arquivo de resultados:", err)
		return
	}
	defer resultFile.Close()

	// Executar o algoritmo de ordenação em lotes de 1000 casos até 100000
	/*for {
		if tam == len(pessoas) {
			break
		}*/

		// Criar uma cópia do slice de pessoas para cada teste
		pessoasTeste := make([]Pessoa, tam)
		copy(pessoasTeste, pessoas[:tam])

		// Variáveis para contagem de trocas e execuções


		// Medir o tempo de execução
		startTime := time.Now()

		// Ordenar as pessoas utilizando o selection sort
    selectionSort(pessoasTeste)

		// Calcular o tempo de execução
		endTime := time.Now()
		executionTime := endTime.Sub(startTime)

    // Imprimir as pessoas ordenadas
		//for _, pessoa := range pessoasTeste {
		//	fmt.Printf("%d %s\n", pessoa.Codigo, pessoa.Nome)
	//	}

		fmt.Println("Tempo de execução:", executionTime)
		 fmt.Println("Número de trocas: ", trocas)
		 fmt.Println("Número de execuções: ", execucoes)
		fmt.Println("   ")
	//}
}


func selectionSort(pessoas []Pessoa ) {
  n := len(pessoas)
  
	for i := 0; i < n - 1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
      execucoes++
     
       if pessoas[j].Codigo < pessoas[min].Codigo || pessoas[j].Codigo == pessoas[min].Codigo && strings.Compare(pessoas[j].Nome, pessoas[min].Nome) < 0 {
      	min = j
			}
		}  
 
    
		if min != i {    
      pessoas[min], pessoas[i] = pessoas[i], pessoas[min]
      trocas++
		}
	}
}
