package ast

type Node struct {
	Name     string  `json:name`
	Value    string  `json:name`
	Children []*Node `json:children`
}
