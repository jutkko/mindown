package util

import "os"

type Graph struct {
}

func (g *Graph) Export() error {
	return nil
}

type Parser interface {
	Parse(file *os.File) *Graph
}
