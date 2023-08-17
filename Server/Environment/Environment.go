package environment

import "fmt"

type Environment struct {
	//Anterior *Environment
	Anterior  interface{}
	Variables map[string]Symbol
	NameEnv   string
}

func NewEnvironment(ant interface{}, ide string) Environment {
	return Environment{
		Anterior:  ant,
		Variables: make(map[string]Symbol),
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
	fmt.Println("La variable ", id, " no existe HEREEEE")
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
					//fmt.Println("ASIG EN TABLA ", tmpEnv.Tabla[id].TipoArr)
					//fmt.Println("asig val entra", value.TipoArr)
				//	value.TipoArr = tmpEnv.Tabla[id].TipoArr
				//	value.ExtraTipo = tmpEnv.Tabla[id].ExtraTipo
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
