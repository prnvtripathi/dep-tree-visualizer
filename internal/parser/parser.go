package parser

import (
	"fmt"
	"os"

	"golang.org/x/mod/modfile"
)

func main() {
	data, err := os.ReadFile("go.mod")
	if err != nil {
		panic(err)
	}

	f, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		panic(err)
	}

	// Printing the module path
	fmt.Println("File path: ", f.Module.Mod.Path)

	// Print all required modules
	for _, req := range f.Require {
		fmt.Printf("Require: %s %s\n", req.Mod.Path, req.Mod.Version)
	}

	// Print replace directives
	for _, rep := range f.Replace {
		fmt.Printf("Replace: %s => %s\n", rep.Old.Path, rep.New.Path)
	}

	// Print exclude directives
	for _, ex := range f.Exclude {
		fmt.Printf("Exclude: %s %s\n", ex.Mod.Path, ex.Mod.Version)
	}
}

// Main functionality
func ParseModFile(path string) *modfile.File {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	f, err := modfile.Parse(path, data, nil)
	if err != nil {
		panic(err)
	}

	return f
}
