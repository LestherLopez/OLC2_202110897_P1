package instructions

import (
	environment "Server/Environment"
)

type Continue struct {
	Lin   int
	Col   int
	
}

func NewContinue(lin int, col int) Continue {
	expresion := Continue{lin, col}
	return expresion
}

func (p Continue) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	
		//return sin expresion
	return environment.Symbol{Lin: p.Lin, Col: p.Col, Id: "",Tipo: environment.NULL, Valor: nil, Mutable: true, Transfer: environment.CONTINUE}
	
}