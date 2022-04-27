package v1alpha2

import (
	internalinterfaces "github.com/whoisfisher/hncclient/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	HierarchyConfigurations() HierarchyConfigurationInformer
	HNCConfigurations() HNCConfigurationInformer
	SubnamespaceAnchors() SubnamespaceAnchorInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Alertmanagers returns a AlertmanagerInformer.
func (v *version) HierarchyConfigurations() HierarchyConfigurationInformer {
	return &hierarchyConfigurationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PodMonitors returns a PodMonitorInformer.
func (v *version) HNCConfigurations() HNCConfigurationInformer {
	return &hncConfigurationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Probes returns a ProbeInformer.
func (v *version) SubnamespaceAnchors() SubnamespaceAnchorInformer {
	return &subnamespaceAnchorInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
