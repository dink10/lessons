package abstract_factory

type CocaColaFactory struct {
	a, b, c int
}

func NewCocaColaFactory() AbstractFactory {
	return &CocaColaFactory{}
}

func (f *CocaColaFactory) CreateWater(volume float64) AbstractWater {
	return &CocaColaWater{volume: volume}
}

func (f *CocaColaFactory) CreateBottle(volume float64) AbstractBottle {
	return &CocaColaBottle{
		volume: volume,
	}
}
