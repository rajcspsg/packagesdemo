package greet

import "fmt"

type Octopus struct {
	Name string
	Color string
}

func (o Octopus) String() string {
	return fmt.Sprintf("Octpus Name is %s and color is %s", o.Name, o.Color)
}

func (o *Octopus) Reset() {
	o.Name = ""
	o.Color = ""
}
