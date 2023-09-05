package instructions

import (
	environment "Server/Environment"
	interfaces "Server/Interfaces"
	"fmt"
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

func (p Assignation) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	//variable := env.(environment.Environment).GetVariable(p.id_var)
	/*if variable == un {
		errMsg := fmt.Sprintf("ERROR: No se ha definido la variable previamente %s", p.id_var)
		panic(errMsg)
	}*/
	var valor_asignar environment.Symbol
	
	valor_asignar = p.valor.(interfaces.Expression).Ejecutar(ast, env)
	variable := env.(environment.Environment).GetVariable(p.id_var)
	fmt.Print(variable)
	//no se puede cambiar
	if(!variable.Mutable){
		fmt.Printf("No es posible cambiar el valor de la variable %v\n", p.id_var)
	}else if(variable.Mutable){
		if(valor_asignar.Tipo == variable.Tipo){
			valor_asignar.Mutable = variable.Mutable
			valor_asignar.Transfer = variable.Transfer

			env.(environment.Environment).SetVariable(p.id_var, valor_asignar)
			
			return nil
		}else{
			fmt.Printf("No es posible cambiar el valor de la variable %v\n debido a que los timpos no son equivalentes", p.id_var)
		}
		
	}
	//ast.SetPrint(consoleOut + "\n")
	return nil
}