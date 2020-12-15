package main

// TAdd Type concentrating math add ops.
type TAdd struct{}

type mathadd interface {
	add(x, y int8) int8
}

func (t TAdd) add(x, y int8) int8 {
	return x + y
}

// TMath Type concentrating math ops.
type TMath struct{}

type math interface {
	mathadd
	multiply(x, y int8) int8
}

func (t TMath) add(x, y int8) int8 {
	return x + y
}

func (t TMath) multiply(x, y int8) int8 {
	return x * y
}

// TExtMath Type concentrating extended math ops.
type TExtMath struct {
	a TAdd
}

func (t TExtMath) add(x, y int8) int8 {
	return t.a.add(x, y)
}

func (t TExtMath) multiply(x, y int8) int8 {
	return x * y
}
