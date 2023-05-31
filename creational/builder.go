package creational

type IBuilder interface {
	setCarBodyType()
	setEngineType()
	seеTireType()
	getCar() Car
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "passenger" {
		return newPassengerBuilder()
	}
	if builderType == "sport" {
		return newSportBuilder()
	}
	return nil
}

type PassengerBuilder struct {
	CarBodyType string
	EngineType  string
	TireType    string
}

func newPassengerBuilder() *PassengerBuilder {
	return &PassengerBuilder{}
}
func (b *PassengerBuilder) setCarBodyType() {
	b.CarBodyType = "Passenger"
}
func (b *PassengerBuilder) setEngineType() {
	b.EngineType = "Standart"
}
func (b *PassengerBuilder) seеTireType() {
	b.TireType = "Standart"
}
func (b *PassengerBuilder) getCar() Car {
	return Car{
		CarBodyType: b.CarBodyType,
		EngineType:  b.EngineType,
		TireType:    b.TireType,
	}
}

type SportBuilder struct {
	CarBodyType string
	EngineType  string
	TireType    string
}

func newSportBuilder() *SportBuilder {
	return &SportBuilder{}
}
func (b *SportBuilder) setCarBodyType() {
	b.CarBodyType = "Sport"
}
func (b *SportBuilder) setEngineType() {
	b.EngineType = "Sport"
}
func (b *SportBuilder) seеTireType() {
	b.TireType = "Low-profile"
}
func (b *SportBuilder) getCar() Car {
	return Car{
		CarBodyType: b.CarBodyType,
		EngineType:  b.EngineType,
		TireType:    b.TireType,
	}
}

type Car struct {
	CarBodyType string
	EngineType  string
	TireType    string
}

type Director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) BuildCar() Car {
	d.builder.setCarBodyType()
	d.builder.setEngineType()
	d.builder.seеTireType()
	return d.builder.getCar()
}
