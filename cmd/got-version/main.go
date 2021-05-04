package main

import (
	"fmt"

	"github.com/roger-russel/got-version/internal/reader"
	"github.com/roger-russel/got-version/internal/tag"
)

func main() {
	pipe, errPipe := reader.Pipe()

	if errPipe != nil {
		fmt.Println(errPipe)
		return
	}

	last := tag.GetLastTag(pipe)

	fmt.Println(last)
}
