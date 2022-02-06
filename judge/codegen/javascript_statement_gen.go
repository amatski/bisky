package codegen

import "fmt"

type JavascriptStmtGenerator struct {
}

func (s *JavascriptStmtGenerator) Print(value string) string {
	return fmt.Sprintf("console.log(%s);", value)
}

func (s *JavascriptStmtGenerator) FunctionCall(name string, value string) string {
	return fmt.Sprintf("%s(%s)", name, value)
}

func (s *JavascriptStmtGenerator) SolutionCallPrefix() string {
	return ""
}

func (s *JavascriptStmtGenerator) VarAssignment(arg *arg, idx int) (string, string) {
	name := fmt.Sprintf("%s%d", RandomName(), idx)
	return name, fmt.Sprintf("%s %s = %s", arg.Type, name, arg.Value)
}

func (s *JavascriptStmtGenerator) ToArg(value string) *arg {
	return &arg{
		Type:  "let",
		Value: value,
	}
}

func (s *JavascriptStmtGenerator) ToArgFromNumberList(list *NumberList) *arg {
	return &arg{
		Type:  "let",
		Value: fmt.Sprintf("[%s]", list.Elements),
	}
}
