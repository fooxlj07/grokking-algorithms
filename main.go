package main

import (
	chapter3 "grokking-algorithms/chapter3/go"
)

func main() {
	box := chapter3.BuildBox()
	chapter3.FindKey(box)
}
