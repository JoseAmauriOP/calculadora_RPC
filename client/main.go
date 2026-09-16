package main

import (
	"calculadora_RPC/calc"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strconv"
)

func call_calc(host string,  x int, y int, operador string){
	var replyOperacaoInt calc.ResultadoInt
	var replyOperacaoDivisao calc.ResultadoFloat
	var args calc.Numbers
	
	endereco := host + ":1234"
	client, err := rpc.Dial("tcp", endereco)
	if err != nil {
		log.Fatal("Error connecting to the server", err)
	}

	args.A = x
	args.B = y

	if operador == "/"{
		err = client.Call("Calculadora.Divisao", args, &replyOperacaoDivisao)
		if err != nil {
			fmt.Println("Aviso:", err)
		}
		fmt.Printf("Result:%f\n", replyOperacaoDivisao.ResultadoFloat);
	} else {
		switch operador {
		case "+":
			err = client.Call("Calculadora.Soma", args, &replyOperacaoInt)
			if err != nil {
				log.Fatal("erro no processo de operação:", err)
			}
			
		case "-":
			err = client.Call("Calculadora.Subtracao", args, &replyOperacaoInt)
			if err != nil {
				log.Fatal("erro no processo de operação:", err)
			}
			
		case "x":
			err = client.Call("Calculadora.Multiplicacao", args, &replyOperacaoInt)
			if err != nil {
				log.Fatal("erro no processo de operação:", err)
			}
			
		default:
			log.Fatal("operador invalido")
		}

		fmt.Printf("Result:%d\n", replyOperacaoInt.ResultadoInt);
	}

}

func main() {
	if len(os.Args) < 4 {
		log.Fatal("Erro ao passar parametros")
	}

	a, errA := strconv.Atoi(os.Args[2])

	if errA != nil {
		log.Fatal("erro: valores incorretos")
	}

	b, errB := strconv.Atoi(os.Args[4])

	if errB != nil {
		log.Fatal("erro: valores incorretos")
	}
	
	call_calc(os.Args[1], a, b, os.Args[3])
}