package v1alpha2

import (
	"github.com/whoisfisher/hncclient"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	"sigs.k8s.io/hierarchical-namespaces/api/v1alpha2"
)

type HierarchyConfigurationLister interface {
	List(selector labels.Selector) (ret []*v1alpha2.HierarchyConfiguration, err error)
	HierarchyConfigurations(namespace string) HierarchyConfigurationNamespaceLister
	HierarchyConfigurationListerExpansion
}

type hierarchyConfigurationLister struct {
	indexer cache.Indexer
}

func NewHierarchyConfigurationLister(indexer cache.Indexer) HierarchyConfigurationLister {
	return &hierarchyConfigurationLister{indexer: indexer}
}

func (s *hierarchyConfigurationLister) List(selector labels.Selector) (ret []*v1alpha2.HierarchyConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.HierarchyConfiguration))
	})
	return ret, err
}

func (s *hierarchyConfigurationLister) HierarchyConfigurations(namespace string) HierarchyConfigurationNamespaceLister {
	return hierarchyConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

type HierarchyConfigurationNamespaceLister interface {
	List(selector labels.Selector) (ret []*v1alpha2.HierarchyConfiguration, err error)
	Get(name string) (*v1alpha2.HierarchyConfiguration, error)
	HierarchyConfigurationNamespaceListerExpansion
}

type hierarchyConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

func (s hierarchyConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.HierarchyConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.HierarchyConfiguration))
	})
	return ret, err
}

func (s hierarchyConfigurationNamespaceLister) Get(name string) (*v1alpha2.HierarchyConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(hncclient.Resource("hierarchyconfiguration"), name)
	}
	return obj.(*v1alpha2.HierarchyConfiguration), nil
}
