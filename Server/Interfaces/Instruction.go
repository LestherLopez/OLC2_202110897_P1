package interfaces

import environment "Server/Environment"

type Instruction interface {
	Ejecutar(ast *environment.AST, env interface{}) interface{}
}