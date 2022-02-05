package judge

import (
	"testing"

	"github.com/amatski/bisky/judge/codegen"
	"github.com/amatski/bisky/judge/problem"

	"github.com/stretchr/testify/require"
)

func TestJudgeSolutionPython(t *testing.T) {
	f := problem.DiskFetcher{ParentDir: "../data"}
	handler := RequestHandler{
		Fetcher: &f,
	}

	twoSumCorrect := `
class Solution:
    def twoSum(self, nums, target):
        d = {}
        for i in range(0, len(nums)):
            n = nums[i]
            k = target-n
            if k in d:
                return [i, d[k]]
            d[n] = i
        return []
	`

	twoSumIncorrect := `
class Solution:
    def twoSum(self, nums, target):
        return nums
	`

	t.Run("judges correct python solution for 2sum", func(t *testing.T) {
		out, err := handler.JudgeSolution(JudgeRequest{
			Language: codegen.Python,
			Code:     twoSumCorrect,
			Problem:  "two_sum",
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

	t.Run("judges correct python solution for 2sum with multiple test cases", func(t *testing.T) {
		out, err := handler.JudgeSolution(JudgeRequest{
			Language: codegen.Python,
			Code:     twoSumCorrect,
			Problem:  "two_sum",
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

	t.Run("judges incorrect python solution for 2sum with multiple test cases", func(t *testing.T) {
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
			Language:  codegen.Python,
			Code:      twoSumIncorrect,
			Problem:   "two_sum",
			TestCases: testCases,
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
