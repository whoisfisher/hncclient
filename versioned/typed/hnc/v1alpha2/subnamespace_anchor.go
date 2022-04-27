package v1alpha2

import (
	"context"
	"time"

	scheme "github.com/whoisfisher/hncclient/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	"sigs.k8s.io/hierarchical-namespaces/api/v1alpha2"
)

// AlertmanagersGetter has a method to return a AlertmanagerInterface.
// A group's client should implement this interface.
type SubnamespaceAnchorsGetter interface {
	SubnamespaceAnchors(namespace string) SubnamespaceAnchorInterface
}

// AlertmanagerInterface has methods to work with Alertmanager resources.
type SubnamespaceAnchorInterface interface {
	Create(ctx context.Context, subnamespaceAnchor *v1alpha2.SubnamespaceAnchor, opts metav1.CreateOptions) (*v1alpha2.SubnamespaceAnchor, error)
	Update(ctx context.Context, subnamespaceAnchor *v1alpha2.SubnamespaceAnchor, opts metav1.UpdateOptions) (*v1alpha2.SubnamespaceAnchor, error)
	UpdateStatus(ctx context.Context, subnamespaceAnchor *v1alpha2.SubnamespaceAnchor, opts metav1.UpdateOptions) (*v1alpha2.SubnamespaceAnchor, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1alpha2.SubnamespaceAnchor, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1alpha2.SubnamespaceAnchorList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha2.SubnamespaceAnchor, err error)
	SubnamespaceAnchorExpansion
}

// alertmanagers implements AlertmanagerInterface
type subnamespaceAnchors struct {
	client rest.Interface
	ns     string
}

// newAlertmanagers returns a Alertmanagers
func newSubnamespaceAnchors(c *HNCV1Alpha2Client, namespace string) *subnamespaceAnchors {
	return &subnamespaceAnchors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the alertmanager, and returns the corresponding alertmanager object, and an error if there is any.
func (c *subnamespaceAnchors) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1alpha2.SubnamespaceAnchor, err error) {
	result = &v1alpha2.SubnamespaceAnchor{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("subnamespaceanchors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Alertmanagers that match those selectors.
func (c *subnamespaceAnchors) List(ctx context.Context, opts metav1.ListOptions) (result *v1alpha2.SubnamespaceAnchorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.SubnamespaceAnchorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("subnamespaceanchors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested alertmanagers.
func (c *subnamespaceAnchors) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("subnamespaceanchors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a alertmanager and creates it.  Returns the server's representation of the alertmanager, and an error, if there is any.
func (c *subnamespaceAnchors) Create(ctx context.Context, subnamespaceAnchor *v1alpha2.SubnamespaceAnchor, opts metav1.CreateOptions) (result *v1alpha2.SubnamespaceAnchor, err error) {
	result = &v1alpha2.SubnamespaceAnchor{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("subnamespaceanchors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(subnamespaceAnchor).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a alertmanager and updates it. Returns the server's representation of the alertmanager, and an error, if there is any.
func (c *subnamespaceAnchors) Update(ctx context.Context, subnamespaceAnchor *v1alpha2.SubnamespaceAnchor, opts metav1.UpdateOptions) (result *v1alpha2.SubnamespaceAnchor, err error) {
	result = &v1alpha2.SubnamespaceAnchor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("subnamespaceanchors").
		Name(subnamespaceAnchor.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(subnamespaceAnchor).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *subnamespaceAnchors) UpdateStatus(ctx context.Context, subnamespaceAnchor *v1alpha2.SubnamespaceAnchor, opts metav1.UpdateOptions) (result *v1alpha2.SubnamespaceAnchor, err error) {
	result = &v1alpha2.SubnamespaceAnchor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("subnamespaceanchors").
		Name(subnamespaceAnchor.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(subnamespaceAnchor).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the alertmanager and deletes it. Returns an error if one occurs.
func (c *subnamespaceAnchors) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("subnamespaceanchors").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *subnamespaceAnchors) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("subnamespaceanchors").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched alertmanager.
func (c *subnamespaceAnchors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1alpha2.SubnamespaceAnchor, err error) {
	result = &v1alpha2.SubnamespaceAnchor{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("subnamespaceanchors").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
