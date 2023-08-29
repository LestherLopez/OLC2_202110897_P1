package instructions

import (
	environment "Server/Environment"
)

type RemoveLast struct {
	Lin   int
	Col   int
	id_var string
	
}
//lin int, col int, id_var string, type_var environment.TipoExpresion, valor interfaces.Expression, constant bool
func NewRemoveLast(lin int, col int, id_var string) RemoveLast {
	return RemoveLast{lin, col, id_var}
}

func (p RemoveLast) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	
	
	vector := env.(environment.Environment).GetVector(p.id_var)

	
	//value:=environment.Symbol{Lin: valor.Lin, Col: valor.Col, Id: valor.Id, Tipo: valor.Tipo, Valor: valor.Valor, Transfer: environment.NULL}
	newVector := vector.Valor[:len(vector.Valor)-1]
	vector.Valor = newVector
	env.(environment.Environment).SetVector(p.id_var, vector)

	
	return nil
}