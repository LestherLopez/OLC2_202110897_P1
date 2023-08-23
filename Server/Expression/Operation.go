package expressions

import (
	environment "Server/Environment"
	interfaces "Server/Interfaces"
	"fmt"
	"strconv"
)

type Operation struct {
	Lin      int
	Col      int
	Op_izq   interfaces.Expression
	Operador string
	Op_der   interfaces.Expression
}

func NewOperation(lin int, col int, Op1 interfaces.Expression, Operador string, Op2 interfaces.Expression) Operation {
	exp := Operation{Lin: lin, Col: col, Op_izq: Op1, Operador: Operador, Op_der: Op2}
	return exp
}

func (o Operation) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var dominante environment.TipoExpresion

	tabla_dominante := [5][5]environment.TipoExpresion{
		//		INTEGER			FLOAT				STRING				BOOLEAN				NULL
		{environment.INTEGER, environment.FLOAT, environment.STRING, environment.BOOLEAN, environment.NULL},
		//FLOAT
		{environment.FLOAT, environment.FLOAT, environment.STRING, environment.NULL, environment.NULL},
		//STRING
		{environment.STRING, environment.STRING, environment.STRING, environment.STRING, environment.NULL},
		//BOOLEAN
		{environment.BOOLEAN, environment.NULL, environment.STRING, environment.BOOLEAN, environment.NULL},
		//NULL
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
	}

	if(o.Op_der==nil){
		var op1 environment.Symbol
		op1 = o.Op_izq.Ejecutar(ast, env)
		if(o.Operador=="-"){
			if op1.Tipo==0{ //int
				
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: op1.Tipo, Valor: -op1.Valor.(int) }
			}else if op1.Tipo==1{ //float
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: op1.Tipo, Valor: -val1 }
			}
		}
	}

	var op1, op2 environment.Symbol
	op1 = o.Op_izq.Ejecutar(ast, env)
	op2 = o.Op_der.Ejecutar(ast, env)
	
	switch o.Operador {
	case "+":
		{
			//validar tipo dominante
			dominante = tabla_dominante[op1.Tipo][op2.Tipo]
			//valida el tipo
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: op1.Valor.(int) + op2.Valor.(int)}
			} else if dominante == environment.FLOAT {
				
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: val1 + val2}
			} else if dominante == environment.STRING {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: r1 + r2}
			} else {
				ast.SetError("ERROR: No es posible sumar")
			}
		}
	case "-":
		{
			dominante = tabla_dominante[op1.Tipo][op2.Tipo]

			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: op1.Valor.(int) - op2.Valor.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: val1 - val2}
			} else {
				ast.SetError("ERROR: No es posible restar")
			}
		}
	case "*":
		{
			dominante = tabla_dominante[op1.Tipo][op2.Tipo]
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: op1.Valor.(int) * op2.Valor.(int)}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: val1 * val2}
			} else {
				ast.SetError("ERROR: No es posible Multiplicar")
			}
		}
	case "/":
		{
			dominante = tabla_dominante[op1.Tipo][op2.Tipo]
			if dominante == environment.INTEGER {
				if op2.Valor.(int) != 0 {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: op1.Valor.(int) / op2.Valor.(int)}
				} else {
					ast.SetError("ERROR: No es posible dividir en cero")
				}

			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				if val2 != 0 {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: val1 / val2}
				} else {
					ast.SetError("ERROR: No es posible dividir en cero")
				}
			} else {
				ast.SetError("ERROR: No es posible Dividir")
			}

		}
	case "%":
	 	{
			if(op1.Tipo==0 && op2.Tipo==0){
				if(op2.Valor.(int)!=0){
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: dominante, Valor: op1.Valor.(int) % op2.Valor.(int)}
				}else if(op2.Valor.(int)==0){
					ast.SetError("ERROR: MATH ERROR")
				}
				
			} else{
				ast.SetError("ERROR: No es posible obtener el modulo")
			}
		}
	case "+=":
		{
			value := env.(environment.Environment).GetVariable(op1.Id)
			
			if op2.Tipo==value.Tipo{
				//int
				if(op2.Tipo==0){
					value.Valor = value.Valor.(int) + op2.Valor.(int)
					env.(environment.Environment).SetVariable(value.Id, value)
					//return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.INTEGER, Valor: value.Valor}
				//float
				}else if(op2.Tipo==1){
					val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", value.Valor), 64)
					val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
					value.Valor = val1+ val2
					env.(environment.Environment).SetVariable(value.Id, value)
					//return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.FLOAT, Valor: value.Valor}
					//string	
				}else if(op2.Tipo==2){
					r1 := fmt.Sprintf("%v", value.Valor)
					r2 := fmt.Sprintf("%v", op2.Valor)
					value.Valor = r1 + r2
					env.(environment.Environment).SetVariable(value.Id, value)
					//return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.FLOAT, Valor: value.Valor}
				}
			}else{
				ast.SetError("ERROR: No son compatibles los tipos")
			}
		}
	case "-=":
		{
			value := env.(environment.Environment).GetVariable(op1.Id)
			
			if op2.Tipo==value.Tipo{
				//int
				if(op2.Tipo==0){
					value.Valor = value.Valor.(int) - op2.Valor.(int)
					env.(environment.Environment).SetVariable(value.Id, value)
					//return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.INTEGER, Valor: value.Valor}
				//float
				}else if(op2.Tipo==1){
					val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", value.Valor), 64)
					val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
					value.Valor = val1+ val2
					env.(environment.Environment).SetVariable(value.Id, value)
					//return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.FLOAT, Valor: value.Valor}
					//string	
				}
			}else{
				ast.SetError("ERROR: No son compatibles los tipos")
			}
		}
		//relational operations with int, float/ string y character
	case "<":
		{
			dominante = tabla_dominante[op1.Tipo][op2.Tipo]
			if (op1.Tipo==0 && op2.Tipo == 0) {   //int
				
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: op1.Valor.(int) < op2.Valor.(int)}
			} else if (op1.Tipo==1 && op2.Tipo == 1) {  //float
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: val1 < val2}
			} else if(op1.Tipo==2 && op2.Tipo == 2) { //string
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: r1 < r2}
			} else if(op1.Tipo==5 && op2.Tipo == 5){ //char
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: r1 < r2}
			}else{
				ast.SetError("ERROR: No es posible comparar <")
			}
		}
	case ">":
		{
			dominante = tabla_dominante[op1.Tipo][op2.Tipo]
			if (op1.Tipo==0 && op2.Tipo == 0) {   //int
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: op1.Valor.(int) > op2.Valor.(int)}
			} else if (op1.Tipo==1 && op2.Tipo == 1) {  //float
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: val1 > val2}
			} else if(op1.Tipo==2 && op2.Tipo == 2) { //string
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: r1 > r2}
			} else if(op1.Tipo==5 && op2.Tipo == 5){ //char
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: r1 > r2}
			}else{
				ast.SetError("ERROR: No es posible comparar <")
			}
		}
	case "<=":
		{
			dominante = tabla_dominante[op1.Tipo][op2.Tipo]
			if (op1.Tipo==0 && op2.Tipo == 0) {   //int
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: op1.Valor.(int) <= op2.Valor.(int)}
			} else if (op1.Tipo==1 && op2.Tipo == 1) {  //string
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: val1 <= val2}
			} else if(op1.Tipo==2 && op2.Tipo == 2) { //char
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: r1 <= r2}
			} else if(op1.Tipo==5 && op2.Tipo == 5){ //char
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: r1 <= r2}
			}else{
				ast.SetError("ERROR: No es posible comparar <")
			}
		}
	case ">=":
		{
			dominante = tabla_dominante[op1.Tipo][op2.Tipo]
			if (op1.Tipo==0 && op2.Tipo == 0) {   //int
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: op1.Valor.(int) >= op2.Valor.(int)}
			} else if (op1.Tipo==1 && op2.Tipo == 1) {  //string
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: val1 >= val2}
			} else if(op1.Tipo==2 && op2.Tipo == 2) { //char
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: r1 >= r2}
			} else if(op1.Tipo==5 && op2.Tipo == 5){ //char
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: r1 >= r2}
			}else{
				ast.SetError("ERROR: No es posible comparar <")
			}
		}
	case "==":
		{
			if op1.Tipo == op2.Tipo {
				
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: op1.Valor == op2.Valor}
			} else {
				ast.SetError("ERROR: No es posible comparar == ")
			}
		}
	case "!=":
		{
			if op1.Tipo == op2.Tipo {
			
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: op1.Valor != op2.Valor}
			} else {
				ast.SetError("ERROR: No es posible comparar !=")
			}
		}
	case "&&":
		{
			if (op1.Tipo == environment.BOOLEAN) && (op2.Tipo == environment.BOOLEAN) {
				
				
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: op1.Valor.(bool) && op2.Valor.(bool)}
			} else {
				ast.SetError("ERROR: tipo no compatible &&")
			}
		}
	case "||":
		{
			if (op1.Tipo == environment.BOOLEAN) && (op2.Tipo == environment.BOOLEAN) {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: op1.Valor.(bool) || op2.Valor.(bool)}
			} else {
				ast.SetError("ERROR: tipo no compatible ||")
			}
		}
	case "!":
		{
			if(op1.Tipo == environment.BOOLEAN){
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.BOOLEAN, Valor: !op1.Valor.(bool) }
			}
		}
	}
	
	var result interface{}
	return environment.Symbol{Lin: o.Lin, Col: o.Col, Tipo: environment.NULL, Valor: result}
}