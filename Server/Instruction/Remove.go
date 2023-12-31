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
	valor = p.exp.Ejecutar(ast, env)
	vector := env.(environment.Environment).GetVector(p.id_vector)
	if(environment.INTEGER==valor.Tipo){
		//value:=environment.Symbol{Lin: valor.Lin, Col: valor.Col, Id: valor.Id, Tipo: valor.Tipo, Valor: valor.Valor, Transfer: environment.NULL}
		if(valor.Valor.(int)>=0 && valor.Valor.(int)<len(vector.Valor)){

		
			vector.Valor = append(vector.Valor[:valor.Valor.(int)], vector.Valor[valor.Valor.(int)+1:]... )
			env.(environment.Environment).SetVector(p.id_vector, vector)
			 env.(environment.Environment).GetVector(p.id_vector)
			//fmt.Print("vector nuevo")
			
		}else{
			ast.SetError("ERROR: El indice esta fuera del rango del vector")
			fmt.Print("ERROR: El indice esta fuera del rango del vector")
		}
	}else{
		ast.SetError("ERROR: El indice del elemento al que quiere acceder tiene que ser de tipo entero")
		fmt.Print("ERROR: El indice del elemento al que quiere acceder tiene que ser de tipo entero")
	}
	return nil
}