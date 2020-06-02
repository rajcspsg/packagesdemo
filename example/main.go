package main

import (
	"fmt"

	"github.com/rajcspsg/packagesdemo/greet"
)

func main() {
	greet.Hello()
	fmt.Println(greet.Shark)
	oct := greet.Octopus { Name : "Jesse", Color: "Orange" }
	fmt.Println(oct.String())
	oct.Reset()
	fmt.Println(oct.String())
}
