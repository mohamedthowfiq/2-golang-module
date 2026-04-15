package main

import (
	"go-modules/internal/greet"
	"fmt"
)

func main(){

	msg1 := greet.Hello("thowfiq")

	fmt.Println(msg1)
}