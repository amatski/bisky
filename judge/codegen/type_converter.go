package codegen

import (
	"fmt"
	"regexp"
	"strings"
)

// ConvertType just converts the given value depending on the
// output type. currently only used for lists
func ConvertType(outputType string, value string) string {
	// have to check that the output is actually integers here
	// using regex
	// (number list)
	// have to adjust the regex to include spaces instead of commas more or less
	// otherwise it doesn't get converted

	if outputType != Integers && outputType != Doubles {
		return value
	}
	re := regexp.MustCompile(`((-|)\d+(\.\d+|))`)
	nums := re.FindAllString(value, -1)

	return fmt.Sprintf("[%s]", strings.Join(nums, ","))
}
