package v1alpha2

import (
	"context"
	"github.com/whoisfisher/hncclient/listers/hnc/v1alpha2"
	time "time"

	internalinterfaces "github.com/whoisfisher/hncclient/informers/externalversions/internalinterfaces"
	versioned "github.com/whoisfisher/hncclient/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	hncv1alpha2 "sigs.k8s.io/hierarchical-namespaces/api/v1alpha2"
)

type SubnamespaceAnchorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.SubnamespaceAnchorLister
}

type subnamespaceAnchorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

func NewSubnamespaceAnchorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSubnamespaceAnchorInformer(client, namespace, resyncPeriod, indexers, nil)
}

func NewFilteredSubnamespaceAnchorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HNCV1Alpha2().SubnamespaceAnchors(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HNCV1Alpha2().SubnamespaceAnchors(namespace).Watch(context.TODO(), options)
			},
		},
		&hncv1alpha2.SubnamespaceAnchor{},
		resyncPeriod,
		indexers,
	)
}

func (f *subnamespaceAnchorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSubnamespaceAnchorInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *subnamespaceAnchorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&hncv1alpha2.SubnamespaceAnchor{}, f.defaultInformer)
}

func (f *subnamespaceAnchorInformer) Lister() v1alpha2.SubnamespaceAnchorLister {
	return v1alpha2.NewSubnamespaceAnchorLister(f.Informer().GetIndexer())
}
