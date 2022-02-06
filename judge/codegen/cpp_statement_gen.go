package codegen

import "fmt"

type CppStmtGenerator struct {
}

func (s *CppStmtGenerator) Print(value string) (string, error) {
	return fmt.Sprintf("cout << %s << \"\\n\";", value), nil
}

func (s *CppStmtGenerator) FunctionCall(name string, value string) string {
	return fmt.Sprintf("%s(%s)", name, value)
}

func (s *CppStmtGenerator) SolutionCallPrefix() string {
	return "Solution()."
}

func (s *CppStmtGenerator) VarAssignment(arg *arg, idx int) (string, string, error) {
	name := fmt.Sprintf("%s%d", RandomName(), idx)
	return name, fmt.Sprintf("%s %s = %s;", arg.Type, name, arg.Value), nil
}
