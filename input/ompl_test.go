package input_test

import (
	"github.com/jutkko/mindown/input"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ompl", func() {
	var (
		err error
	)

	It("should parse the file into the graph", func() {
		_, err = input.ParseOpml("../testdata/simple.opml")
		Expect(err).NotTo(HaveOccurred())
	})
})
