package abstract_factory

type AbstractFactory interface {
	CreateWater(volume float64) AbstractWater
	CreateBottle(volume float64) AbstractBottle
}

type AbstractWater interface {
	GetVolume() float64
}

type AbstractBottle interface {
	PourWater(water AbstractWater) // Bottle with water.
	GetBottleVolume() float64
	GetWaterVolume() float64
}
