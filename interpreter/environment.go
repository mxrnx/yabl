package interpreter

import . "github.com/knarka/yabl/expr"

type Bind struct {
	name  string
	value Value
}

type Environment []Bind

func (env Environment) lookup(name string) Value {
	for _, bind := range env {
		if bind.name == name {
			return bind.value
		}
	}
	panic("no such bind!")
}
