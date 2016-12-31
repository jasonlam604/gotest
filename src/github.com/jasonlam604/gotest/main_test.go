package main

import (
	"flag"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var systemTest *bool

func init() {
	systemTest = flag.Bool("systemTest", true, "Set to true when running system tests")
}

func TestSystem(t *testing.T) {
	if *systemTest {
		main()
	}
}

func TestReverseMe(t *testing.T) {
	assert.Equal(t, "olleh", reverseMe("hello"))
}

func TestReverseMeAgain(t *testing.T) {
	assert.Equal(t, "hello", reverseMe("olleh"))
}
