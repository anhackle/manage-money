package basic

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 2
	// )

	// actual := AddOne(input)
	// if actual != output {
	// 	t.Errorf("Error")
	// }

	assert.Equal(t, AddOne(2), 3, "AddOne(2) should be equal 3")
}

func TestRequire(t *testing.T) {
	require.Equal(t, 2, 3)
	fmt.Println("Testing require testify")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 2, 3)
	fmt.Println("Testing assert testify")
}
