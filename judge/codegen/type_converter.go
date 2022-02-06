package codegen

import (
	"fmt"
	"regexp"
	"strings"
)

// ConvertType just converts the given value depending on the
// output type. currently only used for lists
func ConvertType(outputType string, value string) string {
	if outputType != Integers && outputType != Doubles {
		return value
	}

	re := regexp.MustCompile(`((-|)\d+(\.\d+|))`)
	nums := re.FindAllString(value, -1)

	return fmt.Sprintf("[%s]", strings.Join(nums, ","))
}
