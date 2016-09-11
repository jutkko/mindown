package input

import (
	"fmt"

	"github.com/gilliek/go-opml/opml"
	"github.com/jutkko/mindown/util"
)

func ParseOpml(filename string) (*util.Graph, error) {
	doc, err := opml.NewOPMLFromFile(filename)
	if err != nil {
		return nil, err
	}

	xml, err := doc.XML()
	if err != nil {
		return nil, err
	}

	fmt.Printf(xml)
	return nil, nil
}
