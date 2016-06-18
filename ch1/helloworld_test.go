package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	assert.Equal(t, HelloWord(), "Hello, 世界", "they should be equal")
}
