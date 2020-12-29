package main

import (
	"fmt"
)

type process struct {
	state string
}

func NewProcess(initialValue string) *processBuilder {
	return &processBuilder{
		build: process{
			state: initialValue,
		},
	}
}

type processBuilder struct {
	build process
}

func (b *processBuilder) StepAdd(value string) *processBuilder {
	b.build.state = b.build.state + value
	return b
}

func main() {
	p1 := NewProcess("xxxxx").StepAdd("y").StepAdd("z")
	fmt.Println((*p1).build.state) // result: xxxxxyz
}
