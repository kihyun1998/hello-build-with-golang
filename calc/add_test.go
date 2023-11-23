package calc

import (
	"hello-build-with-golang/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	num1 := util.RandomInt(1, 1000)
	num2 := util.RandomInt(1, 1000)

	result := Add(num1, num2)
	require.Equal(t, result, num1+num2)
}
