package codegen

import (
	"regexp"
	"strings"
)

var (
	// regex that matches a list of numbers
	numberList = regexp.MustCompile(`\[(((((\s)*(-|)(\d+(\.\d+|)(\s)*)(,)*)*)((\s)*((-|)\d+(\.\d+|)(\s)*)))|(\s)*)\]`)
	decimal    = regexp.MustCompile(`(\d+\.\d*)`)
	//stringList = regexp.MustCompile(`\[(((\s)*("[^"]+")(\s)*(,)*)*)((\s)*("[^"]+")(\s)*)\]`)
)

type NumberList struct {
	Elements string
	Decimal  bool
}

func isNumberList(value string) *NumberList {
	res := numberList.FindAllStringSubmatch(value, -1)
	if len(res) > 0 {
		elements := res[0][1]

		return &NumberList{
			Elements: elements,
			Decimal:  decimal.MatchString(elements),
		}
	}
	return nil
}

func argsForInput(generator StatementGenerator, input string) ([]*arg, error) {
	args := strings.Split(input, "\n")
	filteredArgs := []*arg{}
	for _, argg := range args {
		t := strings.TrimSpace(argg)
		if argg != "\n" && t != "" {
			list := isNumberList(t)
			var a *arg
			if list != nil {
				a = generator.ArgFromNumberList(list)
			} else {
				a = generator.Arg(t)
			}
			filteredArgs = append(filteredArgs, a)
		}
	}

	return filteredArgs, nil
}
