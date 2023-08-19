package instructions

import (
	environment "Server/Environment"
	interfaces "Server/Interfaces"
	"fmt"
)

type If struct {
	Lin   int
	Col   int
	exp_conditional interfaces.Expression
	sentence interfaces.Instruction
	sentence_else interfaces.Instruction
}

func NewIf(lin int, col int, expc interfaces.Expression, senten interfaces.Instruction, senten_else interfaces.Instruction) If {
	return If{Lin: lin, Col: col, exp_conditional: expc, sentence: senten, sentence_else: senten_else}
}

func (p If) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	
	conditional := p.exp_conditional.(interfaces.Expression).Ejecutar(ast, env)
	if(conditional.Tipo != environment.BOOLEAN){
		fmt.Println("El tipo de variable es incorrecto para un If")
		return nil
	}
	if(conditional.Valor.(bool)){
		//hacer nuevo environment
	
		ifEnv := environment.NewEnvironment(env, "If environment")
			//ejecutar sentencias
		p.sentence.Ejecutar(ast, ifEnv)
		
		
			//condicional por si viene return
		
	}else{
		
		elseEnv := environment.NewEnvironment(env, "Else")
		p.sentence_else.Ejecutar(ast, elseEnv)
	}
	
	return nil
}