package main

import (
	"fmt"
)

type process struct {
	initialValue string
	finalValue   string
}

type processBuilder struct {
	build process
}

func (b *processBuilder) StepAdd(pValue string) *processBuilder {
	b.build.finalValue = b.build.finalValue + pValue
	return b
}

func NewProcess(pInitialValue string) *processBuilder {
	return &processBuilder{build: process{finalValue: pInitialValue}}
}

func main() {
	p1 := NewProcess("xxxxx").StepAdd("y").StepAdd("z")
	fmt.Println((*p1).build.finalValue) // result: xxxxxyz
}
