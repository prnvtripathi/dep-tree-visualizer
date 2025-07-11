package types

type NodeType int

const (
	RequireNode NodeType = iota
	ReplaceNode
	ExcludeNode
)

type Node struct {
	Name     string   `json:"name"`
	Version  string   `json:"version,omitempty"`
	Children []*Node  `json:"children,omitempty"`
	Type     NodeType `json:"-"`
}
