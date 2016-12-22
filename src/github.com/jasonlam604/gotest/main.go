package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
)

func main() {
	fmt.Println(reverseMe("hello"))
}

func reverseMe(name string) string {
	return stringutil.Reverse(name)
}
