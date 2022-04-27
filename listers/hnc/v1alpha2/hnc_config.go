package v1alpha2

import (
	"github.com/whoisfisher/hncclient"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	"sigs.k8s.io/hierarchical-namespaces/api/v1alpha2"
)

type HNCConfigurationLister interface {
	List(selector labels.Selector) (ret []*v1alpha2.HNCConfiguration, err error)
	HNCConfigurations(namespace string) HNCConfigurationNamespaceLister
	HNCConfigurationListerExpansion
}

type hncConfigurationLister struct {
	indexer cache.Indexer
}

func NewHNCConfigurationLister(indexer cache.Indexer) HNCConfigurationLister {
	return &hncConfigurationLister{indexer: indexer}
}

func (s *hncConfigurationLister) List(selector labels.Selector) (ret []*v1alpha2.HNCConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.HNCConfiguration))
	})
	return ret, err
}

func (s *hncConfigurationLister) HNCConfigurations(namespace string) HNCConfigurationNamespaceLister {
	return hncConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

type HNCConfigurationNamespaceLister interface {
	List(selector labels.Selector) (ret []*v1alpha2.HNCConfiguration, err error)
	Get(name string) (*v1alpha2.HNCConfiguration, error)
	HNCConfigurationNamespaceListerExpansion
}

type hncConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

func (s hncConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.HNCConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.HNCConfiguration))
	})
	return ret, err
}

func (s hncConfigurationNamespaceLister) Get(name string) (*v1alpha2.HNCConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(hncclient.Resource("hncconfiguration"), name)
	}
	return obj.(*v1alpha2.HNCConfiguration), nil
}
