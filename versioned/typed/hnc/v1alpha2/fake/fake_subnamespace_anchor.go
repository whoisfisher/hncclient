package fake

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/hierarchical-namespaces/api/v1alpha2"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAlertmanagers implements AlertmanagerInterface
type FakeSubnamespaceAnchors struct {
	Fake *FakeHNCV1Alpha2
	ns   string
}

var SubnamespaceAnchorResource = schema.GroupVersionResource{Group: v1alpha2.GroupVersion.Group, Version: v1alpha2.GroupVersion.Version, Resource: "subnamespaceanchors"}

var SubnamespaceAnchorKind = schema.GroupVersionKind{Group: v1alpha2.GroupVersion.Group, Version: v1alpha2.GroupVersion.Version, Kind: "SubnamespaceAnchor"}

// Get takes name of the alertmanager, and returns the corresponding alertmanager object, and an error if there is any.
func (c *FakeSubnamespaceAnchors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.SubnamespaceAnchor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(SubnamespaceAnchorResource, c.ns, name), &v1alpha2.SubnamespaceAnchor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.SubnamespaceAnchor), err
}

// List takes label and field selectors, and returns the list of Alertmanagers that match those selectors.
func (c *FakeSubnamespaceAnchors) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.SubnamespaceAnchorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(SubnamespaceAnchorResource, SubnamespaceAnchorKind, c.ns, opts), &v1alpha2.SubnamespaceAnchorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.SubnamespaceAnchorList{ListMeta: obj.(*v1alpha2.SubnamespaceAnchorList).ListMeta}
	for _, item := range obj.(*v1alpha2.SubnamespaceAnchorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested alertmanagers.
func (c *FakeSubnamespaceAnchors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(SubnamespaceAnchorResource, c.ns, opts))

}

// Create takes the representation of a alertmanager and creates it.  Returns the server's representation of the alertmanager, and an error, if there is any.
func (c *FakeSubnamespaceAnchors) Create(ctx context.Context, subnamespaceAnchor *v1alpha2.SubnamespaceAnchor, opts v1.CreateOptions) (result *v1alpha2.SubnamespaceAnchor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(HNCConfigurationResource, c.ns, subnamespaceAnchor), &v1alpha2.SubnamespaceAnchor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.SubnamespaceAnchor), err
}

// Update takes the representation of a alertmanager and updates it. Returns the server's representation of the alertmanager, and an error, if there is any.
func (c *FakeSubnamespaceAnchors) Update(ctx context.Context, subnamespaceAnchor *v1alpha2.SubnamespaceAnchor, opts v1.UpdateOptions) (result *v1alpha2.SubnamespaceAnchor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(SubnamespaceAnchorResource, c.ns, subnamespaceAnchor), &v1alpha2.HNCConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.SubnamespaceAnchor), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSubnamespaceAnchors) UpdateStatus(ctx context.Context, hncConfiguration *v1alpha2.SubnamespaceAnchor, opts v1.UpdateOptions) (*v1alpha2.SubnamespaceAnchor, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(SubnamespaceAnchorResource, "status", c.ns, hncConfiguration), &v1alpha2.SubnamespaceAnchor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.SubnamespaceAnchor), err
}

// Delete takes name of the alertmanager and deletes it. Returns an error if one occurs.
func (c *FakeSubnamespaceAnchors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(HNCConfigurationResource, c.ns, name), &v1alpha2.SubnamespaceAnchor{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSubnamespaceAnchors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(SubnamespaceAnchorResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.SubnamespaceAnchorList{})
	return err
}

// Patch applies the patch and returns the patched alertmanager.
func (c *FakeSubnamespaceAnchors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.SubnamespaceAnchor, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(SubnamespaceAnchorResource, c.ns, name, pt, data, subresources...), &v1alpha2.SubnamespaceAnchor{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.SubnamespaceAnchor), err
}
