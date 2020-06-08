package interpreter

import . "github.com/knarka/yabl/expr"

type Pointer struct {
	name    string
	address int
}

type Bind struct {
	address int
	value   Value
}

type Environment []Pointer
type Store []Bind

var maxAddr = 0

func (env Environment) lookup(name string) int {
	for _, pointer := range env {
		if pointer.name == name {
			return pointer.address
		}
	}
	panic("no such pointer!")
}

func (sto Store) fetch(address int) Value {
	for _, bind := range sto {
		if bind.address == address {
			return bind.value
		}
	}
	panic("no such bind!")
}

func newAddr() int {
	maxAddr += 1
	return maxAddr - 1
}
