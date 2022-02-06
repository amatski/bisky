package judge

import (
	"testing"

	"github.com/amatski/bisky/judge/codegen"
	"github.com/amatski/bisky/judge/problem"

	"github.com/stretchr/testify/require"
)

func TestJudgeSolutionJavascript(t *testing.T) {
	f := problem.DiskFetcher{ParentDir: "../data"}
	handler := RequestHandler{
		Fetcher: &f,
	}

	twoSumCorrect := `
	/**
	* @param {number[]} nums
	* @param {number} target
	* @return {number[]}
	*/
   var twoSum = function(nums, target) {
	   return nums;
   };
	`

	t.Run("judges correct javascript solution for 2sum", func(t *testing.T) {
		out, err := handler.JudgeSolution(JudgeRequest{
			Language: codegen.Javascript,
			Code:     twoSumCorrect,
			Problem:  "two_sum",
			TestCases: []*problem.TestCase{
				{
					Input: `[1,2,3,4,5]
					5
				`,
					ExpectedOutput: []string{`[ 1, 2, 3, 4, 5 ]`},
				},
				{
					Input: `[1,2,3,4,5,6]
					5
				`,
					ExpectedOutput: []string{`[ 1, 2, 3, 4, 5, 6 ]`},
				},
			},
		})
		require.NoError(t, err)
		require.NotNil(t, out)
		require.True(t, out.Results[0].Passed)
		require.Equal(t, "[ 1, 2, 3, 4, 5 ]", out.Results[0].Stdout)
		require.True(t, out.Results[1].Passed)
	})

}
