package instructions

import (
	environment "Server/Environment"
	interfaces "Server/Interfaces"
)

type Assignation struct {
	Lin   int
	Col   int
	id_var string
	valor interfaces.Expression
}

func NewAssignation(lin int, col int, id_var string, valor interfaces.Expression) Assignation {
	return Assignation{lin, col, id_var, valor}
}

func (p Assignation) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	//variable := env.(environment.Environment).GetVariable(p.id_var)
	/*if variable == un {
		errMsg := fmt.Sprintf("ERROR: No se ha definido la variable previamente %s", p.id_var)
		panic(errMsg)
	}*/
	var valor_asignar environment.Symbol
	
	valor_asignar = p.valor.Ejecutar(ast, env)
	env.(environment.Environment).SetVariable(p.id_var, valor_asignar)

	//ast.SetPrint(consoleOut + "\n")
	return valor_asignar
}