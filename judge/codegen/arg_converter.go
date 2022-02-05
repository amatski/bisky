package codegen

import (
	"fmt"
	"regexp"
)

var (
	numberList = regexp.MustCompile(`\[(((\s)*(\d+(\.|)(\d+|)(\s)*)(,)*)*)\]`)
	decimal    = regexp.MustCompile(`(\d+\.\d*)`)
)

type arg struct {
	Type  string
	Value string
}

func toArg(language string, s string) *arg {
	if language == Cpp {
		res := numberList.FindAllStringSubmatch(s, -1)
		if len(res) > 0 {
			elements := res[0][1]
			vecType := "int"
			if decimal.MatchString(s) {
				vecType = "double"
			}

			return &arg{
				Type:  fmt.Sprintf("vector<%s>", vecType),
				Value: fmt.Sprintf("{%s}", elements),
			}
		}
		return &arg{
			Type:  "auto",
			Value: s,
		}
	}
	return &arg{
		Type:  "",
		Value: s,
	}
}

func (a *arg) Full() string {
	return fmt.Sprintf("%s%s", a.Type, a.Value)
}
