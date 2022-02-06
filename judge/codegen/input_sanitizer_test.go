package codegen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInputRegex(t *testing.T) {
	assert.True(t, isNumberList("[0,1,2]"))
	assert.True(t, isNumberList("[0, 1,2]"))
	assert.True(t, isNumberList("[ 0, 1,2]"))
	assert.True(t, isNumberList("[ 0, 1, 2]"))
	assert.True(t, isNumberList("[ 0, 1, 2 ]"))
	assert.True(t, isNumberList("[ 0, 1.5, 2 ]"))
	assert.True(t, isNumberList("[ 1.5, 1.5, 2.3 ]"))
	assert.False(t, isNumberList("[ 0, 1, 2, ]"))
}
