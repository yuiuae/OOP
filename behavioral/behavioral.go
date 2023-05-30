package behavioral

import "fmt"

type Lever struct {
	command Command
}

func (l *Lever) push() {
	l.command.execute()
}

type Command interface {
	execute()
}

type AddWaterCommand struct {
	tank Tank
}

func (c *AddWaterCommand) execute() {
	c.tank.add()
}

type PoreAllWaterCommand struct {
	tank Tank
}

func (c *PoreAllWaterCommand) execute() {
	c.tank.pour()
}

type Tank interface {
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
