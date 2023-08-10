package interfaces

import environment "Server/Environment"

type Expression interface {
	Ejecutar(ast *environment.AST, env interface{}) environment.Symbol
}