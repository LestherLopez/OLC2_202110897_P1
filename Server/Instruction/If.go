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
/*	if(conditional.Tipo != environment.BOOLEAN){
		return nil
	}
	if(conditional.Valor.(bool) ){
		//hacer nuevo environment
		var ifEnv environment.Environment
		ifEnv = environment.NewEnvironment(env.(environment.Environment), "IF")
		element := p.sentence.Ejecutar(ast, ifEnv)
		if(element!=nil){

			//condicional por si viene return
		}
	}else{
		var elseEnv environment.Environment
		elseEnv = environment.NewEnvironment(env.(environment.Environment), "Else")
		p.sentence_else.Ejecutar(ast, elseEnv)
	}*/
	
	//valueToPrint := i.Value.(interfaces.Expression).Ejecutar(ast, env)
	consoleOut := fmt.Sprintf("%v", conditional.Valor)
	ast.SetPrint(consoleOut + "\n")
	return nil
}