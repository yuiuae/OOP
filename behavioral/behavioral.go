package behavioral

import "fmt"

type Lever struct {
	TankCommand Command
}

func (l *Lever) Push() {
	l.TankCommand.execute()
}

type Command interface {
	execute()
}

type AddWaterCommand struct {
	Tank WaterTank
}

func (c *AddWaterCommand) execute() {
	c.Tank.add()
}

type PoreAllWaterCommand struct {
	Tank WaterTank
}

func (c *PoreAllWaterCommand) execute() {
	c.Tank.pour()
}

type WaterTank interface {
	add()
	pour()
}
type Pool struct {
	volume int
}

func (p *Pool) add() {
	p.volume = p.volume + 40
	fmt.Println("Add 40 liters of waters to the pool")
}

func (p *Pool) pour() {
	p.volume = 0
	fmt.Println("Pour all the water out of the pool")
}
