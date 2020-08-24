/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/suzerain-io/pinniped/generated/1.19/apis/crdpinniped/v1alpha1"
	scheme "github.com/suzerain-io/pinniped/generated/1.19/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PinnipedDiscoveryInfosGetter has a method to return a PinnipedDiscoveryInfoInterface.
// A group's client should implement this interface.
type PinnipedDiscoveryInfosGetter interface {
	PinnipedDiscoveryInfos(namespace string) PinnipedDiscoveryInfoInterface
}

// PinnipedDiscoveryInfoInterface has methods to work with PinnipedDiscoveryInfo resources.
type PinnipedDiscoveryInfoInterface interface {
	Create(ctx context.Context, pinnipedDiscoveryInfo *v1alpha1.PinnipedDiscoveryInfo, opts v1.CreateOptions) (*v1alpha1.PinnipedDiscoveryInfo, error)
	Update(ctx context.Context, pinnipedDiscoveryInfo *v1alpha1.PinnipedDiscoveryInfo, opts v1.UpdateOptions) (*v1alpha1.PinnipedDiscoveryInfo, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.PinnipedDiscoveryInfo, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PinnipedDiscoveryInfoList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PinnipedDiscoveryInfo, err error)
	PinnipedDiscoveryInfoExpansion
}

// pinnipedDiscoveryInfos implements PinnipedDiscoveryInfoInterface
type pinnipedDiscoveryInfos struct {
	client rest.Interface
	ns     string
}

// newPinnipedDiscoveryInfos returns a PinnipedDiscoveryInfos
func newPinnipedDiscoveryInfos(c *CrdV1alpha1Client, namespace string) *pinnipedDiscoveryInfos {
	return &pinnipedDiscoveryInfos{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the pinnipedDiscoveryInfo, and returns the corresponding pinnipedDiscoveryInfo object, and an error if there is any.
func (c *pinnipedDiscoveryInfos) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PinnipedDiscoveryInfo, err error) {
	result = &v1alpha1.PinnipedDiscoveryInfo{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pinnipeddiscoveryinfos").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PinnipedDiscoveryInfos that match those selectors.
func (c *pinnipedDiscoveryInfos) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PinnipedDiscoveryInfoList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PinnipedDiscoveryInfoList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pinnipeddiscoveryinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pinnipedDiscoveryInfos.
func (c *pinnipedDiscoveryInfos) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("pinnipeddiscoveryinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a pinnipedDiscoveryInfo and creates it.  Returns the server's representation of the pinnipedDiscoveryInfo, and an error, if there is any.
func (c *pinnipedDiscoveryInfos) Create(ctx context.Context, pinnipedDiscoveryInfo *v1alpha1.PinnipedDiscoveryInfo, opts v1.CreateOptions) (result *v1alpha1.PinnipedDiscoveryInfo, err error) {
	result = &v1alpha1.PinnipedDiscoveryInfo{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("pinnipeddiscoveryinfos").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pinnipedDiscoveryInfo).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a pinnipedDiscoveryInfo and updates it. Returns the server's representation of the pinnipedDiscoveryInfo, and an error, if there is any.
func (c *pinnipedDiscoveryInfos) Update(ctx context.Context, pinnipedDiscoveryInfo *v1alpha1.PinnipedDiscoveryInfo, opts v1.UpdateOptions) (result *v1alpha1.PinnipedDiscoveryInfo, err error) {
	result = &v1alpha1.PinnipedDiscoveryInfo{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pinnipeddiscoveryinfos").
		Name(pinnipedDiscoveryInfo.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(pinnipedDiscoveryInfo).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the pinnipedDiscoveryInfo and deletes it. Returns an error if one occurs.
func (c *pinnipedDiscoveryInfos) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pinnipeddiscoveryinfos").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pinnipedDiscoveryInfos) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pinnipeddiscoveryinfos").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched pinnipedDiscoveryInfo.
func (c *pinnipedDiscoveryInfos) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PinnipedDiscoveryInfo, err error) {
	result = &v1alpha1.PinnipedDiscoveryInfo{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("pinnipeddiscoveryinfos").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}