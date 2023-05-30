// Selection Sort
// Autores: David Brocardo e Leonardo Balan
#include <algorithm>
#include <ctime>
#include <fstream>
#include <iostream>
#include <string>
#include <vector>
using namespace std;
// variavel grobal
int comp = 0;
float troca = 0;

// Funcao SelectionSort para char
void selectionSort(std::vector<std::string> &dados) {
  //(std::string dados[], int n) {
  int n = dados.size();
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
        std::swap(dados[i], dados[min]);
        troca++;
    }
    //comp++;
  }
}
int main(int argc, char *argv[]) {
  FILE *in; // ponteiro para orquivo de entrada
  // declaração de variáveis
  int n = 0;
  int valor = 0;
  char nome;
  n = stoi(argv[1]);
  std::cout << "Valor computado " << n << std::endl;
  // cout << "\n\t\t\t Selection Sort\n\n";

  // cout << "Ordenando um vetor de  " << n << " nomes\n";

  std::ifstream inputFile("nomes 1.txt");
  if (!inputFile) {
    std::cerr << "Erro ao abrir o arquivo de entrada.\n";
    return 1;
  }

  std::vector<std::string> names;
  std::string name;

  // while (n < total) {
  //   n++;
  for (int i = 0; i < n; i++) {
    std::getline(inputFile, name);
    names.push_back(name);
    // cout<< name << endl;
  }
  inputFile.close();

  clock_t inicio = clock(); // Início da contagem de tempo
  // SelectionSort(V, n);
  selectionSort(names);

  // Cálculo do tempo decorrido em segundos
  clock_t fim = clock(); // Fim da contagem de tempo
  double duracao = (double)(fim - inicio) /
                   CLOCKS_PER_SEC; // Calculando a duração em segundos
  cout << endl;

  // Imprime o tempo de execução em segundos
  cout << "Tempo de execução: " << duracao << " segundos" << endl;
  // Imprime o numero de trocas e comparações
  cout << "Trocas: " << troca << " | Comparações: " << comp << endl;
  // delete[] VC;

  cout << endl;

  // fecha os arquivos
  // fclose(in);

  return 0;
}
