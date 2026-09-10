package main

import (
	"calculadora_RPC/operations"
	"log"
	"net"
	"net/rpc"
)

func main() {
	calculadora := new(operations.Calculadora)
	err := rpc.Register(calculadora)
	if err != nil {
		log.Fatal("call failed:", err)
	}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Erro ao identificar a porta:", err)
	}

	rpc.Accept(listener)
}
