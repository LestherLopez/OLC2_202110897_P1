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
	sentence interfaces.Instruction
	sentence_deafult []interface{}
}

func NewSwitch(lin int, col int, expc interfaces.Expression, exp interfaces.Expression,senten interfaces.Instruction, senten_deafult []interface{}) Switch {
	return Switch{Lin: lin, Col: col, exp_conditional: expc, exp: exp,sentence: senten, sentence_deafult: senten_deafult}
}

func (p Switch) Ejecutar(ast *environment.AST, env interface{}) interface{} {

		var conditional environment.Symbol
		conditional  = p.exp_conditional.(interfaces.Expression).Ejecutar(ast, env)
		
		env.(environment.Environment).KeepVariableSwitch("switch", conditional)
		conditional1 := env.(environment.Environment).GetVariableSwitch("switch")
		fmt.Println(conditional1.Valor)
		fmt.Println("entro al guardado de expresion")
		if(p.sentence!=nil){
			var SwitchEnv environment.Environment
			SwitchEnv = environment.NewEnvironment(env, "Switch case environment")
			p.sentence.(interfaces.Instruction).Ejecutar(ast, SwitchEnv)
		}else{
			var deafultEnv environment.Environment
			deafultEnv = environment.NewEnvironment(env, "Switch case environment")
			for _, inst := range p.sentence_deafult {
				inst.(interfaces.Instruction).Ejecutar(ast, deafultEnv)
			}
		}
		

	return nil
}