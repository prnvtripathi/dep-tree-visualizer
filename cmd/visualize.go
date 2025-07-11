package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/prnvtripathi/dep-tree-visualizer/internal/parser"
	"github.com/prnvtripathi/dep-tree-visualizer/utils"
)

var source string
var format string
var outputFileName string

var visualizeCmd = &cobra.Command{
	Use:     "visualize",
	Aliases: []string{"visualise"},
	Short:   "Visualize dependencies",
	Long: `The visualize command generates a visual representation of the dependency tree.
It supports tree output with colors and exporting to JSON.`,
	Run: func(cmd *cobra.Command, args []string) {
		if source == "" {
			fmt.Println("Please provide a path to go.mod using -s or --source")
			return
		}
		if format != "json" && format != "tree" {
			fmt.Println("Supported formats: tree, json")
			return
		}

		f := parser.ParseModFile(source)
		root := utils.BuildDependencyTree(f)

		switch format {
		case "tree":
			utils.PrintTree(root)
		case "json":
			if outputFileName != "" {
				utils.ExportJSON(root, outputFileName)
			} else {
				utils.PrintJSON(root)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(visualizeCmd)
	visualizeCmd.Flags().StringVarP(&source, "source", "s", "", "Path to the mod file to read from")
	visualizeCmd.Flags().StringVarP(&format, "format", "f", "tree", "Output format: tree, json")
	visualizeCmd.Flags().StringVarP(&outputFileName, "output", "o", "", "Output file name (for JSON only)")
}
