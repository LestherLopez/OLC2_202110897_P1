package expressions

import environment "Server/Environment"

type Access struct {
	Lin   int
	Col   int
	id 	  string
}

func NewAccess(lin int, col int, id string) Access {
	exp := Access{lin, col, id}
	return exp
}

func (p Access) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	result := env.(environment.Environment).GetVariable(p.id)
	return result
}