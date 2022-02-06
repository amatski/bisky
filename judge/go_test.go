package judge

import (
	"log"
	"testing"

	"github.com/amatski/bisky/judge/codegen"
	"github.com/amatski/bisky/judge/problem"

	"github.com/stretchr/testify/require"
)

func TestJudgeSolutionGo(t *testing.T) {
	f := problem.DiskFetcher{ParentDir: "../data"}
	handler := RequestHandler{
		Fetcher: &f,
	}

	twoSumIncorrect := `
func twoSum(nums []int, target int) []int {
	return []int{target, nums[0]-target}
}
	`

	t.Run("judges correct go solution for 2sum", func(t *testing.T) {
		out, err := handler.JudgeSolution(JudgeRequest{
			Language:   codegen.Go,
			Code:       twoSumIncorrect,
			Problem:    "two_sum",
			OutputType: codegen.Integers,
			TestCases: []*problem.TestCase{
				{
					Input: `[1,2,3,4,5]
					5
				`,
					ExpectedOutput: []string{`[5,-4]`},
				},
			},
		})
		require.NoError(t, err)
		require.NotNil(t, out)
		log.Println(out.Results[0].Stdout)
		require.True(t, out.Results[0].Passed)
		require.Equal(t, "[5,-4]", out.Results[0].Stdout)
	})

}
