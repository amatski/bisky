package codegen

import "errors"

type StatementGenerator interface {
	Print(value string) string
	// FunctionCall just returns a function call, doesn't necessarily include a semicolon
	FunctionCall(name string, value string) string
	SolutionCallPrefix() string
	VarAssignment(arg *arg, idx int) (string, string)
	ToArg(string) *arg
}

var (
	StmtGenerators = map[string]StatementGenerator{
		Cpp:        &CppStmtGenerator{},
		Python:     &PythonStmtGenerator{},
		Go:         &GoStmtGenerator{},
		Javascript: &JavascriptStmtGenerator{},
	}
	ErrMissingGenerator = errors.New("missing statement generator")
)

func GetStatementGenerator(language string) (StatementGenerator, error) {
	g, ok := StmtGenerators[language]
	if !ok {
		return nil, ErrMissingGenerator
	}
	return g, nil
}
