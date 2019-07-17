package kata

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestContainerOfWater(t *testing.T) {
	c := ContainerOfWater{}

	assert.Equal(t, 1, c.maxArea([]int{1,1}))
	assert.Equal(t, 2, c.maxArea([]int{1,2,2}))
	assert.Equal(t, 6, c.maxArea([]int{1,3,2,3}))
	assert.Equal(t, 9, c.maxArea([]int{1,3,2,2,3}))
	assert.Equal(t, 49, c.maxArea([]int{1,8,6,2,5,4,8,3,7}))
}
