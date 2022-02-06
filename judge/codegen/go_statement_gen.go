package codegen

import "fmt"

type GoStmtGenerator struct {
}

func (s *GoStmtGenerator) Print(value string) (string, error) {
	return fmt.Sprintf("fmt.Println(%s)", value), nil
}

func (s *GoStmtGenerator) FunctionCall(name string, value string) string {
	return fmt.Sprintf("%s(%s)", name, value)
}

func (s *GoStmtGenerator) SolutionCallPrefix() string {
	// Go directly invokes the function so the prefix is empty
	return ""
}

func (s *GoStmtGenerator) VarAssignment(arg *arg, idx int) (string, string, error) {
	name := fmt.Sprintf("%s%d", RandomName(), idx)
	return name, fmt.Sprintf("%s := %s", name, arg.Literal()), nil
}
