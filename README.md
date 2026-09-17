# Calculadora RPC

Calculadora distribuída em Go utilizando o pacote nativo `net/rpc`. O projeto é composto por um servidor que expõe operações aritméticas via RPC e um cliente de linha de comando que consome esse serviço.

## Sumário

- [Visão geral](#visão-geral)
- [Arquitetura](#arquitetura)
- [Pré-requisitos](#pré-requisitos)
- [Instalação](#instalação)
- [Uso](#uso)
- [Operações suportadas](#operações-suportadas)
- [Tratamento de erros](#tratamento-de-erros)
- [Estrutura do projeto](#estrutura-do-projeto)
- [Licença](#licença)

## Visão geral

O servidor registra o tipo `Calculadora`, que expõe métodos remotos de soma, subtração, multiplicação e divisão. O cliente recebe os operandos e o operador via argumentos de linha de comando, conecta-se ao servidor por TCP e invoca a operação correspondente.

## Arquitetura

```
client/main.go  ──── RPC (TCP:1234) ────▶  server/main.go
                                                 │
                                                 ▼
                                        operations.Calculadora
                                          ├─ Soma
                                          ├─ Subtracao
                                          ├─ Multiplicacao
                                          └─ Divisao
```

Os tipos de requisição e resposta (`Numbers`, `ResultadoInt`, `ResultadoFloat`) ficam no pacote `calc`, compartilhado entre cliente e servidor.

## Pré-requisitos

- [Go](https://go.dev/dl/) 1.26 ou superior

## Instalação

```bash
git clone https://github.com/<seu-usuario>/calculadora_RPC.git
cd calculadora_RPC
go mod init calculadora_RPC
```

## Uso

### 1. Iniciar o servidor

```bash
go run server/main.go
```

O servidor sobe na porta `1234` e aguarda conexões.

### 2. Executar o cliente

```bash
go run client/main.go <host> <x> <operador> <y>
```

| Argumento   | Descrição                                  |
|-------------|---------------------------------------------|
| `host`      | Endereço do servidor (ex.: `localhost`)      |
| `x`         | Primeiro operando (inteiro)                  |
| `operador`  | Operação: `+`, `-`, `x` ou `/`               |
| `y`         | Segundo operando (inteiro)                   |

#### Exemplos

```bash
go run client/main.go localhost 10 + 5
# Result:15

go run client/main.go localhost 10 - 5
# Result:5

go run client/main.go localhost 10 x 5
# Result:50

go run client/main.go localhost 10 / 5
# Result:2.000000
```

## Operações suportadas

| Operador | Método RPC               | Tipo de retorno |
|----------|---------------------------|-----------------|
| `+`      | `Calculadora.Soma`         | `int`           |
| `-`      | `Calculadora.Subtracao`    | `int`           |
| `x`      | `Calculadora.Multiplicacao`| `int`           |
| `/`      | `Calculadora.Divisao`      | `float64`       |

## Tratamento de erros

- Argumentos ausentes ou inválidos (operandos não numéricos) encerram o cliente com mensagem de erro.
- Operador desconhecido é rejeitado pelo cliente antes de chamar o servidor.
- Divisão por zero retorna um erro tratado pelo servidor (`Calculadora.Divisao`), exibido no cliente sem interromper sua execução.

## Estrutura do projeto

```
calculadora_RPC/
├── calc/
│   └── calc.go            # Tipos compartilhados (Numbers, ResultadoInt, ResultadoFloat)
├── client/
│   └── main.go             # Cliente RPC (linha de comando)
├── operations/
│   ├── calculadora.go       # Struct Calculadora
│   ├── soma.go
│   ├── subtracao.go
│   ├── multiplicacao.go
│   └── divisao.go
├── server/
│   └── main.go              # Servidor RPC
└── go.mod
```

