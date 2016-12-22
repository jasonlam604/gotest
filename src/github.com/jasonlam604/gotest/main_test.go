package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseMe(t *testing.T) {
	assert.Equal(t, "olleh", reverseMe("hello"))
}

func TestReverseMeAgain(t *testing.T) {
	assert.Equal(t, "hello", reverseMe("olleh"))
}
