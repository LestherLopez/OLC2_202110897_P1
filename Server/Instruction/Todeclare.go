package instructions

import (
	environment "Server/Environment"
	primitive "Server/Expression"
	interfaces "Server/Interfaces"
	"fmt"
)

type Todeclare struct {
	Lin   int
	Col   int
	id_var string
	type_var environment.TipoExpresion
	valor interfaces.Expression
	constant bool
}
//lin int, col int, id_var string, type_var environment.TipoExpresion, valor interfaces.Expression, constant bool
func NewTodeclare(lin int, col int, id_var string, type_var environment.TipoExpresion, valor interfaces.Expression, constant bool) Todeclare {
	return Todeclare{lin, col, id_var, type_var, valor, constant}
}

func (p Todeclare) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	/*fmt.Println("hola")
	var valor_asignar environment.Symbol
	if(p.valor!=nil){
		if(valor_asignar.Tipo==p.type_var){
		valor_asignar = p.valor.Ejecutar(ast, env)
		env.(environment.Environment).KeepVariable(p.id_var, valor_asignar)
		}
	}*/
	
	//variable con tipo y valor
	if(p.type_var != environment.NULL && p.valor != nil){
			valueTodeclare := p.valor.(interfaces.Expression).Ejecutar(ast, env)
			if(valueTodeclare.Tipo==p.type_var){
					env.(environment.Environment).KeepVariable(p.id_var, valueTodeclare)
					fmt.Println("Los tipos coinciden")
			} else{
				fmt.Println("Los tipos no coinciden")
			}
	//variable con valor pero sin tipo
	}else if(p.type_var==environment.NULL && p.valor != nil){
			valueTodeclare := p.valor.(interfaces.Expression).Ejecutar(ast, env)
			env.(environment.Environment).KeepVariable(p.id_var, valueTodeclare)
			fmt.Println("var sin tipo pero con valor")
	//variable con tipo pero sin valor
	}else if(p.type_var!=environment.NULL && p.valor == nil){
		fmt.Println("var con tipo pero sin valor")
		val := primitive.NewPrimitive(p.Lin, p.Col, nil, environment.NULL) 
		valueTodeclare := val.Ejecutar(ast, env)
		env.(environment.Environment).KeepVariable(p.id_var, valueTodeclare)
	}
	
	return nil
	

	//ast.SetPrint(consoleOut + "\n")
	//return valor_asignar
}