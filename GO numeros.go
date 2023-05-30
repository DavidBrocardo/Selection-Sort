
package main

import (
   "bufio"
    "fmt"
    "os"
    "strconv"
    "time"
)



// Variáveis globais para armazenar o número de comparações e trocas
var comparacoes int
var trocas int



// Algoritmo de ordenação selectionSort
func selectionSort(dados []int , n int ) {

	for i := 0; i < n - 1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
      comparacoes++
      if dados[j] < dados[min] {
				min = j
			}
		}
    
		if min != i {     
      aux := dados[min]
      dados[min] = dados[i]
      dados[i] = aux
      trocas++
		}
	}
}

func main() {
  tam, err := strconv.Atoi(os.Args[1])
fmt.Println("Valor computado ", tam)
    // Abre o arquivo 'numeros1.txt' para leitura
    file, err := os.Open("numeros 1.txt")
    if err != nil {
        fmt.Println("Erro ao abrir o arquivo:", err)
        return
    }
    defer file.Close() // Fecha o arquivo após terminar a leitura
  
    // Cria um scanner para ler o arquivo linha por linha
    scanner := bufio.NewScanner(file)

    // Cria um slice de inteiros para armazenar os números do arquivo
    var nums []int

    // Lê as linhas do arquivo até chegar ao final ou ler 1000 números
 
    for scanner.Scan() {
        // Converte a linha do arquivo em um inteiro e adiciona ao slice 'nums'
        num, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Println("Erro ao converter o número:", err)
            return
        }
        nums = append(nums, num)

        // Interrompe a leitura após ler 1000 números
        if len(nums) == tam {
            break
        }
    }

    // Inicia o cálculo do tempo de execução
    start := time.Now()

    // Chama a função 'quickSort' para classificar os números em ordem crescente
    selectionSort(nums, len(nums))

    // Encerra o cálculo do tempo de execução
    end := time.Now()

    // Imprime os números ordenados no console
    //fmt.Println("Números ordenados:", nums)

    // Imprime o número de trocas e comparações realizadas
    fmt.Println("Número de trocas:", trocas)
    fmt.Println("Número de comparações:", comparacoes)

    // Imprime o tempo de execução
    fmt.Println("Tempo de execução:", end.Sub(start))
    fmt.Println("  ")
}
