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

func isNumberList(value string) bool {
	return numberList.MatchString(value)
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
