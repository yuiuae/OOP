package structural

import "fmt"

type Vacuum struct {
}

func (v *Vacuum) CleanInHouse() {
	fmt.Println("User vacuums the floor")
}
