package instructions

import (
	environment "Server/Environment"
	expressions "Server/Expression"
	interfaces "Server/Interfaces"
	"fmt"
)

type For struct {
	Lin   int
	Col   int
	id_var string 
	first_exp interfaces.Expression
	second_exp interfaces.Expression
	sentences []interface{}
}

func NewFor(lin int, col int, id_var string, first_exp interfaces.Expression,  second_exp interfaces.Expression, sentences []interface{}) For {
	return For{lin, col, id_var, first_exp, second_exp, sentences}
}

func (p For) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	/*for i in 1...5 {
		print(i)
		}*/
	
	//expresion
	if(p.second_exp==nil){

	}else if(p.second_exp!=nil){
		var expresion1, expresion2 environment.Symbol
		expresion1 = p.first_exp.Ejecutar(ast, env)
		expresion2 = p.second_exp.Ejecutar(ast, env)
		fmt.Println("entra 1")
		if(expresion1.Tipo==environment.INTEGER && expresion2.Tipo==environment.INTEGER){
			fmt.Println("entra 2")
			if(expresion2.Valor.(int)>=expresion1.Valor.(int)){
				forEnv := environment.NewEnvironment(env, "For environment")
				id := NewTodeclare(p.Lin, p.Col, p.id_var, environment.INTEGER, p.first_exp, false)
				id.Ejecutar(ast, forEnv)
				//guardar variable que inicia el rango
				

				valint1 := expresion1.Valor.(int)
				valint2 := expresion2.Valor.(int)
				for valint1 != (valint2+1) {
					for _, inst := range p.sentences {
						fmt.Println("bloque")
						inst.(interfaces.Instruction).Ejecutar(ast, forEnv)
					}
				
					
									
					
					valint1 = valint1 + 1
					exp := expressions.NewPrimitive(p.Lin,p.Col,1,environment.INTEGER)
					exp.Ejecutar(ast, forEnv)
					aumento := NewIncreaseDecrease(p.Lin, p.Col, p.id_var, "+=", exp)
					aumento.Ejecutar(ast, forEnv)
				}
				fmt.Println("entra 3")
			}else{
				ast.SetError("ERROR: El segundo valor tiene que ser mayor o igual que el primero")
			}
		}else{
			ast.SetError("ERROR: Los tipos del rango deben de ser Int los dos")
		}
	}

	
	print("For")
	//ast.SetPrint(consoleOut + "\n")
	return nil
}