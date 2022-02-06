package codegen

import "fmt"

type GoStmtGenerator struct {
}

func (s *GoStmtGenerator) Print(value string) string {
	return s.FunctionCall("fmt.Println", value)
}

func (s *GoStmtGenerator) FunctionCall(name string, value string) string {
	return fmt.Sprintf("%s(%s)", name, value)
}

func (s *GoStmtGenerator) SolutionCallPrefix() string {
	// Go directly invokes the function so the prefix is empty
	return ""
}

func (s *GoStmtGenerator) VarAssignment(arg *arg, idx int) (string, string) {
	name := fmt.Sprintf("%s%d", RandomName(), idx)
	return name, fmt.Sprintf("%s := %s", name, arg.Literal())
}

func (s *GoStmtGenerator) ToArg(value string) *arg {
	res := numberList.FindAllStringSubmatch(value, -1)
	if len(res) > 0 {
		elements := res[0][1]
		sliceType := "int"
		if decimal.MatchString(value) {
			sliceType = "float64"
		}

		return &arg{
			Type:  fmt.Sprintf("[]%s", sliceType),
			Value: fmt.Sprintf("{%s}", elements),
		}
	}

	return &arg{
		Type:  "",
		Value: value,
	}
}
