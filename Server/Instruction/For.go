package instructions

import (
	environment "Server/Environment"
	interfaces "Server/Interfaces"
	"fmt"
)

type For struct {
	Lin   int
	Col   int
	exp_conditional interfaces.Expression
	sentences []interface{}
}

func NewFor(lin int, col int, exp_conditional interfaces.Expression, sentences []interface{}) For {
	return For{lin, col, exp_conditional, sentences}
}

func (p For) Ejecutar(ast *environment.AST, env interface{}) interface{} {

	var conditional environment.Symbol
	conditional = p.exp_conditional.(interfaces.Expression).Ejecutar(ast, env)
	ForEnv := environment.NewEnvironment(env, "For environment")
	//For in go
	if(conditional.Tipo != environment.BOOLEAN){
		fmt.Println("El tipo de variable es incorrecto para un For, tiene que ser boolean")
		ast.SetError("ERROR: El tipo de variable es incorrecto para un For, tiene que ser boolean")
		return nil
	}
	for conditional.Valor.(bool) {
		for _, inst := range p.sentences {
			inst.(interfaces.Instruction).Ejecutar(ast, ForEnv)
		}
	
		
		conditional =p.exp_conditional.(interfaces.Expression).Ejecutar(ast, ForEnv)
	}
	print("For")
	//ast.SetPrint(consoleOut + "\n")
	return nil
}