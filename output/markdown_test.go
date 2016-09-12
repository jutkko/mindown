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
		Expect(content).Should(ContainSubstring("# Vim-notes"))
	})

})
