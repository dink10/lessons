package composite

type File struct {
	name string
}

func (f *File) Add(_ Component) {}

func (f *File) Name() string {
	return f.name
}

func (f *File) Children() []Component {
	return []Component{}
}

func (f *File) Print(prefix string) string {
	return prefix + "/" + f.Name() + "\n"
}

func NewFile(name string) *File {
	return &File{name: name}
}
