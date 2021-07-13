package iterator

import "testing"

func TestIterator(t *testing.T) {
	shelf := &BookShelf{}

	books := []string{"A", "B", "C", "D", "E", "F", "G"}

	for _, b := range books {
		shelf.Add(&Book{Name: b})
	}

	for iterator := shelf.Iterator(); iterator.HasNext(); iterator.Next() {
		index := iterator.Index()
		value := iterator.Value().(*Book)

		if value.Name != books[index] {
			t.Errorf("Expected book name %s, got %s", books[index], value.Name)
		}
	}
}
