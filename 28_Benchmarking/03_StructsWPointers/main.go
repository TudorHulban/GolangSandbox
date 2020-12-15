package main

type aStruct1 struct {
	state string
	f1    string
	f2    string
	f3    string
	f4    string
}

type aStruct2 struct {
	state *string
	f1    *string
	f2    *string
	f3    *string
	f4    *string
}

func aFunction1(p aStruct1) string {
	return p.state
}

// worst performance
func aFunction2(p aStruct2) string {
	return *p.state
}

func aFunction3(p *aStruct1) string {
	return p.state
}

func aFunction4(p *aStruct1) string {
	return (*p).state
}

// performance is worst than passing by value
func aFunction5(p *aStruct2) string {
	return *p.state
}

// same performance as passing the struct by value
func aFunction6(p string) string {
	return p
}

func main() {}
