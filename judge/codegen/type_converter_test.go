package codegen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnswerConverter(t *testing.T) {
	ans := ConvertType(Integers, "[1,2,3]")
	assert.Equal(t, "[1,2,3]", ans)

	ans = ConvertType(Integers, "[1]")
	assert.Equal(t, "[1]", ans)

	ans = ConvertType(Integers, "[1.5,       2]")
	assert.Equal(t, "[1.5,2]", ans)
	ans = ConvertType(Integers, "[1.5       2]")
	assert.Equal(t, "[1.5,2]", ans)
	ans = ConvertType(Integers, "[1.5       2 -5]")
	assert.Equal(t, "[1.5,2,-5]", ans)
}
