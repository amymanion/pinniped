/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/suzerain-io/pinniped/kubernetes/1.19/api/apis/crdpinniped/v1alpha1"
)

// FakePinnipedDiscoveryInfos implements PinnipedDiscoveryInfoInterface
type FakePinnipedDiscoveryInfos struct {
	Fake *FakeCrdV1alpha1
	ns   string
}

var pinnipeddiscoveryinfosResource = schema.GroupVersionResource{Group: "crd.pinniped.dev", Version: "v1alpha1", Resource: "pinnipeddiscoveryinfos"}

var pinnipeddiscoveryinfosKind = schema.GroupVersionKind{Group: "crd.pinniped.dev", Version: "v1alpha1", Kind: "PinnipedDiscoveryInfo"}

// Get takes name of the pinnipedDiscoveryInfo, and returns the corresponding pinnipedDiscoveryInfo object, and an error if there is any.
func (c *FakePinnipedDiscoveryInfos) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PinnipedDiscoveryInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pinnipeddiscoveryinfosResource, c.ns, name), &v1alpha1.PinnipedDiscoveryInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinnipedDiscoveryInfo), err
}

// List takes label and field selectors, and returns the list of PinnipedDiscoveryInfos that match those selectors.
func (c *FakePinnipedDiscoveryInfos) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PinnipedDiscoveryInfoList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pinnipeddiscoveryinfosResource, pinnipeddiscoveryinfosKind, c.ns, opts), &v1alpha1.PinnipedDiscoveryInfoList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PinnipedDiscoveryInfoList{ListMeta: obj.(*v1alpha1.PinnipedDiscoveryInfoList).ListMeta}
	for _, item := range obj.(*v1alpha1.PinnipedDiscoveryInfoList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pinnipedDiscoveryInfos.
func (c *FakePinnipedDiscoveryInfos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pinnipeddiscoveryinfosResource, c.ns, opts))

}

// Create takes the representation of a pinnipedDiscoveryInfo and creates it.  Returns the server's representation of the pinnipedDiscoveryInfo, and an error, if there is any.
func (c *FakePinnipedDiscoveryInfos) Create(ctx context.Context, pinnipedDiscoveryInfo *v1alpha1.PinnipedDiscoveryInfo, opts v1.CreateOptions) (result *v1alpha1.PinnipedDiscoveryInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pinnipeddiscoveryinfosResource, c.ns, pinnipedDiscoveryInfo), &v1alpha1.PinnipedDiscoveryInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinnipedDiscoveryInfo), err
}

// Update takes the representation of a pinnipedDiscoveryInfo and updates it. Returns the server's representation of the pinnipedDiscoveryInfo, and an error, if there is any.
func (c *FakePinnipedDiscoveryInfos) Update(ctx context.Context, pinnipedDiscoveryInfo *v1alpha1.PinnipedDiscoveryInfo, opts v1.UpdateOptions) (result *v1alpha1.PinnipedDiscoveryInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pinnipeddiscoveryinfosResource, c.ns, pinnipedDiscoveryInfo), &v1alpha1.PinnipedDiscoveryInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinnipedDiscoveryInfo), err
}

// Delete takes name of the pinnipedDiscoveryInfo and deletes it. Returns an error if one occurs.
func (c *FakePinnipedDiscoveryInfos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pinnipeddiscoveryinfosResource, c.ns, name), &v1alpha1.PinnipedDiscoveryInfo{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePinnipedDiscoveryInfos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pinnipeddiscoveryinfosResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PinnipedDiscoveryInfoList{})
	return err
}

// Patch applies the patch and returns the patched pinnipedDiscoveryInfo.
func (c *FakePinnipedDiscoveryInfos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PinnipedDiscoveryInfo, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pinnipeddiscoveryinfosResource, c.ns, name, pt, data, subresources...), &v1alpha1.PinnipedDiscoveryInfo{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PinnipedDiscoveryInfo), err
}
