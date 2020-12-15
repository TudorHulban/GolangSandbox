package main

// test with - go test -bench=.

import (
	"testing"
)

var numberTasks = 100000
var fieldName = "Name"

func benchmarkFilterDirect(pNoTasks int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		work := *NewTasks(pNoTasks)
		showFieldDirect(fieldName, &work)
	}
}

func BenchmarkDirect(b *testing.B) {
	benchmarkFilterDirect(numberTasks, b)
}

func benchmarkFilterAssertion(pNoTasks int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		work := *NewTasks(pNoTasks)
		showFieldDirect(fieldName, &work)
	}
}

func BenchmarkAssert(b *testing.B) {
	benchmarkFilterAssertion(numberTasks, b)
}

func benchmarkFilterName(pNoTasks int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		work := *NewTasks(pNoTasks)
		showFieldName("Name", &work)
	}
}

func BenchmarkName(b *testing.B) {
	benchmarkFilterName(numberTasks, b)
}

// --------------- Test Results

/*
var numberTasks = 10000
var fieldName = "ID"

goos: linux
goarch: amd64
BenchmarkDirect-4   	     100	  14030771 ns/op
BenchmarkAssert-4   	     100	  14082028 ns/op
*/

/*
var numberTasks = 100000
var fieldName = "ID"

goos: linux
goarch: amd64
BenchmarkDirect-4   	       5	 207640264 ns/op
BenchmarkAssert-4   	      10	 203703491 ns/op
*/

/*
var numberTasks = 100000
var fieldName = "Name"

goos: linux
goarch: amd64
BenchmarkDirect-4   	       5	 215689393 ns/op
BenchmarkAssert-4   	       5	 202420484 ns/op
BenchmarkName-4     	      10	 158550310 ns/op
*/
