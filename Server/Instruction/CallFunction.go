package instructions

import (
	environment "Server/Environment"
	interfaces "Server/Interfaces"
	"fmt"
)

type CallFunction struct {
	Lin   int
	Col   int
	id 	  string
}

func NewCallFunction(lin int, col int, id string, ) CallFunction {
	exp := CallFunction{lin, col, id}
	return exp
}

func (p CallFunction) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	funcion := env.(environment.Environment).GetFunction(p.id)
	FunctionEnv := environment.NewEnvironment(env, "Function "+p.id+" environment")
	fmt.Print("VAMOS A IMPRIMIR ENVIRONMENT ID")
	fmt.Print(FunctionEnv.NameEnv)
	for _, inst := range funcion.Bloque_sentencias {
		inst.(interfaces.Instruction).Ejecutar(ast, FunctionEnv)

	}
	return nil
}