package cmd

import (
	"fmt"
	"log"

	"github.com/disiqueira/gotree"
	"github.com/prnvtripathi/dep-tree-visualizer/internal/parser"
	"github.com/spf13/cobra"
	"golang.org/x/mod/modfile"
)

var source string

// visualizeCmd represents the visualize command
var visualizeCmd = &cobra.Command{
	Use:   "visualize",
	Short: "Visualize dependencies",
	Long: `The visualize command generates a visual representation of the dependency tree.
This command can be used to create graphs or charts that illustrate the relationships
between different packages and their dependencies, making it easier to understand
the structure of a project or system.`,
	Run: func(cmd *cobra.Command, args []string) {
		if source == "" {
			log.Println("Please provide a path to go.mod using -s or --source")
			return
		}
		fmt.Println("Visualizing dependencies from:", source)
		f := parser.ParseModFile(source)
		createTree(f)
	},
}

func init() {
	rootCmd.AddCommand(visualizeCmd)
	visualizeCmd.Flags().StringVarP(&source, "source", "s", "", "Path to the mod file to read from")
}

func createTree(f *modfile.File) {
	root := gotree.New("Dependencies")

	// Add required modules
	for _, req := range f.Require {
		root.Add(fmt.Sprintf("%s@%s", req.Mod.Path, req.Mod.Version))
	}

	// Add replace directives if any
	if len(f.Replace) > 0 {
		replacements := root.Add("Replacements")
		for _, rep := range f.Replace {
			replacements.Add(fmt.Sprintf("%s => %s", rep.Old.Path, rep.New.Path))
		}
	}

	// Add excluded modules if any
	if len(f.Exclude) > 0 {
		excludes := root.Add("Excludes")
		for _, ex := range f.Exclude {
			excludes.Add(fmt.Sprintf("%s@%s", ex.Mod.Path, ex.Mod.Version))
		}
	}

	fmt.Println(root.Print())

}
