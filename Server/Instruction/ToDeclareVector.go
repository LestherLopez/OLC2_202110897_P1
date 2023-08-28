package instructions

import (
	environment "Server/Environment"
	interfaces "Server/Interfaces"
	"fmt"
)

type ToDecalreVector struct {
	Lin   int
	Col   int
	id_var string
	type_var environment.TipoExpresion
	valores []interface{}
}
//lin int, col int, id_var string, type_var environment.TipoExpresion, valor interfaces.Expression, constant bool
func NewToDecalreVector(lin int, col int, id_var string, type_var environment.TipoExpresion, valores []interface{}) ToDecalreVector {
	return ToDecalreVector{lin, col, id_var, type_var, valores}
}

func (p ToDecalreVector) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	var isType bool  
	isType = true
	//se definen las no constantes
	fmt.Println(p.valores[0])
		
		if(p.type_var != environment.NULL && p.valores != nil){
				for _, inst := range p.valores {
					valor := inst.(interfaces.Expression).Ejecutar(ast, env)
					
					if p.type_var==valor.Tipo {
						isType = true
					}else{
						
						ast.SetError("ERROR: el valor "+fmt.Sprintf("%v", valor.Valor)+" no corresponde al tipo ingresado")
						isType = false
					}
					
				}
		}else if(p.valores==nil){
				isType = true
		}
		if(isType){
			value:=environment.SymbolVector{Lin: p.Lin, Col: p.Col, Id: p.id_var, Tipo: p.type_var, Valor: p.valores, Transfer: environment.NULL}
			env.(environment.Environment).KeepVector(p.id_var,  value)
		}else if(isType && p.valores==nil){
			value:=environment.SymbolVector{Lin: p.Lin, Col: p.Col, Id: p.id_var, Tipo: p.type_var, Valor: nil, Transfer: environment.NULL}
			env.(environment.Environment).KeepVector(p.id_var,  value)
		}

	return nil

}