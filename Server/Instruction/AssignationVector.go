package instructions

import (
	environment "Server/Environment"
	interfaces "Server/Interfaces"
	"fmt"
)

type AssignationVector struct {
	Lin   int
	Col   int
	first_id string
	first_exp interfaces.Expression
	second_exp interfaces.Expression
}

func NewAssignationVector(lin int, col int, first_id string, first_exp interfaces.Expression, second_exp interfaces.Expression) AssignationVector {
	return AssignationVector{lin, col, first_id, first_exp,  second_exp}
}

func (p AssignationVector) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	//variable := env.(environment.Environment).GetVariable(p.id_var)
	/*if variable == un {
		errMsg := fmt.Sprintf("ERROR: No se ha definido la variable previamente %s", p.id_var)
		panic(errMsg)
	}*/

		//ejecutar expresiones
	
	var indice, valor_asignar environment.Symbol
	indice = p.first_exp.(interfaces.Expression).Ejecutar(ast, env)
	valor_asignar = p.second_exp.(interfaces.Expression).Ejecutar(ast, env)
	vector := env.(environment.Environment).GetVector(p.first_id)
	//no se puede cambiar
	if(indice.Tipo==environment.INTEGER){
		
		if(indice.Valor.(int)>=0 && indice.Valor.(int)<len(vector.Valor)){
			if(valor_asignar.Tipo==vector.Tipo){
				
				index:=indice.Valor.(int)
				
				
				vector.Valor[index] = p.second_exp
			
				env.(environment.Environment).SetVector(vector.Id, vector)
				
			}else{
				ast.SetError("ERROR: El tipo del vector no corresponde al tipo de valor a asignar")
				fmt.Print("ERROR: El tipo del vector no corresponde al tipo de valor a asignar")
			}
		}else{
			ast.SetError("ERROR: El indice esta fuera del rango del vector")
			fmt.Print("ERROR: El indice esta fuera del rango del vector")
		}
	}else{
		ast.SetError("ERROR: El indice esta fuera del rango del vector")
		fmt.Printf("No es posible cambiar el valor de la variable %v\n", p.first_id)
	}
	//ast.SetPrint(consoleOut + "\n")
	return nil
}