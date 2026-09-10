package main

import (
	"log"
	"net/rpc"
	"fmt"
	"calculadora_RPC/calc"
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
	args := calc.Numbers{A:5, B:0}

	
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

	fmt.Println(replyMultiplicacao.ResultadoInt, replySoma.ResultadoInt, replySubtracao.ResultadoInt, replyDivisao.ResultadoFloat)
}