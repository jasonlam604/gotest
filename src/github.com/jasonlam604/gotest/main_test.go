package main

import (
	"flag"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
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
