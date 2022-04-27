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
type FakeHNCConfigurations struct {
	Fake *FakeHNCV1Alpha2
	ns   string
}

var HNCConfigurationResource = schema.GroupVersionResource{Group: v1alpha2.GroupVersion.Group, Version: v1alpha2.GroupVersion.Version, Resource: "hncconfigurations"}

var HNCConfigurationKind = schema.GroupVersionKind{Group: v1alpha2.GroupVersion.Group, Version: v1alpha2.GroupVersion.Version, Kind: "HNCConfiguration"}

// Get takes name of the alertmanager, and returns the corresponding alertmanager object, and an error if there is any.
func (c *FakeHNCConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.HNCConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(HNCConfigurationResource, c.ns, name), &v1alpha2.HierarchyConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.HNCConfiguration), err
}

// List takes label and field selectors, and returns the list of Alertmanagers that match those selectors.
func (c *FakeHNCConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.HNCConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(HNCConfigurationResource, HNCConfigurationKind, c.ns, opts), &v1alpha2.HNCConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.HNCConfigurationList{ListMeta: obj.(*v1alpha2.HNCConfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha2.HNCConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested alertmanagers.
func (c *FakeHNCConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(HNCConfigurationResource, c.ns, opts))

}

// Create takes the representation of a alertmanager and creates it.  Returns the server's representation of the alertmanager, and an error, if there is any.
func (c *FakeHNCConfigurations) Create(ctx context.Context, HNCConfiguration *v1alpha2.HNCConfiguration, opts v1.CreateOptions) (result *v1alpha2.HNCConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(HNCConfigurationResource, c.ns, HNCConfiguration), &v1alpha2.HierarchyConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.HNCConfiguration), err
}

// Update takes the representation of a alertmanager and updates it. Returns the server's representation of the alertmanager, and an error, if there is any.
func (c *FakeHNCConfigurations) Update(ctx context.Context, hncConfiguration *v1alpha2.HNCConfiguration, opts v1.UpdateOptions) (result *v1alpha2.HNCConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(HNCConfigurationResource, c.ns, hncConfiguration), &v1alpha2.HNCConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.HNCConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHNCConfigurations) UpdateStatus(ctx context.Context, hncConfiguration *v1alpha2.HNCConfiguration, opts v1.UpdateOptions) (*v1alpha2.HNCConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(HNCConfigurationResource, "status", c.ns, hncConfiguration), &v1alpha2.HNCConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.HNCConfiguration), err
}

// Delete takes name of the alertmanager and deletes it. Returns an error if one occurs.
func (c *FakeHNCConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(HNCConfigurationResource, c.ns, name), &v1alpha2.HNCConfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHNCConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(HNCConfigurationResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.HNCConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched alertmanager.
func (c *FakeHNCConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.HNCConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(HNCConfigurationResource, c.ns, name, pt, data, subresources...), &v1alpha2.HNCConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.HNCConfiguration), err
}
