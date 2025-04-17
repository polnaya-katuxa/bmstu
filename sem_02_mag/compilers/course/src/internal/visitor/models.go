package visitor

import "github.com/llir/llvm/ir/value"

type variableLeft struct {
	name string
	key  value.Value
}

type expression struct {
	value value.Value
}
