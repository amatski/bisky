package problem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiskFetch(t *testing.T) {
	d := DiskFetcher{ParentDir: "../../data"}
	problemId := "two_sum"
	ext := ".py"
	codeTemplate, err := d.CodeTemplate(problemId, ext)
	assert.NoError(t, err)
	expected := "class Solution:\n    def twoSum(self, nums: List[int], target: int) -> List[int]:\n        "
	assert.Equal(t, expected, codeTemplate)
}
