# Aprendendo Go

## ⚙️ Descrição do repositório

Este é um simples projeto em Go que demonstra o uso de arrays, slices, maps e structs. O código inclui funções para criar e manipular cada um desses tipos de dados.

### 🛠️ Estrutura do Código

O código é composto de um pacote principal e uma estrutura User definida pelo usuário, além de várias funções.

### ⚡️ Struct

A struct `User` possui dois campos: `name` do tipo `string` e `age` do tipo `uint8.`

```Go
type User struct {
	name string
	age uint8
}
```

### Funções

Existem três funções principais: `arrayAndSlice()`, `mapObj()`, e `userStruct()`.

`arrayAndSlice()`: Esta função demonstra como criar e manipular um slice em Go. Ela inicia um slice com três elementos, depois adiciona mais três elementos. Em cada etapa, ela imprime o slice, seu tamanho e sua capacidade.

`mapObj()`: Esta função cria um map simples com uma chave "name" e um valor "luis", e depois imprime o map.

`userStruct()`: Esta função cria uma instância da struct User com nome "Wagner" e idade 29, e depois imprime a struct.

`main()`: Esta é a função principal que chama as outras funções.

```go
func main() {
	arrayAndSlice()
	mapObj()
	userStruct()
}
```

### Execução

Para executar o código, use o seguinte comando:

```go
go run main.go
```

### Contribuição

Sinta-se à vontade para contribuir com este projeto, fazendo um fork do repositório e abrindo uma solicitação de pull com suas mudanças.
