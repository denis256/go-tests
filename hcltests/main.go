package main

import (
	"encoding/json"
	"fmt"

	_ "github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func main() {
	// Create a cty.Value (example: a map value)
	value := cty.MapVal(map[string]cty.Value{
		"key1": cty.StringVal("value1"),
	})

	// Convert cty.Value to HCL string
	hclString, err := ctyValueToHCLString(value)
	if err != nil {
		fmt.Println("Error converting cty.Value to HCL string:", err)
		return
	}

	// Print the HCL string
	fmt.Printf("HCL String:\n%s\n", hclString)
}

// ctyValueToHCLString converts a cty.Value to an HCL string
func ctyValueToHCLString(value cty.Value) (string, error) {
	// Convert cty.Value to JSON
	jsonBytes, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error marshaling cty.Value to JSON: %w", err)
	}

	// Create an HCL block with the JSON string
	hclString := fmt.Sprintf("example = %s", string(jsonBytes))

	return hclString, nil
}
