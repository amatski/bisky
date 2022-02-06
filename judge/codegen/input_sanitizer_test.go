package codegen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumberList(t *testing.T) {
	assert.NotNil(t, isNumberList("[0,1,2]"))
	assert.Equal(t, "0,1,2", isNumberList("[0,1,2]").Elements)
	assert.NotNil(t, isNumberList("[0, 1,2]"))
	assert.NotNil(t, isNumberList("[ 0, 1,2]"))
	assert.NotNil(t, isNumberList("[ 0, 1, 2]"))
	assert.NotNil(t, isNumberList("[ 0, 1, 2 ]"))
	assert.NotNil(t, isNumberList("[ 0, 1.5, 2 ]"))
	assert.Equal(t, "0", isNumberList("[0]").Elements)
	assert.Equal(t, "1.5", isNumberList("[1.5]").Elements)
	assert.Equal(t, "1.5,  2 ", isNumberList("[1.5,  2 ]").Elements)
	assert.NotNil(t, isNumberList("[ 1.5, 1.5, 2.3 ]"))
	assert.Nil(t, isNumberList("[ 0, 1, 2, ]"))
	assert.NotNil(t, isNumberList("[]"))
	assert.Equal(t, " ", isNumberList("[ ]").Elements)
	assert.Nil(t, isNumberList("[ 0, 1, 2, ]"))
	assert.Nil(t, isNumberList("[ 0, -1, 2, ]"))
	assert.Equal(t, " 0, -1, 2", isNumberList("[ 0, -1, 2]").Elements)
	assert.NotNil(t, isNumberList("[ 0, -1, 2 ]"))
	assert.Nil(t, isNumberList("[ 0, -1, +2 ]"))
	assert.Equal(t, " 0, -1, 2 ", isNumberList("[ 0, -1, 2 ]").Elements)
	assert.Equal(t, "-1", isNumberList("[-1]").Elements)
	assert.Equal(t, "-1,-1", isNumberList("[-1,-1]").Elements)
	assert.Nil(t, isNumberList("[ 0, -1, - 2 ]"))
	assert.Nil(t, isNumberList("[ 0, -1, 2., 3 ]"))
}
