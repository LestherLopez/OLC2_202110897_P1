package instructions

import (
	environment "Server/Environment"
	interfaces "Server/Interfaces"
	"fmt"
)

type While struct {
	Lin   int
	Col   int
	exp_conditional interfaces.Expression
	sentences []interface{}
}

func NewWhile(lin int, col int, exp_conditional interfaces.Expression, sentences []interface{}) While {
	return While{lin, col, exp_conditional, sentences}
}

func (p While) Ejecutar(ast *environment.AST, env interface{}) interface{} {

	var conditional environment.Symbol
	conditional = p.exp_conditional.(interfaces.Expression).Ejecutar(ast, env)
	whileEnv := environment.NewEnvironment(env, "While environment")
	//while in go
	if(conditional.Tipo != environment.BOOLEAN){
		fmt.Println("El tipo de variable es incorrecto para un while, tiene que ser boolean")
		ast.SetError("ERROR: El tipo de variable es incorrecto para un while, tiene que ser boolean")
		return nil
	}
	for conditional.Valor.(bool) {
		for _, inst := range p.sentences {
			inst.(interfaces.Instruction).Ejecutar(ast, whileEnv)
		}
	
		
		conditional =p.exp_conditional.(interfaces.Expression).Ejecutar(ast, whileEnv)
	}
	print("while")
	//ast.SetPrint(consoleOut + "\n")
	return nil
}