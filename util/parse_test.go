package util_test

import (
	. "github.com/jutkko/mindown/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parse", func() {
	Describe("Node", func() {
		var (
			node1, node2, node3 *Node
		)

		BeforeEach(func() {
			node1 = NewNode("title1", "value1")
			node2 = NewNode("title2", "value2")
			node3 = NewNode("title3", "value3")
		})

		It("adds the correct child to existing node", func() {
			node1.AddChild(node2)
			node1Children := node1.GetChildren()
			Expect(node1Children).Should(ConsistOf(node2))
		})

		Context("adding multiple children", func() {
			It("adds the correct children to the existing node", func() {
				node1.AddChild(node2)
				node1.AddChild(node3)
				node1Children := node1.GetChildren()
				Expect(node1Children).Should(ConsistOf(node2, node3))
			})
		})
	})

	Describe("Graph", func() {
		var (
			node1 *Node
			graph *Graph
		)

		BeforeEach(func() {
			node1 = NewNode("title1", "value1")
			graph = &Graph{}
		})

		It("adds the correct child to existing node", func() {
			graph.AddNode(node1)
			Expect(graph.GetNodes()).Should(ConsistOf(node1))
		})
	})

})
