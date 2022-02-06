package codegen

import (
	"fmt"
)

type arg struct {
	Type  string
	Value string
}

func (a *arg) Literal() string {
	return fmt.Sprintf("%s%s", a.Type, a.Value)
}
