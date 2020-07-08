// SPDX-License-Identifier: Apache-2.0
// Copyright(c) 2018-2019 Saaras Inc.

/*
Copyright 2019  Heptio

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/saarasio/enroute/enroute-dp/apis/enroute/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGlobalConfigs implements GlobalConfigInterface
type FakeGlobalConfigs struct {
	Fake *FakeEnrouteV1beta1
	ns   string
}

var globalconfigsResource = schema.GroupVersionResource{Group: "enroute.saaras.io", Version: "v1beta1", Resource: "globalconfigs"}

var globalconfigsKind = schema.GroupVersionKind{Group: "enroute.saaras.io", Version: "v1beta1", Kind: "GlobalConfig"}

// Get takes name of the globalConfig, and returns the corresponding globalConfig object, and an error if there is any.
func (c *FakeGlobalConfigs) Get(name string, options v1.GetOptions) (result *v1beta1.GlobalConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(globalconfigsResource, c.ns, name), &v1beta1.GlobalConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.GlobalConfig), err
}

// List takes label and field selectors, and returns the list of GlobalConfigs that match those selectors.
func (c *FakeGlobalConfigs) List(opts v1.ListOptions) (result *v1beta1.GlobalConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(globalconfigsResource, globalconfigsKind, c.ns, opts), &v1beta1.GlobalConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.GlobalConfigList{ListMeta: obj.(*v1beta1.GlobalConfigList).ListMeta}
	for _, item := range obj.(*v1beta1.GlobalConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested globalConfigs.
func (c *FakeGlobalConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(globalconfigsResource, c.ns, opts))

}

// Create takes the representation of a globalConfig and creates it.  Returns the server's representation of the globalConfig, and an error, if there is any.
func (c *FakeGlobalConfigs) Create(globalConfig *v1beta1.GlobalConfig) (result *v1beta1.GlobalConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(globalconfigsResource, c.ns, globalConfig), &v1beta1.GlobalConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.GlobalConfig), err
}

// Update takes the representation of a globalConfig and updates it. Returns the server's representation of the globalConfig, and an error, if there is any.
func (c *FakeGlobalConfigs) Update(globalConfig *v1beta1.GlobalConfig) (result *v1beta1.GlobalConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(globalconfigsResource, c.ns, globalConfig), &v1beta1.GlobalConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.GlobalConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGlobalConfigs) UpdateStatus(globalConfig *v1beta1.GlobalConfig) (*v1beta1.GlobalConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(globalconfigsResource, "status", c.ns, globalConfig), &v1beta1.GlobalConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.GlobalConfig), err
}

// Delete takes name of the globalConfig and deletes it. Returns an error if one occurs.
func (c *FakeGlobalConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(globalconfigsResource, c.ns, name), &v1beta1.GlobalConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGlobalConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(globalconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.GlobalConfigList{})
	return err
}

// Patch applies the patch and returns the patched globalConfig.
func (c *FakeGlobalConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.GlobalConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(globalconfigsResource, c.ns, name, pt, data, subresources...), &v1beta1.GlobalConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.GlobalConfig), err
}