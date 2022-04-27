package hnc

import (
	hncv1alpha2 "github.com/whoisfisher/hncclient/informers/externalversions/hnc/v1alpha2"
	internalinterfaces "github.com/whoisfisher/hncclient/informers/externalversions/internalinterfaces"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	V1alpha2() hncv1alpha2.Interface
}

type group struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &group{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

func (g *group) V1alpha2() hncv1alpha2.Interface {
	return hncv1alpha2.New(g.factory, g.namespace, g.tweakListOptions)
}
