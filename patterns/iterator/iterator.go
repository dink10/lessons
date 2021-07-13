package iterator

type Iterator interface {
	Index() int
	Value() interface{}
	HasNext() bool
	Next()
	Prev()
	Reset()
	End()
}

type Aggregator interface {
	Iterator() Iterator
}
