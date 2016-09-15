package output_test

import (
	"io/ioutil"
	"os"

	"github.com/jutkko/mindown/input"
	"github.com/jutkko/mindown/output"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Markdown", func() {
	var (
		tempFile *os.File
		err      error
		filename string
		content  []byte
	)

	BeforeEach(func() {
		tempFile, err = ioutil.TempFile("", "")
		Expect(err).NotTo(HaveOccurred())
		filename = tempFile.Name()
		content = make([]byte, 50)
	})

	It("should output the correct file given simple graph", func() {
		graph, err := input.ParseOpml("../testdata/simple.opml")
		Expect(err).NotTo(HaveOccurred())

		output.WriteMarkdown(filename, graph)
		tempFile.Read(content)
		Expect(content).Should(ContainSubstring("# Vim-notes\n"))

		Expect(os.Remove(filename)).To(Succeed())
	})

	It("should output the correct file given simple-two-root graph", func() {
		graph, err := input.ParseOpml("../testdata/simple-two-root.opml")
		Expect(err).NotTo(HaveOccurred())

		output.WriteMarkdown(filename, graph)
		tempFile.Read(content)
		Expect(content).Should(ContainSubstring("# Vim-notes\n"))
		Expect(content).Should(ContainSubstring("# Vim-notes1\n"))

		Expect(os.Remove(filename)).To(Succeed())
	})

	Context("when the file cannot be written to", func() {
		It("should return an error", func() {
			err := output.WriteMarkdown("-/some-random-file-name", nil)
			Expect(err).To(HaveOccurred())
		})
	})
})
