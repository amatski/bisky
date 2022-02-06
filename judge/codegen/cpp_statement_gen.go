package codegen

import "fmt"

type CppStmtGenerator struct {
}

func (s *CppStmtGenerator) Print(value string) string {
	return fmt.Sprintf("cout << %s << \"\\n\";", value)
}

func (s *CppStmtGenerator) FunctionCall(name string, value string) string {
	return fmt.Sprintf("%s(%s)", name, value)
}

func (s *CppStmtGenerator) SolutionCallPrefix() string {
	return "Solution()."
}

func (s *CppStmtGenerator) VarAssignment(arg *arg, idx int) (string, string) {
	name := fmt.Sprintf("%s%d", RandomName(), idx)
	return name, fmt.Sprintf("%s %s = %s;", arg.Type, name, arg.Value)
}

func (s *CppStmtGenerator) ToArg(value string) *arg {
	list := isNumberList(value)
	if list != nil {
		vecType := "int"
		if decimal.MatchString(value) {
			vecType = "double"
		}

		return &arg{
			Type:  fmt.Sprintf("vector<%s>", vecType),
			Value: fmt.Sprintf("{%s}", list.Elements),
		}
	}
	return &arg{
		Type:  "auto",
		Value: value,
	}
}
