package structural

import "fmt"

type Dishwasher struct {
}

func (d *Dishwasher) WashUp() {
	fmt.Println("User washes dishes")
}
