package input_test

import (
	"github.com/jutkko/mindown/input"
	"github.com/jutkko/mindown/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ompl", func() {
	It("should parse the file into a graph", func() {
		node := util.NewNode("Vim-notes", "")
		expectedGraph := &util.Graph{}
		expectedGraph.AddNode(node)

		actualGraph, err := input.ParseOpml("../testdata/simple.opml")
		Expect(err).NotTo(HaveOccurred())
		Expect(actualGraph).To(Equal(expectedGraph))
	})

	It("should parse the file into a graph", func() {
		node := util.NewNode("Vim-notes", "")
		node1 := util.NewNode("Vim-notes1", "")
		expectedGraph := &util.Graph{}
		expectedGraph.AddNode(node)
		expectedGraph.AddNode(node1)

		actualGraph, err := input.ParseOpml("../testdata/simple-two-root.opml")
		Expect(err).NotTo(HaveOccurred())
		Expect(actualGraph).To(Equal(expectedGraph))
	})

	It("should parse the file into a graph", func() {
		node := util.NewNode("Vim-notes", "")

		nodeIntro := util.NewNode("Intro", "")
		nodeToVim := util.NewNode("To vim", "")
		nodeIntro.AddChild(nodeToVim)
		nodeToReadManual := util.NewNode("To reading the manual", "")
		nodeUserMan := util.NewNode("User manual", "")
		nodeToReadManual.AddChild(nodeUserMan)
		nodeHowUseMan := util.NewNode("How to use the manuals", "")
		nodeNavigate := util.NewNode("Navigate", "")
		nodeHowUseMan.AddChild(nodeNavigate)
		nodeSearch := util.NewNode("Search", "")
		nodeHowUseMan.AddChild(nodeSearch)
		nodeToReadManual.AddChild(nodeHowUseMan)
		nodeIntro.AddChild(nodeToReadManual)

		nodeMotion := util.NewNode("Motion", "")
		nodeBasic := util.NewNode("Basic", "")
		nodeMotion.AddChild(nodeBasic)
		nodeRefManual := util.NewNode("Reference manual", "")
		nodeMotion.AddChild(nodeRefManual)

		node.AddChild(nodeIntro)
		node.AddChild(nodeMotion)

		expectedGraph := &util.Graph{}
		expectedGraph.AddNode(node)

		actualGraph, err := input.ParseOpml("../testdata/vim-notes-example-simple.opml")
		Expect(err).NotTo(HaveOccurred())
		Expect(actualGraph).To(Equal(expectedGraph))
	})

	Context("when the file does not exist", func() {
		It("should return an error", func() {
			_, err := input.ParseOpml("../testdata/whassap.opml")
			Expect(err).To(HaveOccurred())
		})
	})

	Context("when the file format is broken", func() {
		It("should return an error", func() {
			_, err := input.ParseOpml("../testdata/broken.opml")
			Expect(err).To(HaveOccurred())
		})
	})
})
