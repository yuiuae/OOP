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
		tank: pool,
	}
	pourCommand := &behavioral.PoreAllWaterCommand{
		tank: pool,
	}
	addLever := &behavioral.Lever{
		command: behavioral.AddWaterCommand,
	}
	addLever.push()

	pourLever := &behavioral.Lever{
		command: behavioral.PoreAllWaterCommand,
	}
	pourLever.push()

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
