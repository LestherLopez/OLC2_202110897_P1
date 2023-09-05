package instructions

import (
	environment "Server/Environment"
	interfaces "Server/Interfaces"
	"fmt"
)

type Print struct {
	Lin   int
	Col   int
	Value []interface{}
}

func NewPrint(lin int, col int, val []interface{}) Print {
	return Print{lin, col, val}
}

func (p Print) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	var consoleOut string
	for _, inst := range p.Value {
		valueToPrint := inst.(interfaces.Expression).Ejecutar(ast, env)
		consoleOut = consoleOut+ fmt.Sprintf("%v", valueToPrint.Valor)
		consoleOut = consoleOut + " "
		
	}
	ast.SetPrint(consoleOut + "\n")
	return nil
}