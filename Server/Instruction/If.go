package instructions

import (
	environment "Server/Environment"
	interfaces "Server/Interfaces"
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

func (i If) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	conditional := i.exp_conditional.Ejecutar(ast, env)
	if(conditional.Tipo != environment.BOOLEAN){
		return nil
	}
	if(conditional.Valor.(bool) ){
		element := i.sentence.Ejecutar(ast, env)
		if(element!=nil){
			//condicional por si viene return
		}
	}else{
		i.sentence_else.Ejecutar(ast, env)
	}
	
	//valueToPrint := i.Value.(interfaces.Expression).Ejecutar(ast, env)
	//consoleOut := fmt.Sprintf("%v", valueToPrint.Valor)
	//ast.SetPrint(consoleOut + "\n")
	return nil
}