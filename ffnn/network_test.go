package ffnn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRelu(t *testing.T) {
	assert.Equal(t, 3.0, ReLu(3))
}
