package environment

import (
	"fmt"
	"strings"
)

type Environment struct {
	
	Anterior  interface{}
	Variables map[string]Symbol
	SwitchVar Symbol
	NameEnv   string
	Vectors   map[string]SymbolVector
	Functions map[string]SymbolFunction

}

func NewEnvironment(ant interface{}, ide string) Environment {
	return Environment{
		Anterior:  ant,
		Variables: make(map[string]Symbol),
		Vectors: make(map[string]SymbolVector),
		Functions: make(map[string]SymbolFunction),
		NameEnv:   ide,
	}
}

func (env Environment) KeepVariable(id string, value Symbol) {

	if variable, ok := env.Variables[id]; ok {
		fmt.Println("La variable "+id+" ya existe", variable)
		return
	}
	env.Variables[id] = value
	
}



func (env Environment) GetVariable(id string) Symbol {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if variable, ok := tmpEnv.Variables[id]; ok {
			return variable
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("La variable ", id, " no existe en este environment")
	return Symbol{Lin: 0, Col: 0, Id: "", Tipo: NULL, Valor: 0, Mutable: false}
}

func (env Environment) SetVariable(id string, value Symbol) Symbol {
	//value.Id = id
	var tmpEnv Environment
	tmpEnv = env
	for {
		if variable, ok := tmpEnv.Variables[id]; ok {
			if tmpEnv.Variables[id].Mutable {
				if tmpEnv.Variables[id].Tipo == value.Tipo {
				
					tmpEnv.Variables[id] = value
					return variable
				} else {
					fmt.Println("El tipo es incorrecto")
					return Symbol{Lin: 0, Col: 0, Id: "", Tipo: NULL, Valor: 0, Mutable: false}
				}
			} else {
				fmt.Println("La variable", id, " no es mutable")
				return Symbol{Lin: 0, Col: 0, Id: "", Tipo: NULL, Valor: 0, Mutable: false}
			}
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("La variable ", id, " no existe")
	return Symbol{Lin: 0, Col: 0, Id: "", Tipo: NULL, Valor: 0, Mutable: false}
}

func (env Environment) KeepVector(id string, value SymbolVector) {
	
	if variable, ok := env.Vectors[id]; ok {
		fmt.Println("El vector "+id+" ya existe", variable)
		return
	}
	env.Vectors[id] = value
	
}

func (env Environment) SetVector(id string, value SymbolVector) SymbolVector {
	//value.Id = id
	var tmpEnv Environment
	tmpEnv = env
	for {
		
		if variable, ok := tmpEnv.Vectors[id]; ok {
					
					tmpEnv.Vectors[id] = value
					return variable
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("La variable ", id, " no existe")
	return SymbolVector{Lin: 0, Col: 0, Id: "", Tipo: NULL, Valor: nil}
}



func (env Environment) GetVector(id string) SymbolVector {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if variable, ok := tmpEnv.Vectors[id]; ok {
			
			return variable
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("La variable ", id, " no existe en este environment")
	return SymbolVector{Lin: 0, Col: 0, Id: "", Tipo: NULL, Valor: nil, Transfer: NULL}
}

 
func (env Environment) KeepFunction(id string, value SymbolFunction) {
	if variable, ok := env.Functions[id]; ok {
		fmt.Println("La funcion " + variable.Id + " ya existe")
		return
	}
	env.Functions[id] = value
}

func (env Environment) GetFunction(id string) SymbolFunction {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if variable, ok := tmpEnv.Functions[id]; ok {
			return variable
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("La funcion ", id, " no existe o es una funcion privada..")
	return SymbolFunction{TipoReturn: NULL}
}

func (env Environment) VerifyFunc() bool {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if strings.Contains(tmpEnv.NameEnv, "FUNCTION") {
			return true
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("la sentencia tiene que estar dentro de un ciclo")
	return false
}


