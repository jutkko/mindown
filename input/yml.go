package input

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"

	"github.com/jutkko/mindown/util"
)

type Instance struct {
	Name      string
	Instances int32
}

type Manifest struct {
	Name            string
	Instance_Groups []Instance
}

func ParseYaml(filename string) (*util.Graph, error) {
	source, _ := ioutil.ReadFile(filename)

	m := Manifest{}

	err := yaml.Unmarshal(source, &m)
	if err != nil {
		return nil, err
	}

	result := &util.Graph{}

	node := util.NewNode(m.Name, "")
	result.AddNode(node)

	for _, instance := range m.Instance_Groups {
		name := instance.Name
		if instance.Instances > 1 {
			name += fmt.Sprintf("*%d", instance.Instances)
		}

		instanceNode := util.NewNode(name, "")
		node.AddChild(instanceNode)
	}

	return result, nil
}
