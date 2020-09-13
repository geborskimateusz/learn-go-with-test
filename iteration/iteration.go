package iteration

import (
	"strings"
)

func AppendViaAddAndAssignment(character string) string {
	var str string

	for i := 0; i < 5; i++ {
		str += character
	}

	return str
}

func AppendViaSB(character string) string {
	var str strings.Builder

	for i := 0; i < 5; i++ {
		str.WriteString(character)
	}

	return str.String()
}
