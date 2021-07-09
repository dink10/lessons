package abstract_factory

type PepsiFactory struct {
}

func NewPepsiFactory() AbstractFactory {
	return &PepsiFactory{}
}

func (f *PepsiFactory) CreateWater(volume float64) AbstractWater {
	return &PepsiWater{volume: volume}
}

func (f *PepsiFactory) CreateBottle(volume float64) AbstractBottle {
	return &PepsiBottle{
		volume: volume,
	}
}
