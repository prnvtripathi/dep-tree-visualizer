package utils

import (
	"fmt"

	"github.com/disiqueira/gotree"
	"github.com/fatih/color"
	"golang.org/x/mod/modfile"

	"github.com/prnvtripathi/dep-tree-visualizer/types"
)

func BuildDependencyTree(f *modfile.File) *types.Node {
	root := &types.Node{Name: "Dependencies"}

	for _, req := range f.Require {
		root.Children = append(root.Children, &types.Node{
			Name:    req.Mod.Path,
			Version: req.Mod.Version,
			Type:    types.RequireNode,
		})
	}

	if len(f.Replace) > 0 {
		replace := &types.Node{Name: "Replacements", Type: types.ReplaceNode}
		for _, rep := range f.Replace {
			replace.Children = append(replace.Children, &types.Node{
				Name: fmt.Sprintf("%s => %s", rep.Old.Path, rep.New.Path),
				Type: types.ReplaceNode,
			})
		}
		root.Children = append(root.Children, replace)
	}

	if len(f.Exclude) > 0 {
		exclude := &types.Node{Name: "Excludes", Type: types.ExcludeNode}
		for _, ex := range f.Exclude {
			exclude.Children = append(exclude.Children, &types.Node{
				Name:    ex.Mod.Path,
				Version: ex.Mod.Version,
				Type:    types.ExcludeNode,
			})
		}
		root.Children = append(root.Children, exclude)
	}

	return root
}

func PrintTree(n *types.Node) {
	tree := convertToGoTree(n)
	fmt.Println(tree.Print())
}

func convertToGoTree(n *types.Node) gotree.Tree {
	label := n.Name
	if n.Version != "" {
		label += "@" + n.Version
	}

	switch n.Type {
	case types.RequireNode:
		label = color.GreenString(label)
	case types.ReplaceNode:
		label = color.YellowString(label)
	case types.ExcludeNode:
		label = color.RedString(label)
	}

	t := gotree.New(label)
	for _, child := range n.Children {
		t.AddTree(convertToGoTree(child))
	}
	return t
}
