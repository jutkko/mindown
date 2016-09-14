package output_test

import (
	"fmt"
	"io/ioutil"

	"github.com/jutkko/mindown/input"
	"github.com/jutkko/mindown/output"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Markdown", func() {
	It("should output the correct file given simple graph", func() {
		graph, err := input.ParseOpml("../testdata/simple.opml")
		Expect(err).NotTo(HaveOccurred())

		tempFile, err := ioutil.TempFile("", "")
		Expect(err).NotTo(HaveOccurred())

		fmt.Printf("%s\n", tempFile.Name())
		output.WriteMarkdown(tempFile.Name(), graph)
		content := make([]byte, 20)
		tempFile.Read(content)
		Expect(content).Should(ContainSubstring("# Vim-notes\n"))
	})

	It("should output the correct file given simple-two-root graph", func() {
		graph, err := input.ParseOpml("../testdata/simple-two-root.opml")
		Expect(err).NotTo(HaveOccurred())

		tempFile, err := ioutil.TempFile("", "")
		Expect(err).NotTo(HaveOccurred())

		fmt.Printf("%s\n", tempFile.Name())
		output.WriteMarkdown(tempFile.Name(), graph)
		content := make([]byte, 20)
		tempFile.Read(content)
		Expect(content).Should(ContainSubstring("# Vim-notes\n"))
		Expect(content).Should(ContainSubstring("# Vim-notes1\n"))
	})

	Context("when the file cannot be written to", func() {
		It("should return an error", func() {
			err := output.WriteMarkdown("-/some-random-file-name", nil)
			Expect(err).To(HaveOccurred())
		})
	})
})
