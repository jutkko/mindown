package input

import (
	"github.com/gilliek/go-opml/opml"
	"github.com/jutkko/mindown/util"
)

func ParseOpml(filename string) (*util.Graph, error) {
	doc, err := opml.NewOPMLFromFile(filename)
	if err != nil {
		return nil, err
	}

	result := &util.Graph{}
	for _, outline := range doc.Body.Outlines {
		result.AddNode(recursivelyConvertToNodes(outline))
	}
	return result, nil
}

// This algorithm does not take care of circular graphs, it will cause
// stackoverflow
func recursivelyConvertToNodes(o opml.Outline) *util.Node {
	result := util.NewNode(o.Text, "")
	if len(o.Outlines) < 1 {
		return result
	}

	for _, outline := range o.Outlines {
		result.AddChild(recursivelyConvertToNodes(outline))
	}
	return result
}
