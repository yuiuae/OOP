package main

import (
	"fmt"

	"github.com/yuiuae/OOP/behavioral"
	"github.com/yuiuae/OOP/creational"
	"github.com/yuiuae/OOP/structural"
)

func main() {
	// behavioral example - Command
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

	// creational example - 	Factory Method
	sl, _ := creational.GetLocomotive("Steam locomotive")
	fmt.Println((sl))
	et, _ := creational.GetLocomotive("Electric train")
	fmt.Println(et)

	// structural example - Decorator
	basicKnowledge := &structural.BasicKnowledge{}
	humanKnowledge := &structural.HumanKnowledge{
		Education: basicKnowledge,
	}
	humanTechKnowledge := &structural.HumanTechKnowledge{
		Education: humanKnowledge,
	}
	fmt.Println("Total number of teaching hours for humanities and technical education is ", humanTechKnowledge.GetHours())

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
