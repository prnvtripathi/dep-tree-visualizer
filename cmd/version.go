package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const CliVersion = "v0.1.0"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Know the installed version of deps-tree",
	Long:  `This command will help you to know the installed version of deps-tree`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deps-tree version:", CliVersion, "\nTo check and update the latest version, run 'deps-tree update'")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
