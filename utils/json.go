package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/prnvtripathi/dep-tree-visualizer/types"
)

func PrintJSON(root *types.Node) {
	data, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	fmt.Println(string(data))
}

func ExportJSON(root *types.Node, filename string) {
	data, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
	fmt.Printf("✅ Exported dependency tree to %s\n", filename)
}
