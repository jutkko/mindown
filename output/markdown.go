package output

import (
	"io/ioutil"

	"github.com/jutkko/mindown/util"
)

func WriteMarkdown(filename string, graph *util.Graph) error {
	err := ioutil.WriteFile(filename, []byte("# Vim-notes"), 0644)
	if err != nil {
		panic("Cannot write " + err.Error())
	}
	return nil
}
