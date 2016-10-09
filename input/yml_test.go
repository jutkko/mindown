package input_test

import (
	"github.com/jutkko/mindown/input"
	"github.com/jutkko/mindown/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Input/Yml", func() {
	It("should parse the deployment name into a graph", func() {
		graph, err := input.ParseYaml("../testdata/simple.yml")
		Expect(err).NotTo(HaveOccurred())
		Expect(graph.GetNodes()[0].GetTitle()).To(Equal("my-redis"))
	})

	It("should parse the instance groups into a graph, with number of instances", func() {
		master := util.NewNode("redis-master", "")
		slave := util.NewNode("redis-slave*2", "")

		graph, err := input.ParseYaml("../testdata/simple-with-instance-groups.yml")
		Expect(err).NotTo(HaveOccurred())

		Expect(graph.GetNodes()[0].GetTitle()).To(Equal("my-redis"))
		children := graph.GetNodes()[0].GetChildren()
		Expect(children).To(ContainElement(master))
		Expect(children).To(ContainElement(slave))
	})

	It("should parse the instance groups into a graph, with number of instances", func() {
		nats := util.NewNode("nats_z1", "")
		etcd := util.NewNode("etcd_z1", "")
		consul := util.NewNode("consul_z1", "")
		blobstore := util.NewNode("blobstore_z1", "")
		postgres := util.NewNode("postgres_z1", "")
		api := util.NewNode("api_z1", "")
		ha_prozy := util.NewNode("ha_proxy_z1", "")
		hm9000 := util.NewNode("hm9000_z1", "")
		doppler := util.NewNode("doppler_z1", "")
		loggregator := util.NewNode("loggregator_trafficcontroller_z1", "")
		uaa := util.NewNode("uaa_z1", "")
		router := util.NewNode("router_z1", "")
		runner := util.NewNode("runner_z1*2", "")

		graph, err := input.ParseYaml("../testdata/cf.yml")
		Expect(err).NotTo(HaveOccurred())

		Expect(graph.GetNodes()[0].GetTitle()).To(Equal("cf"))
		children := graph.GetNodes()[0].GetChildren()

		Expect(children).To(ContainElement(nats))
		Expect(children).To(ContainElement(etcd))
		Expect(children).To(ContainElement(consul))
		Expect(children).To(ContainElement(blobstore))
		Expect(children).To(ContainElement(postgres))
		Expect(children).To(ContainElement(api))
		Expect(children).To(ContainElement(ha_prozy))
		Expect(children).To(ContainElement(hm9000))
		Expect(children).To(ContainElement(doppler))
		Expect(children).To(ContainElement(loggregator))
		Expect(children).To(ContainElement(uaa))
		Expect(children).To(ContainElement(router))
		Expect(children).To(ContainElement(runner))
	})
})
