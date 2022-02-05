package compiler

import (
	"context"
	"testing"

	"github.com/amatski/bisky/judge/codegen"

	"github.com/stretchr/testify/assert"
)

func TestGoCompile(t *testing.T) {
	r := NewRequestHandler()
	ctx := context.Background()
	code := "package main\nimport \"fmt\"\nfunc main(){fmt.Println(\"hi\")}"
	out, err := r.HandleCompileRequest(ctx, CompileEvent{
		Language: codegen.Go,
		Code:     code,
	})
	assert.NoError(t, err)
	assert.Equal(t, "hi\n", out)
}

func TestCppCompile(t *testing.T) {
	r := NewRequestHandler()
	ctx := context.Background()
	code := "#include <stdio.h>\nint main(){printf(\"hi\");}"
	out, err := r.HandleCompileRequest(ctx, CompileEvent{
		Language: codegen.Cpp,
		Code:     code,
	})
	assert.NoError(t, err)
	assert.Equal(t, "hi", out)
}

func TestPythonCompile(t *testing.T) {
	r := NewRequestHandler()
	ctx := context.Background()
	code := "print(515+1)"
	out, err := r.HandleCompileRequest(ctx, CompileEvent{
		Language: codegen.Python,
		Code:     code,
	})
	assert.NoError(t, err)
	assert.Equal(t, "516\n", out)
}

func TestJavascriptCompile(t *testing.T) {
	r := NewRequestHandler()
	ctx := context.Background()
	code := "console.log(512+4)"
	out, err := r.HandleCompileRequest(ctx, CompileEvent{
		Language: codegen.Javascript,
		Code:     code,
	})
	assert.NoError(t, err)
	assert.Equal(t, "516\n", out)
}
