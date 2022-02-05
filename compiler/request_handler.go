package compiler

import (
	"context"
	"errors"

	"github.com/amatski/algotrainer/backend/bisky/judge/codegen"
)

type RequestHandler struct {
	Compilers map[string]Compiler
}

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{
		Compilers: map[string]Compiler{
			codegen.Python: NewGenericCompiler(
				OneStepCompileFn(func(filename string, outBinName string) (string, []string) {
					return "python", []string{filename}
				}),
				codegen.LanguageToExt[codegen.Python],
			),
			codegen.Go: NewGenericCompiler(
				TwoStepCompileFn(func(filename string, outBinName string) (string, []string) {
					return "go", []string{"build", filename}
				}),
				codegen.LanguageToExt[codegen.Go],
			),
			codegen.Cpp: NewGenericCompiler(
				TwoStepCompileFn(func(filename string, outBinName string) (string, []string) {
					return "g++", []string{filename, "-std=c++11", "-Wno-everything", "-O3", "-o", outBinName}
				}),
				codegen.LanguageToExt[codegen.Cpp],
			),
			codegen.Javascript: NewGenericCompiler(
				OneStepCompileFn(func(filename string, outBinName string) (string, []string) {
					return "node", []string{filename}
				}),
				codegen.LanguageToExt[codegen.Javascript],
			),
		},
	}
}

func (l *RequestHandler) HandleCompileRequest(ctx context.Context, event CompileEvent) (string, error) {
	if _, ok := l.Compilers[event.Language]; !ok {
		return "could not find compile hook", errors.New("missing compile hook")
	}

	return l.Compilers[event.Language].Compile(event.Code)
}
