package codegen

import (
	"fmt"
	"regexp"
	"strings"
)

func ConvertType(outputType string, answer string) string {
	if outputType != Integers && outputType != Doubles {
		return answer
	}

	re := regexp.MustCompile(`((-|)\d+(\.\d+|))`)
	nums := re.FindAllString(answer, -1)

	return fmt.Sprintf("[%s]", strings.Join(nums, ","))
}
