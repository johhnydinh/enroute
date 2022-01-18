// SPDX-License-Identifier: Apache-2.0
// Copyright(c) 2018-2021 Saaras Inc.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/saarasio/enroute/enroute-dp/apis/enroute/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RouteFilterLister helps list RouteFilters.
// All objects returned here must be treated as read-only.
type RouteFilterLister interface {
	// List lists all RouteFilters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.RouteFilter, err error)
	// RouteFilters returns an object that can list and get RouteFilters.
	RouteFilters(namespace string) RouteFilterNamespaceLister
	RouteFilterListerExpansion
}

// routeFilterLister implements the RouteFilterLister interface.
type routeFilterLister struct {
	indexer cache.Indexer
}

// NewRouteFilterLister returns a new RouteFilterLister.
func NewRouteFilterLister(indexer cache.Indexer) RouteFilterLister {
	return &routeFilterLister{indexer: indexer}
}

// List lists all RouteFilters in the indexer.
func (s *routeFilterLister) List(selector labels.Selector) (ret []*v1.RouteFilter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.RouteFilter))
	})
	return ret, err
}

// RouteFilters returns an object that can list and get RouteFilters.
func (s *routeFilterLister) RouteFilters(namespace string) RouteFilterNamespaceLister {
	return routeFilterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RouteFilterNamespaceLister helps list and get RouteFilters.
// All objects returned here must be treated as read-only.
type RouteFilterNamespaceLister interface {
	// List lists all RouteFilters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.RouteFilter, err error)
	// Get retrieves the RouteFilter from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.RouteFilter, error)
	RouteFilterNamespaceListerExpansion
}

// routeFilterNamespaceLister implements the RouteFilterNamespaceLister
// interface.
type routeFilterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RouteFilters in the indexer for a given namespace.
func (s routeFilterNamespaceLister) List(selector labels.Selector) (ret []*v1.RouteFilter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.RouteFilter))
	})
	return ret, err
}

// Get retrieves the RouteFilter from the indexer for a given namespace and name.
func (s routeFilterNamespaceLister) Get(name string) (*v1.RouteFilter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("routefilter"), name)
	}
	return obj.(*v1.RouteFilter), nil
}
