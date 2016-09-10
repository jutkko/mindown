package input

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jutkko/mindown/util"
)

func ParseOmpl(f *os.File) *util.Graph {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return nil
}
