package codegen

import (
	"fmt"
	"strings"

	"github.com/amatski/bisky/judge/problem"
)

func printStatement(language string, value string) (string, error) {
	if language == Cpp {
		return fmt.Sprintf("cout << %s << \"\\n\";", value), nil
	}

	if language == Python {
		return fmt.Sprintf("print(%s)", value), nil
	}
	return "", nil
}

func functionCall(name string, value string) string {
	return fmt.Sprintf("%s(%s)", name, value)
}

func varAssignment(language string, arg *arg, idx int) (string, string, error) {
	name := fmt.Sprintf("%s%d", RandomName(), idx)
	if language == Cpp {
		return name, fmt.Sprintf("%s %s = %s;", arg.Type, name, arg.Value), nil
	}

	if language == Python {
		return name, fmt.Sprintf("%s = %s", name, arg.Value), nil
	}
	return "", "", nil
}

func argsForInput(language string, input string) ([]*arg, error) {
	args := strings.Split(input, "\n")
	filteredArgs := []*arg{}
	for _, arg := range args {
		t := strings.TrimSpace(arg)
		if arg != "\n" && t != "" {
			filteredArgs = append(filteredArgs, toArg(language, t))
		}
	}

	return filteredArgs, nil
}

func TestCaseCalls(testCases []*problem.TestCase, language string, problemId string, secret string) (string, error) {
	body, err := NewCode(language)
	if err != nil {
		return "", err
	}
	tcIdx := 0
	for _, testCase := range testCases {
		secretPrintStmt, err := printStatement(language, fmt.Sprintf("\"%s\"", secret))
		if err != nil {
			return "", err
		}

		// clean up input into args specific for the language
		args, err := argsForInput(language, testCase.Input)
		if err != nil {
			return "", err
		}

		// create an assignment for each var from the converted args
		argNames := []string{}
		for _, arg := range args {
			name, assignment, err := varAssignment(language, arg, tcIdx)
			if err != nil {
				return "", err
			}
			body.AddCode(assignment)
			argNames = append(argNames, name)
			tcIdx++
		}

		body.AddCode(secretPrintStmt)
		fnStmt := functionCall("Solution()."+problem.SnakeCaseToCamelCase(problemId), strings.Join(argNames, ","))

		printAnswerStmt, err := printStatement(language, fnStmt)
		if err != nil {
			return "", err
		}
		body.AddCode(printAnswerStmt)
	}
	return body.Code(), nil
}
