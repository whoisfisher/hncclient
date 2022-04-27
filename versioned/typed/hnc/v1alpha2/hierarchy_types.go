package v1alpha2

import (
	"context"
	"github.com/whoisfisher/hncclient/versioned/scheme"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	"sigs.k8s.io/hierarchical-namespaces/api/v1alpha2"
)

// AlertmanagersGetter has a method to return a AlertmanagerInterface.
// A group's client should implement this interface.
type HierarchyConfigurationsGetter interface {
	HierarchyConfigurations(namespace string) HierarchyConfigurationInterface
}

// AlertmanagerInterface has methods to work with Alertmanager resources.
type HierarchyConfigurationInterface interface {
	Create(ctx context.Context, hierarchyConfiguration *v1alpha2.HierarchyConfiguration, opts metav1.CreateOptions) (*v1alpha2.HierarchyConfiguration, error)
	Update(ctx context.Context, hierarchyConfiguration *v1alpha2.HierarchyConfiguration, opts metav1.UpdateOptions) (*v1alpha2.HierarchyConfiguration, error)
	UpdateStatus(ctx context.Context, hierarchyConfiguration *v1alpha2.HierarchyConfiguration, opts metav1.UpdateOptions) (*v1alpha2.HierarchyConfiguration, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1alpha2.HierarchyConfiguration, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1alpha2.HierarchyConfigurationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha2.HierarchyConfiguration, err error)
	HierarchyConfigurationExpansion
}

// alertmanagers implements AlertmanagerInterface
type hierarchyConfigurations struct {
	client rest.Interface
	ns     string
}

// newAlertmanagers returns a Alertmanagers
func newHierarchyConfigurations(c *HNCV1Alpha2Client, namespace string) *hierarchyConfigurations {
	return &hierarchyConfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the alertmanager, and returns the corresponding alertmanager object, and an error if there is any.
func (c *hierarchyConfigurations) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1alpha2.HierarchyConfiguration, err error) {
	result = &v1alpha2.HierarchyConfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hierarchyconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Alertmanagers that match those selectors.
func (c *hierarchyConfigurations) List(ctx context.Context, opts metav1.ListOptions) (result *v1alpha2.HierarchyConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.HierarchyConfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hierarchyconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested alertmanagers.
func (c *hierarchyConfigurations) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("hierarchyconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a alertmanager and creates it.  Returns the server's representation of the alertmanager, and an error, if there is any.
func (c *hierarchyConfigurations) Create(ctx context.Context, hierarchyConfiguration *v1alpha2.HierarchyConfiguration, opts metav1.CreateOptions) (result *v1alpha2.HierarchyConfiguration, err error) {
	result = &v1alpha2.HierarchyConfiguration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("hierarchyconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hierarchyConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a alertmanager and updates it. Returns the server's representation of the alertmanager, and an error, if there is any.
func (c *hierarchyConfigurations) Update(ctx context.Context, hierarchyConfiguration *v1alpha2.HierarchyConfiguration, opts metav1.UpdateOptions) (result *v1alpha2.HierarchyConfiguration, err error) {
	result = &v1alpha2.HierarchyConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hierarchyconfigurations").
		Name(hierarchyConfiguration.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hierarchyConfiguration).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *hierarchyConfigurations) UpdateStatus(ctx context.Context, hierarchyConfiguration *v1alpha2.HierarchyConfiguration, opts metav1.UpdateOptions) (result *v1alpha2.HierarchyConfiguration, err error) {
	result = &v1alpha2.HierarchyConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hierarchyconfigurations").
		Name(hierarchyConfiguration.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(hierarchyConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the alertmanager and deletes it. Returns an error if one occurs.
func (c *hierarchyConfigurations) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hierarchyconfigurations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hierarchyConfigurations) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hierarchyconfigurations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched alertmanager.
func (c *hierarchyConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha2.HierarchyConfiguration, err error) {
	result = &v1alpha2.HierarchyConfiguration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("hierarchyconfigurations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
