package abstract_factory

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	cocaColaFactory := NewCocaColaFactory()

	colaWater := cocaColaFactory.CreateWater(1.5)
	colaBottle := cocaColaFactory.CreateBottle(1.5)

	colaBottle.PourWater(colaWater)

	if colaBottle.GetWaterVolume() != colaBottle.GetBottleVolume() {
		t.Errorf("Expected volume %.1f, but %.1f", colaBottle.GetWaterVolume(), colaBottle.GetBottleVolume())
	}
}
