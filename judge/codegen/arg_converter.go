package codegen

import (
	"fmt"
	"regexp"
)

var (
	// regex that matches a list of numbers
	numberList = regexp.MustCompile(`\[(((\s)*(\d+(\.|)(\d+|)(\s)*)(,)*)*)\]`)
	decimal    = regexp.MustCompile(`(\d+\.\d*)`)
)

type arg struct {
	Type  string
	Value string
}

func (a *arg) Literal() string {
	return fmt.Sprintf("%s%s", a.Type, a.Value)
}
