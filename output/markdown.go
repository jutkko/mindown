package output

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/jutkko/mindown/util"
)

func WriteMarkdown(filename string, graph *util.Graph) error {
	err := ioutil.WriteFile(filename, []byte(fmt.Sprintf("%s\n", graph.GetNodes()[0].GetTitle())), 0644)
	err := ioutil.WriteFile(filename, []byte(fmt.Sprintf("# %s\n", graph.GetNodes()[0].GetTitle())), 0644)

	if err != nil {
		return errors.New(fmt.Sprintf("Failed to write file: %s %s", filename, err.Error()))
	}

	return nil
}
