package util

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(v any) {
	json, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println("json marshal error")
	}
	fmt.Println(string(json))
}
