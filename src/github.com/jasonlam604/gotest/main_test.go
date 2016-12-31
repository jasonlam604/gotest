package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"flag"
	"os"
)

func TestMain(m *testing.M) {
    flag.Parse()
    code := m.Run()
    os.Exit(code)
}

func TestReverseMe(t *testing.T) {
	assert.Equal(t, "olleh", reverseMe("hello"))
}

func TestReverseMeAgain(t *testing.T) {
	assert.Equal(t, "hello", reverseMe("olleh"))
}
