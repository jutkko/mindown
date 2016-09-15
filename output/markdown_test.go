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
		content = make([]byte, 300)
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

	It("should output the correct file given a more complex graph", func() {
		graph, err := input.ParseOpml("../testdata/vim-notes-example-simple.opml")
		Expect(err).NotTo(HaveOccurred())

		output.WriteMarkdown(filename, graph)
		tempFile.Read(content)

		Expect(content).Should(ContainSubstring("# Vim-notes\n"))
		Expect(content).Should(ContainSubstring("## Intro\n"))
		Expect(content).Should(ContainSubstring("### To vim\n"))
		Expect(content).Should(ContainSubstring("### To reading the manual\n"))
		Expect(content).Should(ContainSubstring("#### User manual\n"))
		Expect(content).Should(ContainSubstring("#### How to use the manuals\n"))
		Expect(content).Should(ContainSubstring("##### Navigate\n"))
		Expect(content).Should(ContainSubstring("##### Search\n"))
		Expect(content).Should(ContainSubstring("## Motion\n"))
		Expect(content).Should(ContainSubstring("### Basic\n"))
		Expect(content).Should(ContainSubstring("### Reference manual\n"))

		Expect(os.Remove(filename)).To(Succeed())
	})

	It("should stop using the hashes after 6 levels", func() {
		graph, err := input.ParseOpml("../testdata/vim-notes-example-6-levels.opml")
		Expect(err).NotTo(HaveOccurred())

		output.WriteMarkdown(filename, graph)
		tempFile.Read(content)

		Expect(content).Should(ContainSubstring("# Vim-notes\n"))
		Expect(content).Should(ContainSubstring("## Intro\n"))
		Expect(content).Should(ContainSubstring("### To vim\n"))
		Expect(content).Should(ContainSubstring("### To reading the manual\n"))
		Expect(content).Should(ContainSubstring("#### User manual\n"))
		Expect(content).Should(ContainSubstring("#### How to use the manuals\n"))
		Expect(content).Should(ContainSubstring("##### Navigate\n"))
		Expect(content).Should(ContainSubstring("###### In nav\n"))
		Expect(content).Should(ContainSubstring("\nIn nav 2\n"))
		Expect(content).Should(ContainSubstring("##### Search\n"))
		Expect(content).Should(ContainSubstring("## Motion\n"))
		Expect(content).Should(ContainSubstring("### Basic\n"))
		Expect(content).Should(ContainSubstring("### Reference manual\n"))

		Expect(os.Remove(filename)).To(Succeed())
	})

	Context("when the file cannot be written to", func() {
		It("should return an error", func() {
			err := output.WriteMarkdown("-/some-random-file-name", nil)
			Expect(err).To(HaveOccurred())
		})
	})
})
