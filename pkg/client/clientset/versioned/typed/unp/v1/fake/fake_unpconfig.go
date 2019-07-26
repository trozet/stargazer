/*
Copyright The Kubernetes Authors.

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
	unpv1 "github.com/nimbess/stargazer/pkg/crd/api/unp/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUnpConfigs implements UnpConfigInterface
type FakeUnpConfigs struct {
	Fake *FakeNimbessV1
	ns   string
}

var unpconfigsResource = schema.GroupVersionResource{Group: "nimbess", Version: "v1", Resource: "unpconfigs"}

var unpconfigsKind = schema.GroupVersionKind{Group: "nimbess", Version: "v1", Kind: "UnpConfig"}

// Get takes name of the unpConfig, and returns the corresponding unpConfig object, and an error if there is any.
func (c *FakeUnpConfigs) Get(name string, options v1.GetOptions) (result *unpv1.UnpConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(unpconfigsResource, c.ns, name), &unpv1.UnpConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*unpv1.UnpConfig), err
}

// List takes label and field selectors, and returns the list of UnpConfigs that match those selectors.
func (c *FakeUnpConfigs) List(opts v1.ListOptions) (result *unpv1.UnpConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(unpconfigsResource, unpconfigsKind, c.ns, opts), &unpv1.UnpConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &unpv1.UnpConfigList{ListMeta: obj.(*unpv1.UnpConfigList).ListMeta}
	for _, item := range obj.(*unpv1.UnpConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested unpConfigs.
func (c *FakeUnpConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(unpconfigsResource, c.ns, opts))

}

// Create takes the representation of a unpConfig and creates it.  Returns the server's representation of the unpConfig, and an error, if there is any.
func (c *FakeUnpConfigs) Create(unpConfig *unpv1.UnpConfig) (result *unpv1.UnpConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(unpconfigsResource, c.ns, unpConfig), &unpv1.UnpConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*unpv1.UnpConfig), err
}

// Update takes the representation of a unpConfig and updates it. Returns the server's representation of the unpConfig, and an error, if there is any.
func (c *FakeUnpConfigs) Update(unpConfig *unpv1.UnpConfig) (result *unpv1.UnpConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(unpconfigsResource, c.ns, unpConfig), &unpv1.UnpConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*unpv1.UnpConfig), err
}

// Delete takes name of the unpConfig and deletes it. Returns an error if one occurs.
func (c *FakeUnpConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(unpconfigsResource, c.ns, name), &unpv1.UnpConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUnpConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(unpconfigsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &unpv1.UnpConfigList{})
	return err
}

// Patch applies the patch and returns the patched unpConfig.
func (c *FakeUnpConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *unpv1.UnpConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(unpconfigsResource, c.ns, name, pt, data, subresources...), &unpv1.UnpConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*unpv1.UnpConfig), err
}
