package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonStr := `{"name": "John", "age": 30, "is_student": false}`
	var result map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		fmt.Println(err)
		return
	}

	generateStruct(result)
}

func generateStruct(data map[string]interface{}) {
	fmt.Println("type GeneratedStruct struct {")
	for key, value := range data {
		fmt.Printf("    %s %s `json:\"%s\"`\n", capitalize(key), goType(value), key)
	}
	fmt.Println("}")
}

func goType(value interface{}) string {
	switch value.(type) {
	case string:
		return "string"
	case float64:
		return "int"
	case bool:
		return "bool"
	default:
		return "interface{}"
	}
}

func capitalize(s string) string {
	if len(s) == 0 {
		return ""
	}
	return string(s[0]-32) + s[1:] // Приводим к верхнему регистру
}
