package compiler

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/pkg/errors"
)

type CompileEvent struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

// Compiler is an interface for compiling code
type Compiler interface {
	Compile(code string) (string, error)
}

// compile command generator function
type commandFn func(filename string, outBinName string) (string, []string)

// intermediate compile function
type compileFn func(filename string) (string, error)

type GenericCompiler struct {
	fn  compileFn
	ext string
}

func NewGenericCompiler(fn compileFn, ext string) *GenericCompiler {
	return &GenericCompiler{
		fn:  fn,
		ext: ext,
	}
}

func (c *GenericCompiler) Compile(code string) (string, error) {
	r := rand.NewSource(time.Now().UnixNano())
	// Create a tempfile with the code in it
	file, err := ioutil.TempFile("/tmp", fmt.Sprintf("tmp_code_%d_%d_*%s", rand.New(r).Int(), time.Now().UnixNano(), c.ext))
	defer os.Remove(file.Name())
	if err != nil {
		return "tempfile err", err
	}

	err = ioutil.WriteFile(file.Name(), []byte(code), 0644)
	if err != nil {
		return "writefile err", err
	}

	return c.fn(file.Name())
}

// OneStepCompileFn just runs a single step compile
func OneStepCompileFn(command commandFn) compileFn {
	return func(filename string) (string, error) {
		// the output bin name isn't used
		name, args := command(filename, "")
		cmd := exec.Command(name, args...)
		cmd.Dir = "/tmp"
		out, err := cmd.CombinedOutput()
		if err != nil {
			return string(out), errors.Wrap(err, string(out))
		}

		return string(out), nil
	}
}

// TwoStepCompileFn first compiles then runs the generated binary
func TwoStepCompileFn(command commandFn) compileFn {
	return func(filename string) (string, error) {
		binName := fileNameWithoutExtension(filename)
		name, args := command(filename, binName)
		cmd := exec.Command(name, args...)
		cmd.Dir = "/tmp"
		out, err := cmd.CombinedOutput()
		if err != nil {
			return string(out), errors.Wrap(err, string(out))
		}

		// Remove the compiled binary after we execute it
		defer os.Remove(binName)

		err = os.Chmod(filename, 0777)
		if err != nil {
			return "chmod err", err
		}

		cmd = exec.Command(binName)
		cmd.Dir = "/tmp"
		out, err = cmd.CombinedOutput()
		if err != nil {
			return string(out), errors.Wrap(err, string(out))
		}

		return string(out), nil
	}
}
