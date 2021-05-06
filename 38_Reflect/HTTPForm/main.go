package main

type state struct {
	field1 string  `tformname:"Field One" tformvalues:"a|b|c"`
	field2 int     `tformname:"Field Two" tisvalidrange:"0|10"`
	field3 float64 `tformname:"Field Three" tisvalidrange:"0.1|0.7"`
}

func main() {

}
