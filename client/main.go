package main

import (
	"calculadora_RPC/calc"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
)

func main() {
	var replySoma calc.ResultadoInt	
	var replySubtracao calc.ResultadoInt	
	var replyMultiplicacao calc.ResultadoInt
	var replyDivisao calc.ResultadoFloat

	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error connecting to the server", err)
	}

	if len(os.Args) < 4 {
		log.Fatal("Erro ao passar parametros")
	}

	a, errA := strconv.Atoi(os.Args[2])

	if errA != nil {
		log.Fatal("erro: valores incorretos")
	}

	b, errB := strconv.Atoi(os.Args[3])

	if errB != nil {
		log.Fatal("erro: valores incorretos")
	}
	args := calc.Numbers{A:a, B:b}

	
	err = client.Call("Calculadora.Soma", args, &replySoma)
	if err != nil {
		log.Fatal("erro no processo de operação:", err)
	}

	err = client.Call("Calculadora.Subtracao", args, &replySubtracao)
	if err != nil {
		log.Fatal("erro no processo de operação:", err)
	}

	err = client.Call("Calculadora.Multiplicacao", args, &replyMultiplicacao)
	if err != nil {
		log.Fatal("erro no processo de operação:", err)
	}

	err = client.Call("Calculadora.Divisao", args, &replyDivisao)
	if err != nil {
		fmt.Println("Aviso:", err)
	} 

}