package expressions

import environment "Server/Environment"

type Parameters struct {
	Lin   int
	Col   int
	Type  environment.TipoExpresion
	Id	  string
}

func NewParameters(lin int, col int, tipo environment.TipoExpresion, id string) Parameters {
	exp := Parameters{lin, col, tipo, id}
	return exp
}

func (p Parameters) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	return environment.Symbol{
		Lin:   p.Lin,
		Col:   p.Col,
		Tipo:  p.Type,
		Id:	   p.Id,
	}
}