// SPDX-License-Identifier: Apache-2.0
// Copyright(c) 2018-2020 Saaras Inc.

// Copyright © 2018 Heptio
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package dag provides a data model, in the form of a directed acyclic graph,
// of the relationship between Kubernetes Ingress, Service, and Secret objects.
package dag

import (
	"sync"

	v1 "k8s.io/api/core/v1"
	net_v1 "k8s.io/api/networking/v1"
	"k8s.io/client-go/tools/cache"

	gatewayhostv1 "github.com/saarasio/enroute/enroute-dp/apis/enroute/v1"
	"github.com/sirupsen/logrus"
)

// A KubernetesCache holds Kubernetes objects and associated configuration and produces
// DAG values.
type KubernetesCache struct {
	// GatewayHostRootNamespaces specifies the namespaces where root
	// GatewayHosts can be defined. If empty, roots can be defined in any
	// namespace.
	GatewayHostRootNamespaces []string

	mu sync.RWMutex
	logrus.FieldLogger

	ingresses    map[Meta]*net_v1.Ingress
	gatewayhosts map[Meta]*gatewayhostv1.GatewayHost
	secrets      map[Meta]*v1.Secret
	delegations  map[Meta]*gatewayhostv1.TLSCertificateDelegation
	services     map[Meta]*v1.Service

	routefilters map[RouteFilterMeta]*gatewayhostv1.RouteFilter
	httpfilters  map[HttpFilterMeta]*gatewayhostv1.HttpFilter
}

// Meta holds the name and namespace of a Kubernetes object.
type Meta struct {
	name, namespace string
}

type RouteFilterMeta struct {
	filter_type, name, namespace string
}

type HttpFilterMeta struct {
	filter_type, name, namespace string
}

// Insert inserts obj into the KubernetesCache.
// If an object with a matching type, name, and namespace exists, it will be overwritten.
func (kc *KubernetesCache) Insert(obj interface{}) {
	kc.mu.Lock()
	defer kc.mu.Unlock()
	switch obj := obj.(type) {
	case *v1.Secret:
		m := Meta{name: obj.Name, namespace: obj.Namespace}
		if kc.secrets == nil {
			kc.secrets = make(map[Meta]*v1.Secret)
		}
		kc.secrets[m] = obj
	case *v1.Service:
		m := Meta{name: obj.Name, namespace: obj.Namespace}
		if kc.services == nil {
			kc.services = make(map[Meta]*v1.Service)
		}
		kc.services[m] = obj
	case *net_v1.Ingress:
		m := Meta{name: obj.Name, namespace: obj.Namespace}
		if kc.ingresses == nil {
			kc.ingresses = make(map[Meta]*net_v1.Ingress)
		}
		kc.ingresses[m] = obj
	case *gatewayhostv1.GatewayHost:
		m := Meta{name: obj.Name, namespace: obj.Namespace}
		if kc.gatewayhosts == nil {
			kc.gatewayhosts = make(map[Meta]*gatewayhostv1.GatewayHost)
		}
		kc.gatewayhosts[m] = obj
	case *gatewayhostv1.TLSCertificateDelegation:
		m := Meta{name: obj.Name, namespace: obj.Namespace}
		if kc.delegations == nil {
			kc.delegations = make(map[Meta]*gatewayhostv1.TLSCertificateDelegation)
		}
		kc.delegations[m] = obj

	case *gatewayhostv1.HttpFilter:
		m := HttpFilterMeta{filter_type: obj.Spec.Type, name: obj.Name, namespace: obj.Namespace}
		if kc.httpfilters == nil {
			kc.httpfilters = make(map[HttpFilterMeta]*gatewayhostv1.HttpFilter)
		}
		kc.httpfilters[m] = obj

	case *gatewayhostv1.RouteFilter:
		m := RouteFilterMeta{filter_type: obj.Spec.Type, name: obj.Name, namespace: obj.Namespace}
		if kc.routefilters == nil {
			kc.routefilters = make(map[RouteFilterMeta]*gatewayhostv1.RouteFilter)
		}
		kc.routefilters[m] = obj

	default:
		// not an interesting object
	}
}

// Remove removes obj from the KubernetesCache.
// If no object with a matching type, name, and namespace exists in the DAG, no action is taken.
func (kc *KubernetesCache) Remove(obj interface{}) {
	switch obj := obj.(type) {
	default:
		kc.remove(obj)
	case cache.DeletedFinalStateUnknown:
		kc.Remove(obj.Obj) // recurse into ourselves with the tombstoned value
	}
}

func (kc *KubernetesCache) remove(obj interface{}) {
	kc.mu.Lock()
	defer kc.mu.Unlock()
	switch obj := obj.(type) {
	case *v1.Secret:
		m := Meta{name: obj.Name, namespace: obj.Namespace}
		delete(kc.secrets, m)
	case *v1.Service:
		m := Meta{name: obj.Name, namespace: obj.Namespace}
		delete(kc.services, m)
	case *net_v1.Ingress:
		m := Meta{name: obj.Name, namespace: obj.Namespace}
		delete(kc.ingresses, m)
	case *gatewayhostv1.GatewayHost:
		m := Meta{name: obj.Name, namespace: obj.Namespace}
		delete(kc.gatewayhosts, m)
	case *gatewayhostv1.TLSCertificateDelegation:
		m := Meta{name: obj.Name, namespace: obj.Namespace}
		delete(kc.delegations, m)

	case *gatewayhostv1.HttpFilter:
		m := HttpFilterMeta{filter_type: obj.Spec.Type, name: obj.Name, namespace: obj.Namespace}
		delete(kc.httpfilters, m)

	case *gatewayhostv1.RouteFilter:
		m := RouteFilterMeta{filter_type: obj.Spec.Type, name: obj.Name, namespace: obj.Namespace}
		delete(kc.routefilters, m)
	default:
		// not interesting
	}
}
