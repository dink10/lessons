package facade

import "strings"

// Man == facade
type Man struct {
	house *House
	tree  *Tree
	child *Child
}

func NewMan() *Man {
	return &Man{
		house: &House{},
		tree:  &Tree{},
		child: &Child{},
	}
}

func (m *Man) Todo() string {
	result := []string{
		m.house.Build(),
		m.tree.Grow(),
		m.child.Born(),
	}

	return strings.Join(result, ":")
}

type House struct {
}

func (h *House) Build() string {
	return "House is building"
}

type Tree struct {
}

func (h *Tree) Grow() string {
	return "Tree is growing"
}

type Child struct {
}

func (c *Child) Born() string {
	return "Child born"
}
