package codegen

type StatementGenerator interface {
	Print(value string) string
	// FunctionCall just returns a function call, doesn't necessarily include a semicolon
	FunctionCall(name string, value string) string
	SolutionCallPrefix() string
	VarAssignment(arg *arg, idx int) (string, string)
}

var (
	StmtGenerators = map[string]StatementGenerator{
		Cpp:    &CppStmtGenerator{},
		Python: &PythonStmtGenerator{},
		Go:     &GoStmtGenerator{},
	}
)

func GetStatementGenerator(language string) (StatementGenerator, error) {
	g, ok := StmtGenerators[language]
	if !ok {
		return nil, ErrInvalidLanguage
	}
	return g, nil
}
