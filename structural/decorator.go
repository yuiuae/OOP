package structural

type IEducation interface {
	GetHours() int
}

type BasicKnowledge struct{}

func (k *BasicKnowledge) GetHours() int {
	return 200
}

type HumanKnowledge struct {
	Education IEducation
}

func (k *HumanKnowledge) GetHours() int {
	return k.Education.GetHours() + 70
}

type HumanTechKnowledge struct {
	Education IEducation
}

func (k *HumanTechKnowledge) GetHours() int {
	return k.Education.GetHours() + 40
}
