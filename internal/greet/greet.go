package greet

import "strings"

// this is exported functions start with Captial letter
// other  packages can call it -> p1,p2
func Hello(name string)string{

	clean:= normalizeName(name)

	return "Hello " + clean
}



func normalizeName(name string)string{

	n:=strings.TrimSpace(name)

	if n ==""{
		return "Guest"
	}

	return strings.ToUpper(name)
}