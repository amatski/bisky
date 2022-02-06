package codegen

import (
	"fmt"
	"strings"

	"github.com/amatski/bisky/judge/problem"
	"github.com/pkg/errors"
)

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

func TestCaseCalls(testCases []*problem.TestCase, language string, problemId string, secret string) (string, error) {
	body, err := NewCode(language)
	if err != nil {
		return "", err
	}

	generator, err := GetStatementGenerator(language)
	if err != nil {
		return "", errors.WithMessage(err, "getting statement generator")
	}

	tcIdx := 0
	for _, testCase := range testCases {
		secretPrintStmt := generator.Print(fmt.Sprintf("\"%s\"", secret))

		// clean up input into args specific for the language
		args, err := argsForInput(generator, testCase.Input)
		if err != nil {
			return "", err
		}

		// create an assignment for each var from the converted args
		argNames := []string{}
		for _, arg := range args {
			name, assignment := generator.VarAssignment(arg, tcIdx)
			body.AddCode(assignment)
			argNames = append(argNames, name)
			tcIdx++
		}

		body.AddCode(secretPrintStmt)
		fnStmt := generator.FunctionCall(generator.SolutionCallPrefix()+problem.SnakeCaseToCamelCase(problemId), strings.Join(argNames, ","))

		printAnswerStmt := generator.Print(fnStmt)
		body.AddCode(printAnswerStmt)
	}
	return body.Code(), nil
}
