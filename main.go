package main

import (
	"fmt"

	"github.com/mattfarina/go-test"

	"github.com/Masterminds/vcs"
)

func main() {
	test.Foo()

	n := &vcs.CommitInfo{
		Message: "foo",
	}

	fmt.Println(n.Message)
}
