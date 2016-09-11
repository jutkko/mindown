package input

import (
	"fmt"

	"github.com/gilliek/go-opml/opml"
	"github.com/jutkko/mindown/util"
)

func ParseOpml(filename string) (*util.Graph, error) {
	doc, _ := opml.NewOPMLFromFile(filename)
	xml, _ := doc.XML()
	fmt.Printf(xml)

	return nil, nil
}
