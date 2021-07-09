package abstract_factory

type CocaColaWater struct {
	volume float64
}

func (w *CocaColaWater) GetVolume() float64 {
	return w.volume
}

type CocaColaBottle struct {
	water  AbstractWater
	volume float64
}

func (b *CocaColaBottle) PourWater(water AbstractWater) {
	b.water = water
}

func (b *CocaColaBottle) GetBottleVolume() float64 {
	return b.volume
}

func (b *CocaColaBottle) GetWaterVolume() float64 {
	return b.water.GetVolume()
}
