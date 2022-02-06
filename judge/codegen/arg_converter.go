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
	res := numberList.FindAllStringSubmatch(s, -1)
	if language == Cpp {
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

	if language == Go {
		if len(res) > 0 {
			elements := res[0][1]
			sliceType := "int"
			if decimal.MatchString(s) {
				sliceType = "float64"
			}

			return &arg{
				Type:  fmt.Sprintf("[]%s", sliceType),
				Value: fmt.Sprintf("{%s}", elements),
			}
		}
	}

	if language == Javascript {
		return &arg{
			Type:  "let",
			Value: s,
		}
	}

	return &arg{
		Type:  "",
		Value: s,
	}
}

func (a *arg) Literal() string {
	return fmt.Sprintf("%s%s", a.Type, a.Value)
}
