package expressions

import environment "Server/Environment"

type GetId struct {
	Lin   int
	Col   int
	id 	  string
}

func NewGetId(lin int, col int, id string) GetId {
	exp := GetId{lin, col, id}
	return exp
}

func (p GetId) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	return environment.Symbol{Id: p.id}
}