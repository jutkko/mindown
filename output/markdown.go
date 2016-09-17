package output

import (
	"errors"
	"fmt"
	"os"

	"github.com/jutkko/mindown/util"
)

const DEPTH_LIMIT int = 6

func WriteMarkdown(filename string, forceWrite bool, graph *util.Graph) error {
	if !forceWrite {
		if _, err := os.Stat(filename); !os.IsNotExist(err) {
			return errors.New("File exists")
		}
	}

	// Try to remove the existing file
	os.Remove(filename)

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return errors.New("Failed to open file")
	}
	defer file.Close()

	if graph == nil {
		return errors.New("Graph is nil")
	}

	for _, node := range graph.GetNodes() {
		err := writeMarkdownRecursively(1, file, node)

		if err != nil {
			return errors.New(fmt.Sprintf("Failed to write file: %s %s", filename, err.Error()))
		}
	}

	return nil
}

func writeMarkdownRecursively(depth int, file *os.File, node *util.Node) error {
	err := appendToFile(fmt.Sprintf("%s%s\n", getHash(depth), node.GetTitle()), file)

	if err != nil || len(node.GetChildren()) == 0 {
		return err
	}

	for _, node := range node.GetChildren() {
		err := writeMarkdownRecursively(depth+1, file, node)
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

func appendToFile(data string, file *os.File) error {
	if _, err := file.WriteString(data); err != nil {
		return err
	}

	return nil
}
