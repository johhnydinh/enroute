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

// GlobalConfigLister helps list GlobalConfigs.
// All objects returned here must be treated as read-only.
type GlobalConfigLister interface {
	// List lists all GlobalConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.GlobalConfig, err error)
	// GlobalConfigs returns an object that can list and get GlobalConfigs.
	GlobalConfigs(namespace string) GlobalConfigNamespaceLister
	GlobalConfigListerExpansion
}

// globalConfigLister implements the GlobalConfigLister interface.
type globalConfigLister struct {
	indexer cache.Indexer
}

// NewGlobalConfigLister returns a new GlobalConfigLister.
func NewGlobalConfigLister(indexer cache.Indexer) GlobalConfigLister {
	return &globalConfigLister{indexer: indexer}
}

// List lists all GlobalConfigs in the indexer.
func (s *globalConfigLister) List(selector labels.Selector) (ret []*v1.GlobalConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.GlobalConfig))
	})
	return ret, err
}

// GlobalConfigs returns an object that can list and get GlobalConfigs.
func (s *globalConfigLister) GlobalConfigs(namespace string) GlobalConfigNamespaceLister {
	return globalConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GlobalConfigNamespaceLister helps list and get GlobalConfigs.
// All objects returned here must be treated as read-only.
type GlobalConfigNamespaceLister interface {
	// List lists all GlobalConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.GlobalConfig, err error)
	// Get retrieves the GlobalConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.GlobalConfig, error)
	GlobalConfigNamespaceListerExpansion
}

// globalConfigNamespaceLister implements the GlobalConfigNamespaceLister
// interface.
type globalConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GlobalConfigs in the indexer for a given namespace.
func (s globalConfigNamespaceLister) List(selector labels.Selector) (ret []*v1.GlobalConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.GlobalConfig))
	})
	return ret, err
}

// Get retrieves the GlobalConfig from the indexer for a given namespace and name.
func (s globalConfigNamespaceLister) Get(name string) (*v1.GlobalConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("globalconfig"), name)
	}
	return obj.(*v1.GlobalConfig), nil
}
