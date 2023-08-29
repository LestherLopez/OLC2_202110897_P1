package instructions

import (
	environment "Server/Environment"
	interfaces "Server/Interfaces"
	"fmt"
)

type Remove struct {
	Lin   int
	Col   int
	id_vector string
	exp    interfaces.Expression
}
//lin int, col int, id_vector string, type_var environment.TipoExpresion, valor interfaces.Expression, constant bool
func NewRemove(lin int, col int, id_vector string, exp interfaces.Expression) Remove {
	return Remove{lin, col, id_vector, exp}
}

func (p Remove) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	var valor environment.Symbol
	
	vector := env.(environment.Environment).GetVector(p.id_vector)
	for _, inst := range vector.Valor {
		fmt.Print(inst)
	}
	if(vector.Tipo==valor.Tipo){
		//value:=environment.Symbol{Lin: valor.Lin, Col: valor.Col, Id: valor.Id, Tipo: valor.Tipo, Valor: valor.Valor, Transfer: environment.NULL}
		newVector := vector.Valor[:len(vector.Valor)-1]
		vector.Valor = newVector
		env.(environment.Environment).SetVector(p.id_vector, vector)
		vectorn := env.(environment.Environment).GetVector(p.id_vector)
		fmt.Print("vector nuevo")
		for _, inst := range vectorn.Valor {
			fmt.Print(inst)
		}
	}else{
		ast.SetError("ERROR: El tipo a ingresar en el vector no corresponde a su tipo")
	}
	return nil
}