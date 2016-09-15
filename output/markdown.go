package output

import (
	"errors"
	"fmt"
	"os"

	"github.com/jutkko/mindown/util"
)

const DEPTH_LIMIT int = 6

func WriteMarkdown(filename string, graph *util.Graph) error {
	if graph == nil {
		return errors.New("Graph is nil")
	}

	for _, node := range graph.GetNodes() {
		err := writeMarkdownRecursively(1, filename, node)

		if err != nil {
			return errors.New(fmt.Sprintf("Failed to write file: %s %s", filename, err.Error()))
		}
	}

	return nil
}

func writeMarkdownRecursively(depth int, filename string, node *util.Node) error {
	err := appendToFile(fmt.Sprintf("%s%s\n", getHash(depth), node.GetTitle()), filename)

	if err != nil || len(node.GetChildren()) == 0 {
		return err
	}

	for _, node := range node.GetChildren() {
		err := writeMarkdownRecursively(depth+1, filename, node)
		if err != nil {
			return err
		}
	}

	return nil
}

func getHash(level int) (result string) {
	result = ""

	if level <= DEPTH_LIMIT {
		for i := 0; i < level; i++ {
			result += "#"
		}
		result += " "
	}

	return
}

func appendToFile(data, filename string) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()
	if _, err := f.WriteString(data); err != nil {
		return err
	}

	return nil
}
