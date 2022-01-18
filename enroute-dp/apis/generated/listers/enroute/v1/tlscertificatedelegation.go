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

// TLSCertificateDelegationLister helps list TLSCertificateDelegations.
// All objects returned here must be treated as read-only.
type TLSCertificateDelegationLister interface {
	// List lists all TLSCertificateDelegations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.TLSCertificateDelegation, err error)
	// TLSCertificateDelegations returns an object that can list and get TLSCertificateDelegations.
	TLSCertificateDelegations(namespace string) TLSCertificateDelegationNamespaceLister
	TLSCertificateDelegationListerExpansion
}

// tLSCertificateDelegationLister implements the TLSCertificateDelegationLister interface.
type tLSCertificateDelegationLister struct {
	indexer cache.Indexer
}

// NewTLSCertificateDelegationLister returns a new TLSCertificateDelegationLister.
func NewTLSCertificateDelegationLister(indexer cache.Indexer) TLSCertificateDelegationLister {
	return &tLSCertificateDelegationLister{indexer: indexer}
}

// List lists all TLSCertificateDelegations in the indexer.
func (s *tLSCertificateDelegationLister) List(selector labels.Selector) (ret []*v1.TLSCertificateDelegation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TLSCertificateDelegation))
	})
	return ret, err
}

// TLSCertificateDelegations returns an object that can list and get TLSCertificateDelegations.
func (s *tLSCertificateDelegationLister) TLSCertificateDelegations(namespace string) TLSCertificateDelegationNamespaceLister {
	return tLSCertificateDelegationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TLSCertificateDelegationNamespaceLister helps list and get TLSCertificateDelegations.
// All objects returned here must be treated as read-only.
type TLSCertificateDelegationNamespaceLister interface {
	// List lists all TLSCertificateDelegations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.TLSCertificateDelegation, err error)
	// Get retrieves the TLSCertificateDelegation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.TLSCertificateDelegation, error)
	TLSCertificateDelegationNamespaceListerExpansion
}

// tLSCertificateDelegationNamespaceLister implements the TLSCertificateDelegationNamespaceLister
// interface.
type tLSCertificateDelegationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TLSCertificateDelegations in the indexer for a given namespace.
func (s tLSCertificateDelegationNamespaceLister) List(selector labels.Selector) (ret []*v1.TLSCertificateDelegation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.TLSCertificateDelegation))
	})
	return ret, err
}

// Get retrieves the TLSCertificateDelegation from the indexer for a given namespace and name.
func (s tLSCertificateDelegationNamespaceLister) Get(name string) (*v1.TLSCertificateDelegation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("tlscertificatedelegation"), name)
	}
	return obj.(*v1.TLSCertificateDelegation), nil
}
