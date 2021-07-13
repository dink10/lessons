package iterator

type Book struct {
	Name string
}

type BookShelf struct {
	Books []*Book
}

func (bs *BookShelf) Iterator() Iterator {
	return &BookIterator{
		shelf: bs,
	}
}

func (bs *BookShelf) Add(book *Book) {
	bs.Books = append(bs.Books, book)
}
