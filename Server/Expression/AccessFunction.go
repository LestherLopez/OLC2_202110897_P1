package expressions

import environment "Server/Environment"

type AccessFunction struct {
	Lin   int
	Col   int
	id 	  string
}

func NewAccessFunction(lin int, col int, id string, ) AccessFunction {
	exp := AccessFunction{lin, col, id}
	return exp
}

func (p AccessFunction) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	result := env.(environment.Environment).GetFunction(p.id)
	return result.Return
}