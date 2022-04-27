package fake

import (
	v1alpha2 "github.com/whoisfisher/hncclient/versioned/typed/hnc/v1alpha2"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeHNCV1Alpha2 struct {
	*testing.Fake
}

func (c *FakeHNCV1Alpha2) HierarchyConfigurations(namespace string) v1alpha2.HierarchyConfigurationInterface {
	return &FakeHierarchyConfigurations{c, namespace}
}

func (c *FakeHNCV1Alpha2) HNCConfigurations(namespace string) v1alpha2.HNCConfigurationInterface {
	return &FakeHNCConfigurations{c, namespace}
}

func (c *FakeHNCV1Alpha2) SubnamespaceAnchors(namespace string) v1alpha2.SubnamespaceAnchorInterface {
	return &FakeSubnamespaceAnchors{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeHNCV1Alpha2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
