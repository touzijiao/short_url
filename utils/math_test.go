package utils

import (
	"github.com/stretchr/testify/assert" //一个断言包
	"testing"
)

func TestInit(t *testing.T) {

	assert.Equal(t, 62, Length, "Tonken Length error (断言错误信息)")

}

func TestIdToString(t *testing.T) {
	id := 72

	expectValue := "1a"
	assert.Equal(t, expectValue, IdToString(id))
}

func TestStringToId(t *testing.T) {
	str := "1a"
	expectValue := 72
	assert.Equal(t, expectValue, StringToId(str))
}
