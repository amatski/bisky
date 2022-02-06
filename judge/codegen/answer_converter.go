package codegen

import (
	"regexp"
	"strings"
)

func ConvertAnswer(language string, outputType string, answer string) string {
	if outputType != Integers && outputType != Doubles {
		return answer
	}

	re := regexp.MustCompile(`(\d+(\.\d+|))`)
	nums := re.FindAllString(answer, -1)

	return strings.Join(nums, ",")
}
