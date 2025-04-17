package ast

type Node interface {
}

type String interface {
}

type StringNode struct {
	Value string
}

type Boolean interface {
}

type BoolNode struct {
	Value bool
}

type Number interface {
}

type FloatNode struct {
	Value float64
}

type IntNode struct {
	Value int64
}

type OperatorPowerNode struct {
	Left, Right Number
}

type OperatorMulNode struct {
	Left, Right Number
}

type OperatorDivNode struct {
	Left, Right Number
}

type OperatorModNode struct {
	Left, Right *IntNode
}

type OperatorIntDivNode struct {
	Left, Right *IntNode
}

type OperatorAddNode struct {
	Left, Right Number
}

type OperatorSubNode struct {
	Left, Right Number
}

type OperatorStrCatNode struct {
	Left, Right *StringNode
}

type OperatorNotNode struct {
	Operand Boolean
}

type OperatorUnaryMinusNode struct {
	Operand Number
}

type OperatorStrLenNode struct {
	Operand String
}

type OperatorGreaterNode struct {
	Left, Right Number
}

type OperatorGreaterOrEqualNode struct {
	Left, Right Number
}

type OperatorLessNode struct {
	Left, Right Number
}

type OperatorLessOrEqualNode struct {
	Left, Right Number
}

type OperatorEqualNode struct {
	Left, Right Number
}

type OperatorNotEqualNode struct {
	Left, Right Number
}

type OperatorOrNode struct {
	Left, Right Boolean
}

type OperatorAndNode struct {
	Left, Right Boolean
}
