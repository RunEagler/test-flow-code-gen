package testflow

import "strings"

// ToUpperCamelCase convert snake case string to upper camel case string
func ToUpperCamelCase(snakeParamName string) string {
	terms := strings.Split(snakeParamName, "_")
	upperNames := make([]string, 0)
	for _, term := range terms {
		upperNames = append(upperNames, strings.ToUpper(string(term[0]))+term[1:])
	}
	return strings.Join(upperNames, "")
}

// ToLowerCamelCase convert snake case string to lower camel case string
func ToLowerCamelCase(snakeParamName string) string {
	upperCameCase := ToUpperCamelCase(snakeParamName)
	return strings.ToLower(string(upperCameCase[0])) + upperCameCase[1:]
}
