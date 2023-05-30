// Selection Sort com duas chaves
// Autores: David Brocardo e Leonardo Balan

#include <algorithm>
#include <ctime>
#include <fstream>
#include <iostream>
#include <sstream>
#include <vector>
// variavel grobal
using namespace std;
int comp, troca = 0;

struct Pessoa {
  int codigo;
  string nome;
};

void selectionSort(std::vector<Pessoa> &pessoas) {
  int n = pessoas.size();
  for (int i = 0; i < n - 1; i++) {
    int min = i;
    for (int j = i + 1; j < n; j++) {
      comp++;
      if (pessoas[j].codigo < pessoas[min].codigo ||
          (pessoas[j].codigo == pessoas[min].codigo &&
           pessoas[j].nome < pessoas[min].nome)) {
        min = j;
      }
    }
    
    if (min != i) {
      swap(pessoas[i], pessoas[min]);
      troca++;
    }
  }
}

int main(int argc, char *argv[]) {
  vector<Pessoa> pessoas;
  ifstream arquivo("pessoa 1.txt");
  if (!arquivo) {
    std::cerr << "Erro ao abrir o arquivo.\n";
    return 1;
  }

  string linha;
  int total = 0;
  int n = 0;
  total = stoi(argv[1]);
  std::cout << "Valor computado " << total << std::endl;
  while (n < total) {
    n++;
    getline(arquivo, linha);
    stringstream ss(linha);
    Pessoa pessoa;
    ss >> pessoa.codigo >> pessoa.nome;
    pessoas.push_back(pessoa);
  }
  arquivo.close();

  /* std::cout << "Pessoas antes da ordenação:\n";
   for (const auto& pessoa : pessoas) {
       std::cout << "Código: " << pessoa.codigo << ", Nome: " << pessoa.nome <<
   "\n";
   }*/
  clock_t inicio = clock(); // Início da contagem de tempo
  selectionSort(pessoas);

  // Cálculo do tempo decorrido em segundos
  clock_t fim = clock(); // Fim da contagem de tempo
  double duracao = (double)(fim - inicio) /
                   CLOCKS_PER_SEC; // Calculando a duração em segundos
  cout << endl;

  // Imprime o tempo de execução em segundos
  cout << "Tempo de execução: " << duracao << " segundos" << endl;
  // Imprime o numero de trocas e comparações
  cout << "Trocas: " << troca << " | Comparações: " << comp << endl;
  return 0;
}
