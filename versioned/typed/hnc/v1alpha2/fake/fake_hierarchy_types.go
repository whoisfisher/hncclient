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
type FakeHierarchyConfigurations struct {
	Fake *FakeHNCV1Alpha2
	ns   string
}

var HierarchyConfigurationResource = schema.GroupVersionResource{Group: v1alpha2.GroupVersion.Group, Version: v1alpha2.GroupVersion.Version, Resource: "hierarchyconfigurations"}

var HierarchyConfigurationKind = schema.GroupVersionKind{Group: v1alpha2.GroupVersion.Group, Version: v1alpha2.GroupVersion.Version, Kind: "HierarchyConfiguration"}

// Get takes name of the alertmanager, and returns the corresponding alertmanager object, and an error if there is any.
func (c *FakeHierarchyConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.HierarchyConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(HierarchyConfigurationResource, c.ns, name), &v1alpha2.HierarchyConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.HierarchyConfiguration), err
}

// List takes label and field selectors, and returns the list of Alertmanagers that match those selectors.
func (c *FakeHierarchyConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.HierarchyConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(HierarchyConfigurationResource, HierarchyConfigurationKind, c.ns, opts), &v1alpha2.HierarchyConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.HierarchyConfigurationList{ListMeta: obj.(*v1alpha2.HierarchyConfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha2.HierarchyConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested alertmanagers.
func (c *FakeHierarchyConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(HierarchyConfigurationResource, c.ns, opts))

}

// Create takes the representation of a alertmanager and creates it.  Returns the server's representation of the alertmanager, and an error, if there is any.
func (c *FakeHierarchyConfigurations) Create(ctx context.Context, hierarchyConfiguration *v1alpha2.HierarchyConfiguration, opts v1.CreateOptions) (result *v1alpha2.HierarchyConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(HierarchyConfigurationResource, c.ns, hierarchyConfiguration), &v1alpha2.HierarchyConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.HierarchyConfiguration), err
}

// Update takes the representation of a alertmanager and updates it. Returns the server's representation of the alertmanager, and an error, if there is any.
func (c *FakeHierarchyConfigurations) Update(ctx context.Context, hierarchyConfiguration *v1alpha2.HierarchyConfiguration, opts v1.UpdateOptions) (result *v1alpha2.HierarchyConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(HierarchyConfigurationResource, c.ns, hierarchyConfiguration), &v1alpha2.HierarchyConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.HierarchyConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHierarchyConfigurations) UpdateStatus(ctx context.Context, hierarchyConfiguration *v1alpha2.HierarchyConfiguration, opts v1.UpdateOptions) (*v1alpha2.HierarchyConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(HierarchyConfigurationResource, "status", c.ns, hierarchyConfiguration), &v1alpha2.HierarchyConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.HierarchyConfiguration), err
}

// Delete takes name of the alertmanager and deletes it. Returns an error if one occurs.
func (c *FakeHierarchyConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(HierarchyConfigurationResource, c.ns, name), &v1alpha2.HierarchyConfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHierarchyConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(HierarchyConfigurationResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.HierarchyConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched alertmanager.
func (c *FakeHierarchyConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.HierarchyConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(HierarchyConfigurationResource, c.ns, name, pt, data, subresources...), &v1alpha2.HierarchyConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.HierarchyConfiguration), err
}
