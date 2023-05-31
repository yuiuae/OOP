package creational

import "fmt"

type ILocomotive interface {
	setClass(class string)
	setFuelTyoe(fuelType string)
	GetClass() string
	GetFuelType() string
}
type Locomotive struct {
	class    string
	fuelType string
}

func (l *Locomotive) setClass(class string) {
	l.class = class
}
func (l *Locomotive) setFuelTyoe(fuel string) {
	l.fuelType = fuel
}
func (l *Locomotive) GetClass() string {
	return l.class
}
func (l *Locomotive) GetFuelType() string {
	return l.fuelType
}

type SteamLocomotive struct {
	Locomotive
}

func newSteamLocomotive() ILocomotive {
	return &SteamLocomotive{
		Locomotive{
			class:    "Steam locomotive",
			fuelType: "coal",
		},
	}
}

type ElectricTrain struct {
	Locomotive
}

func newElectricTrain() ILocomotive {
	return &ElectricTrain{
		Locomotive{
			class:    "Electric train",
			fuelType: "electricity",
		},
	}
}

func GetLocomotive(class string) (ILocomotive, error) {
	if class == "Steam locomotive" {
		return newSteamLocomotive(), nil
	}
	if class == "Electric train" {
		return newElectricTrain(), nil
	}
	return nil, fmt.Errorf("Wrong Locomotive class")
}
