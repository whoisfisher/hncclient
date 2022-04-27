package v1alpha2

import (
	"github.com/whoisfisher/hncclient/versioned/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/hierarchical-namespaces/api/v1alpha2"
)

type HNCV1Alpha2Interface interface {
	RESTClient() rest.Interface
	//HierarchicalResourceQuotasGetter
	HierarchyConfigurationsGetter
	HNCConfigurationsGetter
	SubnamespaceAnchorsGetter
}

// MonitoringV1Client is used to interact with features provided by the monitoring.coreos.com group.
type HNCV1Alpha2Client struct {
	restClient rest.Interface
}

//
//func (c *HNCV1Alpha2Client) HierarchicalResourceQuotas(namespace string) HierarchicalResourceQuotaInterface {
//	return newHierarchicalResourceQuotas(c, namespace)
//}

func (c *HNCV1Alpha2Client) HierarchyConfigurations(namespace string) HierarchyConfigurationInterface {
	return newHierarchyConfigurations(c, namespace)
}

func (c *HNCV1Alpha2Client) HNCConfigurations(namespace string) HNCConfigurationInterface {
	return newHNCConfigurations(c, namespace)
}

func (c *HNCV1Alpha2Client) SubnamespaceAnchors(namespace string) SubnamespaceAnchorInterface {
	return newSubnamespaceAnchors(c, namespace)
}

// NewForConfig creates a new MonitoringV1Client for the given config.
func NewForConfig(c *rest.Config) (*HNCV1Alpha2Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &HNCV1Alpha2Client{client}, nil
}

// NewForConfigOrDie creates a new MonitoringV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *HNCV1Alpha2Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new MonitoringV1Client for the given RESTClient.
func New(c rest.Interface) *HNCV1Alpha2Client {
	return &HNCV1Alpha2Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha2.GroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *HNCV1Alpha2Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
