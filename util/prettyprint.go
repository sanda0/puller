package util

import (
	"encoding/json"
	"fmt"
)

func PrettyPrintStruct(data interface{}) string {
	prettyJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Sprintf("Error: %s", err)
	}
	return string(prettyJSON)
}
