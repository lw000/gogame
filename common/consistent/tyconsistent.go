package ggconsistent

import "errors"

type Consistent struct {
	numOfVirtualNode int
	hashSortedNodes  []uint32
	circle           map[uint32]string
	nodes            map[string]bool
}

func New() *Consistent {
	return &Consistent{
		numOfVirtualNode: 20,
		circle:           make(map[uint32]string),
		nodes:            make(map[string]bool),
	}
}

func (c *Consistent) Get() (string, error) {
	if len(c.nodes) == 0 {
		return "", errors.New("no host added")
	}

	return "", nil
}

func (c *Consistent) Add(node string) error {

	return nil
}

func (c *Consistent) Remove(node string) error {

	return nil
}
