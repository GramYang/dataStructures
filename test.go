package main

import (
	"dataStructures/circlebuffer"
	"fmt"
)

func main() {
	buffer1:=circlebuffer.New(3)
	buffer1.Add(1,"a")
	buffer1.Add("b")
	fmt.Println(buffer1)
}