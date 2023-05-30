package main

// import "github.com/yuiuae/OOP/behavioral"
// import "github.com/yuiuae/OOP/creational"
import "github.com/yuiuae/OOP/structural"

func main() {
	user := &structural.User{}
	vacuum := &structural.Vacuum{}
	user.FullCleanInHouse(vacuum)
	dishwasher := &structural.Dishwasher{}
	dishwasherDeviceAdapter := &structural.DishwasherAdapter{
		dishwasherDevice: dishwasher,
	}
	user.FullCleanInHouse(dishwasherDeviceAdapter)
}
