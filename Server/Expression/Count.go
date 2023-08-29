package expressions

import (
	environment "Server/Environment"
)

type Count struct {
	Lin   int
	Col   int
	id 	  string
}

func NewCount(lin int, col int, id string) Count {
	exp := Count{lin, col, id}
	return exp
}

func (p Count) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	result := env.(environment.Environment).GetVector(p.id)
	return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.INTEGER, Valor: len(result.Valor)}
}