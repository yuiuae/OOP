package main

import (
	"github.com/yuiuae/OOP/behavioral"
	"github.com/yuiuae/OOP/structural"
)

// import "github.com/yuiuae/OOP/creational"

func main() {
	// behavioral example
	pool := &behavioral.Pool{}
	addCommand := &behavioral.AddWaterCommand{
		Tank: pool,
	}
	pourCommand := &behavioral.PoreAllWaterCommand{
		Tank: pool,
	}
	addLever := &behavioral.Lever{
		TankCommand: addCommand,
	}
	addLever.Push()

	pourLever := &behavioral.Lever{
		TankCommand: pourCommand,
	}
	pourLever.Push()

	// structural example
	// House cleaning
	// The user can use the vacuum cleaner to clean
	// With an adapter, he can use the dishwasher
	user := &structural.User{}
	vacuum := &structural.Vacuum{}
	user.FullCleanInHouse(vacuum)
	dishwasher := &structural.Dishwasher{}
	dishwasherDeviceAdapter := &structural.DishwasherAdapter{
		DishwasherDevice: dishwasher,
	}
	user.FullCleanInHouse(dishwasherDeviceAdapter)
}
