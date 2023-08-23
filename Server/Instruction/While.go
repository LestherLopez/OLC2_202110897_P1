package instructions

import (
	environment "Server/Environment"
	interfaces "Server/Interfaces"
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