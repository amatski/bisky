package judge

import (
	"testing"

	"github.com/amatski/bisky/judge/codegen"
	"github.com/amatski/bisky/judge/problem"

	"github.com/stretchr/testify/require"
)

func TestJudgeSolutionPythonEncoded(t *testing.T) {
	f := problem.DiskFetcher{ParentDir: "../data"}
	handler := RequestHandler{
		Fetcher: &f,
	}

	twoSumCorrect := "CmNsYXNzIFNvbHV0aW9uOgogICAgZGVmIHR3b1N1bShzZWxmLCBudW1zLCB0YXJnZXQpOgogICAgICAgIGQgPSB7fQogICAgICAgIGZvciBpIGluIHJhbmdlKDAsIGxlbihudW1zKSk6CiAgICAgICAgICAgIG4gPSBudW1zW2ldCiAgICAgICAgICAgIGsgPSB0YXJnZXQtbgogICAgICAgICAgICBpZiBrIGluIGQ6CiAgICAgICAgICAgICAgICByZXR1cm4gW2ksIGRba11dCiAgICAgICAgICAgIGRbbl0gPSBpCiAgICAgICAgcmV0dXJuIFtdCgk="
	twoSumIncorrect := "CmNsYXNzIFNvbHV0aW9uOgogICAgZGVmIHR3b1N1bShzZWxmLCBudW1zLCB0YXJnZXQpOgogICAgICAgIHJldHVybiBudW1zCgk="

	t.Run("judges correct encoded python solution for 2sum", func(t *testing.T) {
		out, err := handler.JudgeSolution(JudgeRequest{
			Language:    codegen.Python,
			EncodedCode: &twoSumCorrect,
			Problem:     "two_sum",
			TestCases: []*problem.TestCase{
				{
					Input: `[1,2,3,4,5]
					5
				`,
					ExpectedOutput: []string{`[2, 1]`},
				},
			},
		})
		require.NoError(t, err)
		require.NotNil(t, out)
		require.True(t, out.Results[0].Passed)
		require.Equal(t, "[2, 1]", out.Results[0].Stdout)
	})

	t.Run("judges correct encoded python solution for 2sum with multiple test cases", func(t *testing.T) {
		out, err := handler.JudgeSolution(JudgeRequest{
			Language:    codegen.Python,
			EncodedCode: &twoSumCorrect,
			Problem:     "two_sum",
			TestCases: []*problem.TestCase{
				{
					Input: `[1,2,3,4,5]
					5
				`,
					ExpectedOutput: []string{`[2, 1]`},
				},
				{
					Input: `[1,2,3,4,5]
					1
				`,
					ExpectedOutput: []string{`[]`},
				},
				{
					Input: `[1,2,3,4,5]
					6
				`,
					ExpectedOutput: []string{`[1, 3]`, `[3, 1]`},
				},
				{
					Input: `[1,2,3,4,5]
					6
				`,
					ExpectedOutput: []string{`[]`},
				},
			},
		})
		require.NoError(t, err)
		require.NotNil(t, out)
		require.True(t, out.Results[0].Passed)
		require.True(t, out.Results[1].Passed)
		require.True(t, out.Results[2].Passed)
		require.False(t, out.Results[3].Passed)
	})

	t.Run("judges incorrect encoded python solution for 2sum with multiple test cases", func(t *testing.T) {
		testCases := []*problem.TestCase{
			{
				Input: `[1,2,3,4,5]
				5
				`,
				ExpectedOutput: []string{`[2, 1]`},
			},
			{
				Input: `
				[1,2,3,4,5]
					1
				`,
				ExpectedOutput: []string{`[]`},
			},
			{
				Input: `[1,2,3,4,5]
					6
				`,
				ExpectedOutput: []string{`[1, 3]`, `[3, 1]`},
			},
			{
				Input: `[1,2,3,4,5]
					6
				`,
				ExpectedOutput: []string{`[]`},
			},
		}
		out, err := handler.JudgeSolution(JudgeRequest{
			Language:    codegen.Python,
			EncodedCode: &twoSumIncorrect,
			Problem:     "two_sum",
			TestCases:   testCases,
		})
		require.NoError(t, err)
		require.NotNil(t, out)

		// 1:1 with number of test cases
		require.Equal(t, len(testCases), len(out.Results))
		require.False(t, out.Results[0].Passed)
		require.False(t, out.Results[1].Passed)
		require.False(t, out.Results[2].Passed)
		require.False(t, out.Results[3].Passed)
	})
}
