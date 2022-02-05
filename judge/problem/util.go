package problem

import (
	"strings"
)

func SnakeCaseToCamelCase(snakeCase string) (camelCase string) {
	isToUpper := false

	for k, v := range snakeCase {
		if k == 0 {
			camelCase = string(snakeCase[0])
		} else {
			if isToUpper {
				camelCase += strings.ToUpper(string(v))
				isToUpper = false
			} else {
				if v == '_' {
					isToUpper = true
				} else {
					camelCase += string(v)
				}
			}
		}
	}
	return
}
