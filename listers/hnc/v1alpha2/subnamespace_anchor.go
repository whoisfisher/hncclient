package v1alpha2

import (
	"github.com/whoisfisher/hncclient"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	"sigs.k8s.io/hierarchical-namespaces/api/v1alpha2"
)

type SubnamespaceAnchorLister interface {
	List(selector labels.Selector) (ret []*v1alpha2.SubnamespaceAnchor, err error)
	SubnamespaceAnchors(namespace string) SubnamespaceAnchorNamespaceLister
	SubnamespaceAnchorListerExpansion
}

type subnamespaceAnchorLister struct {
	indexer cache.Indexer
}

func NewSubnamespaceAnchorLister(indexer cache.Indexer) SubnamespaceAnchorLister {
	return &subnamespaceAnchorLister{indexer: indexer}
}

func (s *subnamespaceAnchorLister) List(selector labels.Selector) (ret []*v1alpha2.SubnamespaceAnchor, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.SubnamespaceAnchor))
	})
	return ret, err
}

func (s *subnamespaceAnchorLister) SubnamespaceAnchors(namespace string) SubnamespaceAnchorNamespaceLister {
	return subnamespaceAnchorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

type SubnamespaceAnchorNamespaceLister interface {
	List(selector labels.Selector) (ret []*v1alpha2.SubnamespaceAnchor, err error)
	Get(name string) (*v1alpha2.SubnamespaceAnchor, error)
	SubnamespaceAnchorNamespaceListerExpansion
}

type subnamespaceAnchorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

func (s subnamespaceAnchorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.SubnamespaceAnchor, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.SubnamespaceAnchor))
	})
	return ret, err
}

func (s subnamespaceAnchorNamespaceLister) Get(name string) (*v1alpha2.SubnamespaceAnchor, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(hncclient.Resource("subnamespaceanchor"), name)
	}
	return obj.(*v1alpha2.SubnamespaceAnchor), nil
}
