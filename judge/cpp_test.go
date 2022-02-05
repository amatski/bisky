package judge

import (
	"testing"

	"github.com/amatski/bisky/judge/codegen"
	"github.com/amatski/bisky/judge/problem"

	"github.com/stretchr/testify/require"
)

func TestJudgeSolutionCpp(t *testing.T) {
	f := problem.DiskFetcher{ParentDir: "../data"}
	handler := RequestHandler{
		Fetcher: &f,
	}

	twoSumIncorrect := `
	class Solution {
		public:
			vector<int> twoSum(vector<int>& nums, int target) {
				vector<int> lol = {1,2,3};
				return lol;
			}
	};
	`

	twoSumCorrect := `
	#include <map>
	class Solution {
		public:
			vector<int> twoSum(vector<int>& nums, int target) {
				  map<int, int> map;
				  vector<int> pairs;
				  for(int i = 0; i < nums.size(); i++) {
					  int complement = target - nums[i];
					  if(map.find(complement) != map.end()) {
						  pairs.push_back(map.find(complement)->second);
						  pairs.push_back(i);
						  break;
					  }
					  map.insert(pair<int, int>(nums[i], i));
				  }
				  return pairs;
			};
	};
	`
	t.Run("judges incorrect c++ solution for 2sum correctly formatted", func(t *testing.T) {
		out, err := handler.JudgeSolution(JudgeRequest{
			Language: codegen.Cpp,
			Code:     twoSumIncorrect,
			Problem:  "two_sum",
			TestCases: []*problem.TestCase{
				{
					Input: `

					[1,2,3,4,5]
					5
				`,
					ExpectedOutput: []string{`[2, 1]`},
				},
			},
		})
		require.NoError(t, err)
		require.NotNil(t, out)
		require.False(t, out.Results[0].Passed)
		require.Equal(t, "[1,2,3]", out.Results[0].Stdout)
	})

	t.Run("judges correct c++ solution for 2sum correctly formatted", func(t *testing.T) {
		testCases := []*problem.TestCase{
			{
				Input: `[1,2,3,4,5]
					5
				`,
				ExpectedOutput: []string{`[1,2]`},
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
				ExpectedOutput: []string{`[1,3]`, `[3,1]`},
			},
		}

		out, err := handler.JudgeSolution(JudgeRequest{
			Language:  codegen.Cpp,
			Code:      twoSumCorrect,
			Problem:   "two_sum",
			TestCases: testCases,
		})
		require.NoError(t, err)
		require.NotNil(t, out)
		for _, res := range out.Results {
			require.True(t, res.Passed)
		}
	})
}
