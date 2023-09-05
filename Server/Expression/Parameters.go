package expressions

import environment "Server/Environment"

type Parameters struct {
	Lin   int
	Col   int
	Type  environment.TipoExpresion
	firs_Id	  string
	second_Id string
}

func NewParameters(lin int, col int, tipo environment.TipoExpresion, first_id string, second_id string) Parameters {
	exp := Parameters{lin, col, tipo, first_id, second_id}
	return exp
}

func (p Parameters) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	return environment.Symbol{
		Lin:   p.Lin,
		Col:   p.Col,
		Tipo:  p.Type,
		Id:	   p.firs_Id,
	}
}