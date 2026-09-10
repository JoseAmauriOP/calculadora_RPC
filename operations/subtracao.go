package operations

import "calculadora_RPC/calc"

func (t *Calculadora) Subtracao(args calc.Numbers, reply *calc.ResultadoInt) error {
	reply.ResultadoInt = args.A - args.B

	return nil
}