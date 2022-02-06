package codegen

import (
	"regexp"
	"strings"
)

var (
	// regex that matches a list of numbers
	numberList = regexp.MustCompile(`\[(((\s)*(\d+(\.|)(\d+|)(\s)*)(,)*)*)((\s)*(\d+(\.|)(\d+|)(\s)*))\]`)
	decimal    = regexp.MustCompile(`(\d+\.\d*)`)
	stringList = regexp.MustCompile(`\[(((\s)*("[^"]+")(\s)*(,)*)*)((\s)*("[^"]+")(\s)*)\]`)
)

type NumberList struct {
	Elements string
}

func isNumberList(value string) *NumberList {
	res := numberList.FindAllStringSubmatch(value, -1)
	if len(res) > 0 {
		elements := res[0][1]
		if len(res[0]) >= 10 {
			elements += res[0][9]
		}
		return &NumberList{
			Elements: elements,
		}
	}
	return nil
}

func isStringList(value string) bool {
	return stringList.MatchString(value)
}

func argsForInput(generator StatementGenerator, input string) ([]*arg, error) {
	args := strings.Split(input, "\n")
	filteredArgs := []*arg{}
	for _, arg := range args {
		t := strings.TrimSpace(arg)
		if arg != "\n" && t != "" {
			filteredArgs = append(filteredArgs, generator.ToArg(t))
		}
	}

	return filteredArgs, nil
}
