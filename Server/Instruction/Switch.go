package instructions

import (
	environment "Server/Environment"
	interfaces "Server/Interfaces"
	"fmt"
)

type Switch struct {
	Lin   int
	Col   int
	exp_conditional interfaces.Expression
	exp 			interfaces.Expression
	sentence []interface{}
	sentence_deafult []interface{}
}

func NewSwitch(lin int, col int, expc interfaces.Expression, exp interfaces.Expression,senten []interface{}, senten_deafult []interface{}) Switch {
	return Switch{Lin: lin, Col: col, exp_conditional: expc, exp: exp,sentence: senten, sentence_deafult: senten_deafult}
}

func (p Switch) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	fmt.Println("entre a Switch")
	var conditional, exp environment.Symbol
	conditional = p.exp_conditional.(interfaces.Expression).Ejecutar(ast, env)
	exp = p.exp.(interfaces.Expression).Ejecutar(ast, env)

	if(conditional.Tipo != exp.Tipo){
		fmt.Println("El tipo de variable es incorrecto para comparar")
		return nil
	}
	if(conditional.Valor==exp.Valor){
		//hacer nuevo environment
		var SwitchEnv environment.Environment
		SwitchEnv = environment.NewEnvironment(env, "Switch case environment")
		fmt.Println("a")
			//ejecutar sentencias
		for _, inst := range p.sentence {
				inst.(interfaces.Instruction).Ejecutar(ast, SwitchEnv)
		}

		//return nil
		
		
			//condicional por si viene return
		
	}else{
		var deafultEnv environment.Environment
		deafultEnv = environment.NewEnvironment(env, "deafult environment")
		for _, inst := range p.sentence_deafult {
			inst.(interfaces.Instruction).Ejecutar(ast, deafultEnv)
			
		}

	}
	return nil
}