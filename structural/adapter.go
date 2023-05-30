package structural

import "fmt"

type User struct {
}

func (u *User) FullCleanInHouse(cln Cleaner) {
	fmt.Println("User cleans the whole house")
	cln.CleanInHouse()
}

type Cleaner interface {
	CleanInHouse()
}

type Vacuum struct {
}

func (v *Vacuum) CleanInHouse() {
	fmt.Println("User vacuums the floor")
}

type Dishwasher struct {
}

func (d *Dishwasher) WashUp() {
	fmt.Println("User washes dishes")
}

type DishwasherAdapter struct {
	DishwasherDevice *Dishwasher
}

func (d *DishwasherAdapter) CleanInHouse() {
	fmt.Println("Adapter converts cleaning to washing")
	d.DishwasherDevice.WashUp()
}
