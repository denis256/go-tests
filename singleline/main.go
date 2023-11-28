package main

import (
	"fmt"
	"strings"
)

func mapToHCL(m map[string]interface{}) string {
	var attributes []string

	for key, value := range m {
		attributes = append(attributes, fmt.Sprintf(`%s=%s`, key, formatValue(value)))
	}

	return fmt.Sprintf("{%s}", strings.Join(attributes, ","))
}

func formatValue(value interface{}) string {
	switch v := value.(type) {
	case string:
		// Escape double quotes within the string value
		escapedValue := strings.ReplaceAll(v, `"`, `\"`)
		return fmt.Sprintf(`\"%s\"`, escapedValue)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return fmt.Sprintf("%f", v)
	case bool:
		return fmt.Sprintf("%t", v)
	case map[string]interface{}:
		return mapToHCL(v)
	default:
		return fmt.Sprintf(`\"%v\"`, v)
	}
}

func main() {
	// Example usage
	myMap := map[string]interface{}{
		"assume_role": map[string]interface{}{
			"key1": "value1",
			"key2": 42,
			"key3": true,
			"key4": map[string]interface{}{
				"nestedKey1": "nestedValue1",
				"nestedKey2": 99,
			},
		},
	}

	hclCode := mapToHCL(myMap)
	fmt.Println(hclCode)
}
