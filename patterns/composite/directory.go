package composite

type Directory struct {
	name     string
	children []Component
}

func (d *Directory) Add(child Component) {
	d.children = append(d.children, child)
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) Children() []Component {
	return d.children
}

func (d *Directory) Print(prefix string) string {
	result := prefix + "/" + d.Name() + "\n"
	for _, v := range d.Children() {
		result += v.Print(prefix + "/" + d.Name())
	}

	return result
}

func NewDirectory(name string) *Directory {
	return &Directory{name: name}
}
