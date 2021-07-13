package iterator

type BookIterator struct {
	shelf    *BookShelf
	index    int
	internal int
}

func (b *BookIterator) Index() int {
	return b.index
}

func (b *BookIterator) Value() interface{} {
	return b.shelf.Books[b.index]
}

func (b *BookIterator) HasNext() bool {
	if b.internal < 0 || b.internal >= len(b.shelf.Books) {
		return false
	}

	return true
}

func (b *BookIterator) Next() {
	b.internal++
	if b.HasNext() {
		b.index++
	}
}

func (b *BookIterator) Prev() {
	b.internal--
	if b.HasNext() {
		b.index--
	}
}

func (b *BookIterator) Reset() {
	b.index = 0
	b.internal = 0
}

func (b *BookIterator) End() {
	b.internal = len(b.shelf.Books) - 1
	b.internal = b.index
}
