package codegen

import "fmt"

type PythonStmtGenerator struct {
}

func (s *PythonStmtGenerator) Print(value string) string {
	return s.FunctionCall("print", value)
}

func (s *PythonStmtGenerator) FunctionCall(name string, value string) string {
	return fmt.Sprintf("%s(%s)", name, value)
}

func (s *PythonStmtGenerator) SolutionCallPrefix() string {
	return "Solution()."
}

func (s *PythonStmtGenerator) VarAssignment(arg *arg, idx int) (string, string) {
	name := fmt.Sprintf("%s%d", RandomName(), idx)
	return name, fmt.Sprintf("%s = %s", name, arg.Value)
}
