package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
	"os"
)

func main() {
	fmt.Println(reverseMe("hello"))
	os.Exit(1)
}

func reverseMe(name string) string {
	return stringutil.Reverse(name)
}
