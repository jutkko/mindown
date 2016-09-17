package output_test

import (
	"io/ioutil"
	"math/rand"
	"os"

	"github.com/jutkko/mindown/input"
	"github.com/jutkko/mindown/output"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

var _ = Describe("Markdown", func() {
	var (
		tempFile *os.File
		filename string
		content  []byte
	)

	BeforeEach(func() {
		filename = randSeq(10)
		content = make([]byte, 300)
	})

	It("should output the correct file given simple graph", func() {
		graph, err := input.ParseOpml("../testdata/simple.opml")
		Expect(err).NotTo(HaveOccurred())

		err = output.WriteMarkdown(filename, false, graph)
		Expect(err).NotTo(HaveOccurred())

		tempFile, err = os.OpenFile(filename, os.O_RDONLY, 0600)
		Expect(err).NotTo(HaveOccurred())
		tempFile.Read(content)
		Expect(content).Should(ContainSubstring("# Vim-notes\n"))

		Expect(os.Remove(filename)).To(Succeed())
	})

	Context("when the file already exists", func() {
		It("should return an error", func() {
			graph, err := input.ParseOpml("../testdata/simple.opml")
			Expect(err).NotTo(HaveOccurred())

			_, err = os.Create(filename)
			Expect(err).NotTo(HaveOccurred())

			err = output.WriteMarkdown(filename, false, graph)
			Expect(err).To(HaveOccurred())

			Expect(os.Remove(filename)).To(Succeed())
		})
	})

	Context("when the overwrite flag is provided", func() {
		It("should output the correct file given simple graph", func() {
			graph, err := input.ParseOpml("../testdata/simple.opml")
			Expect(err).NotTo(HaveOccurred())

			err = output.WriteMarkdown(filename, true, graph)
			Expect(err).NotTo(HaveOccurred())

			tempFile, err = os.OpenFile(filename, os.O_RDONLY, 0600)
			Expect(err).NotTo(HaveOccurred())
			tempFile.Read(content)
			Expect(content).Should(ContainSubstring("# Vim-notes\n"))

			Expect(os.Remove(filename)).To(Succeed())
		})

		Context("when the file already exists", func() {
			It("should output the correct file given simple graph overwriting the original file", func() {
				graph, err := input.ParseOpml("../testdata/simple.opml")
				Expect(err).NotTo(HaveOccurred())

				_, err = os.Create(filename)
				Expect(err).NotTo(HaveOccurred())

				Expect(ioutil.WriteFile(filename, []byte("yoyoyo"), 0600))

				err = output.WriteMarkdown(filename, true, graph)
				Expect(err).NotTo(HaveOccurred())

				tempFile, err = os.OpenFile(filename, os.O_RDONLY, 0600)
				Expect(err).NotTo(HaveOccurred())
				tempFile.Read(content)
				Expect(content).Should(ContainSubstring("# Vim-notes\n"))
				Expect(content).ShouldNot(ContainSubstring("yoyoyo"))

				Expect(os.Remove(filename)).To(Succeed())
			})
		})
	})

	It("should output the correct file given simple-two-root graph", func() {
		graph, err := input.ParseOpml("../testdata/simple-two-root.opml")
		Expect(err).NotTo(HaveOccurred())

		err = output.WriteMarkdown(filename, true, graph)
		Expect(err).NotTo(HaveOccurred())

		tempFile, err = os.OpenFile(filename, os.O_RDONLY, 0600)
		tempFile.Read(content)
		Expect(content).Should(ContainSubstring("# Vim-notes\n"))
		Expect(content).Should(ContainSubstring("# Vim-notes1\n"))

		Expect(os.Remove(filename)).To(Succeed())
	})

	It("should output the correct file given a more complex graph", func() {
		graph, err := input.ParseOpml("../testdata/vim-notes-example-simple.opml")
		Expect(err).NotTo(HaveOccurred())

		err = output.WriteMarkdown(filename, true, graph)
		Expect(err).NotTo(HaveOccurred())

		tempFile, err = os.OpenFile(filename, os.O_RDONLY, 0600)
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

		err = output.WriteMarkdown(filename, true, graph)
		Expect(err).NotTo(HaveOccurred())

		tempFile, err = os.OpenFile(filename, os.O_RDONLY, 0600)
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

	Context("when the graph is nil", func() {
		It("should return an error", func() {
			err := output.WriteMarkdown("some-random-file-name", true, nil)
			Expect(err).To(MatchError("Graph is nil"))
		})
	})

	Context("when the file cannot be written to", func() {
		It("should return an error", func() {
			graph, err := input.ParseOpml("../testdata/simple.opml")
			Expect(err).NotTo(HaveOccurred())

			err = output.WriteMarkdown("-/some-random-file-name", true, graph)
			Expect(err.Error()).To(ContainSubstring("Failed to open file"))
		})
	})

	Context("when the file cannot does not exist", func() {
		It("should just work", func() {
			graph, err := input.ParseOpml("../testdata/simple.opml")
			Expect(err).NotTo(HaveOccurred())

			err = output.WriteMarkdown("some-random-file-name", true, graph)
			Expect(err).NotTo(HaveOccurred())
			Expect(os.Remove("some-random-file-name")).To(Succeed())
		})
	})
})
