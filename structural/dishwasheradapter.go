package structural

import "fmt"

type DishwasherAdapter struct {
	dishwasherDevice *Dishwasher
}

func (d *DishwasherAdapter) CleanInHouse() {
	fmt.Println("Adapter converts cleaning to washing")
	d.dishwasherDevice.WashUp()
}
