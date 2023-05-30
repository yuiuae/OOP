package structural

import "fmt"

type DishwasherAdapter struct {
	DishwasherDevice *Dishwasher
}

func (d *DishwasherAdapter) CleanInHouse() {
	fmt.Println("Adapter converts cleaning to washing")
	d.DishwasherDevice.WashUp()
}
