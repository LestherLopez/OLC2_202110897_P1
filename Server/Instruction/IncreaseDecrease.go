package instructions

import (
	environment "Server/Environment"
	interfaces "Server/Interfaces"
	"fmt"
	"strconv"
)

type IncreaseDecrease struct {
	Lin   int
	Col   int
	id_var 	  string
	operator string
	valor 	interfaces.Expression
}

func NewIncreaseDecrease(lin int, col int, id_var string, operator string,valor interfaces.Expression) IncreaseDecrease {
	incdec := IncreaseDecrease{lin, col, id_var, operator,valor}
	return incdec
}

func (p IncreaseDecrease) Ejecutar(ast *environment.AST, env interface{}) interface{} {
	//result := env.(environment.Environment).GetVariable(p.id)
	//return result
	var valor environment.Symbol
	valor = p.valor.Ejecutar(ast,env)
	switch p.operator {
		case "+=":
			{
				value := env.(environment.Environment).GetVariable(p.id_var)
				
				if valor.Tipo==value.Tipo{
					//int
					if(valor.Tipo==0){
						value.Valor = value.Valor.(int) + valor.Valor.(int)
						env.(environment.Environment).SetVariable(p.id_var, value)
						//return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.INTEGER, Valor: value.Valor}
					//float
					}else if(valor.Tipo==1){
						val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", value.Valor), 64)
						val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valor.Valor), 64)
						value.Valor = val1+ val2
						env.(environment.Environment).SetVariable(p.id_var, value)
						//return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.FLOAT, Valor: value.Valor}
						//string	
					}else if(valor.Tipo==2){
						r1 := fmt.Sprintf("%v", value.Valor)
						r2 := fmt.Sprintf("%v", valor.Valor)
						value.Valor = r1 + r2
						env.(environment.Environment).SetVariable(p.id_var, value)
						//return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.FLOAT, Valor: value.Valor}
					}
				}else{
					ast.SetError("ERROR: No son compatibles los tipos")
				}
			}
		case "-=":
			{
				value := env.(environment.Environment).GetVariable(p.id_var)
				
				if valor.Tipo==value.Tipo{
					//int
					if(valor.Tipo==0){
						value.Valor = value.Valor.(int) - valor.Valor.(int)
						env.(environment.Environment).SetVariable(p.id_var, value)
						//return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.INTEGER, Valor: value.Valor}
					//float
					}else if(valor.Tipo==1){
						val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", value.Valor), 64)
						val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", valor.Valor), 64)
						value.Valor = val1+ val2
						env.(environment.Environment).SetVariable(p.id_var, value)
						//return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.FLOAT, Valor: value.Valor}
						//string	
					}
				}else{
					ast.SetError("ERROR: No son compatibles los tipos")
				}
			}
	
	}
	return nil	
}