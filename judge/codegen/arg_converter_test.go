package codegen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type argTestCase struct {
	input    string
	expected string
}

func TestConvertArg(t *testing.T) {
	t.Run("converts list to vector in c++", func(t *testing.T) {
		tests := []argTestCase{
			{
				input:    "[1,2,3]",
				expected: "vector<int>{1,2,3}",
			},
			{
				input:    "[1, 2,3]",
				expected: "vector<int>{1, 2,3}",
			},
			{
				input:    "[1,2,3,4,5]",
				expected: "vector<int>{1,2,3,4,5}",
			},
			{
				input:    "[1]",
				expected: "vector<int>{1}",
			},
			{
				input:    "[]",
				expected: "vector<int>{}",
			},
			{
				input:    "[1,2.5,3]",
				expected: "vector<double>{1,2.5,3}",
			},
			{
				input:    "[1.0,2.5,3.0]",
				expected: "vector<double>{1.0,2.5,3.0}",
			},
			{
				input:    "[133.55559]",
				expected: "vector<double>{133.55559}",
			},
		}

		for _, test := range tests {
			out := toArg(Cpp, test.input)
			assert.NotNil(t, out)
			assert.Equal(t, test.expected, out.Literal())
		}
	})

	t.Run("doesnt convert list in python", func(t *testing.T) {
		tests := []argTestCase{
			{
				input:    "[1,2,3]",
				expected: "[1,2,3]",
			},
			{
				input:    "[1, 2,3]",
				expected: "[1, 2,3]",
			},
			{
				input:    "[1]",
				expected: "[1]",
			},
			{
				input:    "[]",
				expected: "[]",
			},
		}

		for _, test := range tests {
			out := toArg(Python, test.input)
			assert.NotNil(t, out)
			assert.Equal(t, test.expected, out.Literal())
		}
	})

	t.Run("converts list in go", func(t *testing.T) {
		tests := []argTestCase{
			{
				input:    "[1,2,3]",
				expected: "[]int{1,2,3}",
			},
			{
				input:    "[1, 2,3]",
				expected: "[]int{1, 2,3}",
			},
			{
				input:    "[1]",
				expected: "[]int{1}",
			},
			{
				input:    "[]",
				expected: "[]int{}",
			},
		}

		for _, test := range tests {
			out := toArg(Go, test.input)
			assert.NotNil(t, out)
			assert.Equal(t, test.expected, out.Literal())
		}
	})

	t.Run("doesnt convert POD types", func(t *testing.T) {
		pods := []string{"1.20f", "100", "\"string\"", "0.0"}
		for _, lang := range Languages {
			for _, pod := range pods {
				out := toArg(lang, pod)
				assert.NotNil(t, out)
			}
		}
	})
}
