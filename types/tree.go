package types

type Node struct {
	Name     string
	Version  string
	Children []*Node
}
