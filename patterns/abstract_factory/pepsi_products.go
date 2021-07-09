package abstract_factory

type PepsiWater struct {
	volume float64
}

func (w *PepsiWater) GetVolume() float64 {
	return w.volume
}

type PepsiBottle struct {
	water  AbstractWater
	volume float64
}

func (b *PepsiBottle) PourWater(water AbstractWater) {
	b.water = water
}

func (b *PepsiBottle) GetBottleVolume() float64 {
	return b.volume
}

func (b *PepsiBottle) GetWaterVolume() float64 {
	return b.water.GetVolume()
}
