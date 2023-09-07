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
	parameter []interface{}
	numero int
}

func NewCallFunction(lin int, col int, id string, parameter []interface{}, numero int) CallFunction {
	exp := CallFunction{lin, col, id, parameter, numero}
	return exp
}

func (p CallFunction) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	var updatevectorid, updatevectoridfunc string
	var updatevector bool
	fmt.Print("entramos a call")
	funcion := env.(environment.Environment).GetFunction(p.id)
	FunctionEnv := environment.NewEnvironment(env, "Function "+p.id+" environment")
	switch p.numero{
	case 1: //funcion con parametros normales
		if(len(p.parameter) == len(funcion.Bloque_parametros)){
			for i := 0; i < len(funcion.Bloque_parametros); i++ {
				parametro:=funcion.Bloque_parametros[i].(interfaces.Expression).Ejecutar(ast, env)
				valor:=p.parameter[i].(interfaces.Expression).Ejecutar(ast, env)
				updatevectorid=valor.Id
				fmt.Print(parametro)
				fmt.Print(valor)
				if(valor.Tipo==parametro.Tipo){
					if(parametro.IsList){
						
						vector:=env.(environment.Environment).GetVector(valor.Id)	
						FunctionEnv.KeepVector(parametro.Id, vector)
						updatevectoridfunc= parametro.Id
						updatevector = true
					}else{
					
				
						
						fmt.Print(parametro.Id)
					FunctionEnv.KeepVariable(parametro.Id, environment.Symbol{Valor: valor.Valor, Lin: parametro.Lin, Col: parametro.Col, Id: parametro.Id, Tipo: parametro.Tipo, Mutable: true})
					}
				
				
				}else{
					ast.SetError("ERROR: El tipo del parametro ingresado es incorrecto")
					fmt.Print("ERROR: El tipo del parametro ingresado es incorrecto")
				}
			
			}

			//ejecutar sentencias
			for _, inst := range funcion.Bloque_sentencias {
				inst.(interfaces.Instruction).Ejecutar(ast, FunctionEnv)

			}

			//actualizar inout
			if(updatevector){
				vector := env.(environment.Environment).GetVector(updatevectorid)
				element := FunctionEnv.GetVector(updatevectoridfunc)	
				for _, inst := range vector.Valor {
					fmt.Print(inst)
	
				}
	
				vector.Valor=element.Valor
				env.(environment.Environment).SetVector(updatevectorid, vector)	
			}

		}
	case 2:
	
		fmt.Print(FunctionEnv.NameEnv)
		for _, inst := range funcion.Bloque_sentencias {
			inst.(interfaces.Instruction).Ejecutar(ast, FunctionEnv)

		}
	}
	return nil
}