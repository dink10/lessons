package flyweight

import "fmt"

type Flyweighter interface {
	Draw(weight, height int, opacity float64) string
}

type FlyweightFactory struct {
	pool map[string]Flyweighter
}

func (ff *FlyweightFactory) getFlyweight(filename string) Flyweighter {
	if ff.pool == nil {
		ff.pool = make(map[string]Flyweighter)
	}

	if _, ok := ff.pool[filename]; !ok {
		ff.pool[filename] = &ConcreteFlyweight{filename: filename}
	}

	return ff.pool[filename]
}

type ConcreteFlyweight struct {
	filename string
}

func (f *ConcreteFlyweight) Draw(weight, height int, opacity float64) string {
	return fmt.Sprintf("Draw image: %s, width: %d, height: %d, opacity: %.2f", f.filename, weight, height, opacity)
}
