package codegen

import "errors"

const (
	Python     = "Python"
	Cpp        = "Cpp"
	Go         = "Go"
	Javascript = "Javascript"
)

var (
	LanguageToExt map[string]string = map[string]string{
		Python:     ".py",
		Cpp:        ".cpp",
		Go:         ".go",
		Javascript: ".js",
	}

	Languages          = []string{Python, Cpp, Go, Javascript}
	ErrInvalidLanguage = errors.New("invalid language")
)
