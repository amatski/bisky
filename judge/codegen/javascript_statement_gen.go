package codegen

import "fmt"

type JavascriptStmtGenerator struct {
}

func (s *JavascriptStmtGenerator) Print(value string) (string, error) {
	return fmt.Sprintf("console.log(%s)", value), nil
}

func (s *JavascriptStmtGenerator) FunctionCall(name string, value string) string {
	return fmt.Sprintf("%s(%s)", name, value)
}

func (s *JavascriptStmtGenerator) SolutionCallPrefix() string {
	return ""
}

func (s *JavascriptStmtGenerator) VarAssignment(arg *arg, idx int) (string, string, error) {
	name := fmt.Sprintf("%s%d", RandomName(), idx)
	return name, fmt.Sprintf("%s = %s", name, arg.Value), nil
}
