package composite

type Component interface {
	Add(child Component)
	Name() string
	Children() []Component
	Print(prefix string) string
}
