package codegen

import (
	"fmt"
	"regexp"
	"strings"
)

type Code struct {
	codeSb   strings.Builder
	language string
}

func NewCode(language string) (*Code, error) {
	if _, ok := LanguageToExt[language]; !ok {
		return nil, ErrInvalidLanguage
	}
	return &Code{language: language}, nil
}

func (g *Code) AddCode(code string) {
	g.codeSb.WriteString("\n") // just to be safe
	g.codeSb.WriteString(code)
}

func (g *Code) Code() string {
	return g.codeSb.String()
}

// AddCodeForTestCase returns what would be appended to the solution to output the solution's answer for a test case
func (g *Code) InjectValue(key string, value string) {
	re := regexp.MustCompile(fmt.Sprintf(`\$\{%s\}`, key))
	replaced := re.ReplaceAllString(g.codeSb.String(), value)
	g.codeSb = strings.Builder{}
	g.codeSb.WriteString(replaced)
}
