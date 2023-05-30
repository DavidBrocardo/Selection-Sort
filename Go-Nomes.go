//SelectionSort.go
//Autores: David Brocardo e Leonardo Balan
package main

import (
   "bufio"
    "fmt"
    "os"
     "strings"
    "time"
  "strconv"
)


// definição da classe pessoa
/*type Pessoa struct {
	codigo uint8
	nome   string
}*/

// Variáveis globais para armazenar o número de comparações e trocas
var comparacoes int
var trocas int



// Algoritmo de ordenação selectionSort

func selectionSort(dados []string , n int ) {

	for i := 0; i < n - 1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
      comparacoes++
      if strings.Compare(dados[j], dados[min]) < 0 {
				min = j
			}
		}
      // Retorna um número negativo se 'dados[j]' vem antes de 'dados[min]' na ordem alfabética
 
    
		if min != i {    
      dados[min], dados[i] = dados[i], dados[min]
      /*aux := dados[min]
      dados[min] = dados[i]
      dados[i] = aux*/
      trocas++
		}
	}
}

func main() {
  tam, err := strconv.Atoi(os.Args[1])
fmt.Println("Valor computado ", tam)
  // abre o arquivo nomes
    file, err := os.Open("nomes 1.txt")
    if err != nil {
        fmt.Println("Erro ao abrir o arquivo:", err)
        return
    }
    defer file.Close() // Fecha o arquivo após terminar a leitura

    // Cria um scanner para ler o arquivo linha por linha
    scanner := bufio.NewScanner(file)

    // Cria um slice de inteiros para armazenar os números do arquivo
    var nomes []string

    
    // Lê as linhas do arquivo até chegar ao final ou ler 1000 números
    for scanner.Scan() {
        // Adiciona a linha do arquivo ao slice 'nomes'
      nome := scanner.Text()
      nomes = append(nomes, nome)

        // Interrompe a leitura após ler 1000 números
        if len(nomes) == tam {
            break
        }
    }

    // Inicia o cálculo do tempo de execução
    start := time.Now()

   
   selectionSort(nomes, len(nomes))

    // Encerra o cálculo do tempo de execução
    end := time.Now()

    
    


    // Imprime o número de trocas e comparações realizadas
    fmt.Println("Número de trocas:", trocas)
    fmt.Println("Número de comparações:", comparacoes)

    // Imprime o tempo de execução
    fmt.Println("Tempo de execução:", end.Sub(start))
fmt.Println("  ")
}
