package bridge

type Carer interface {
	Race() string
}

type Enginer interface {
	GetSound() string
}

type Car struct {
	engine Enginer
}

func NewCar(enginer Enginer) Carer {
	return &Car{engine: enginer}
}

func (c *Car) Race() string {
	return c.engine.GetSound()
}

type EngineSuzuki struct {
}

func (e *EngineSuzuki) GetSound() string {
	return "SUZUUUUUKI"
}

type EngineHonda struct {
}

func (e *EngineHonda) GetSound() string {
	return "HOOOOOOONDA"
}

type EngineBMW struct {
}

func (e *EngineBMW) GetSound() string {
	return "BMW"
}
