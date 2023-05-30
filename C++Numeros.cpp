// Selection Sort
// Autores: David Brocardo e Leonardo Balan
#include <ctime>
#include <iostream>
using namespace std;
// variavel grobal
int comp=0;
float troca = 0;

// Funcao SelectionSort para inteiros
template <class T> void SelectionSort(T dados[], int n) {
  for (int i = 0; i < (n - 1); i++) {
    // encontra o menor elemento
    int min = i;
    for (int j = i + 1; j < n; j++) {
      comp++;
      if (dados[j] < dados[min])
        min = j;
    }
    // Troca a posição atual com o menor elemento do vetor
    if (min != i) {
      T aux = dados[min];
      dados[min] = dados[i];
      dados[i] = aux;
      troca++;
    }
    
  }
}

int main(int argc, char *argv[]) {
  FILE *in; // ponteiro para orquivo de entrada
  // declaração de variáveis
  int n = 0;
  int valor = 0;
  char nome;
  n = stoi(argv[1]);
  // Imprime o valor computado
  std::cout << "Valor computado " << n << std::endl;
  int *V = new int[n];

  // abertura dos arquivos
  in = fopen("numeros 1.txt", "r");

  if (!in) {
    printf("Erro : não foi possível abrir o arquivo\n");
    exit(1);
  }
  // Para o Teste de "numeros 1.txt"
  ;

  for (int i = 0; i < n; i++) {
    fscanf(in, "%d", &valor);
    V[i] = valor;
  }

  // escreve(V, n);
  // cout << "\n\nOrdenando o vetor com :\n ";
  // Início da contagem do tempo
  clock_t inicio = clock(); // Início da contagem de tempo

  SelectionSort(V, n);

// Cálculo do tempo decorrido em segundos
  clock_t fim = clock(); // Fim da contagem de tempo
  double duracao = (double)(fim - inicio) /                 CLOCKS_PER_SEC; // Calculando a duração em segundos
  cout << endl;
  

  // Imprime o tempo de execução em segundos
  std::cout << "Tempo de execução: " << duracao << " segundos" << std::endl;
  // Imprime o numero de trocas e comparações
  cout << "Trocas: " << troca << " | Comparações: " << comp << endl;
  delete[] V;

  cout << endl;

  // fecha os arquivos
  fclose(in);

  return 0;
}
