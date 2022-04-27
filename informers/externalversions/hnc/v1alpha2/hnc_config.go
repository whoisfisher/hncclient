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

type HNCConfigurationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.HNCConfigurationLister
}

type hncConfigurationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

func NewHNCConfigurationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHNCConfigurationInformer(client, namespace, resyncPeriod, indexers, nil)
}

func NewFilteredHNCConfigurationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HNCV1Alpha2().HNCConfigurations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HNCV1Alpha2().HNCConfigurations(namespace).Watch(context.TODO(), options)
			},
		},
		&hncv1alpha2.HNCConfiguration{},
		resyncPeriod,
		indexers,
	)
}

func (f *hncConfigurationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHNCConfigurationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hncConfigurationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&hncv1alpha2.HNCConfiguration{}, f.defaultInformer)
}

func (f *hncConfigurationInformer) Lister() v1alpha2.HNCConfigurationLister {
	return v1alpha2.NewHNCConfigurationLister(f.Informer().GetIndexer())
}
