package expressions

import (
	environment "Server/Environment"
)

type AccessStruct struct {
	Lin    int
	Col    int
	Struct string
	Id     string
}

func NewAccessStruct(lin int, col int, str string, id string) AccessStruct {
	exp := AccessStruct{lin, col, str, id}
	return exp
}

func (p AccessStruct) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var result, tempStruct environment.Symbol
	tempStruct = env.(environment.Environment).GetVariable(p.Struct)
	//tempStruct = p.Struct.Ejecutar(ast, env) //sym{[1,2,3,4]}

	if tempStruct.Tipo == environment.STRUCT {

		if variable, ok := tempStruct.Valor.(map[string]environment.Symbol)[p.Id]; ok {
			return variable
		}
		ast.SetError("No existe el elemento " + p.Id)
		return result
	}
	ast.SetError("La expresión no es un struct")
	return result
}