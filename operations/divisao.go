package operations

import (
	"calculadora_RPC/calc"
	"errors"
)
func (c *Calculadora) Divisao(args calc.Numbers, reply *calc.ResultadoFloat) error {
	if args.B == 0 {
		return errors.New("Erro: Divisão por zero")
	}
	
	reply.ResultadoFloat = float64(args.A) / float64(args.B)
	
	return nil
}