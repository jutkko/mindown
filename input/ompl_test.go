package input_test

import (
	"github.com/jutkko/mindown/input"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ompl", func() {

	It("should parse the file into the graph", func() {
		_, err := input.ParseOpml("../testdata/simple.opml")
		Expect(err).NotTo(HaveOccurred())
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
