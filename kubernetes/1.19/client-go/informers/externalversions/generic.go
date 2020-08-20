/*
Copyright 2020 VMware, Inc.
SPDX-License-Identifier: Apache-2.0
*/

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"

	v1alpha1 "github.com/suzerain-io/pinniped/kubernetes/1.19/api/apis/crdpinniped/v1alpha1"
	pinnipedv1alpha1 "github.com/suzerain-io/pinniped/kubernetes/1.19/api/apis/pinniped/v1alpha1"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=crd.pinniped.dev, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("pinnipeddiscoveryinfos"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Crd().V1alpha1().PinnipedDiscoveryInfos().Informer()}, nil

		// Group=pinniped.dev, Version=v1alpha1
	case pinnipedv1alpha1.SchemeGroupVersion.WithResource("credentialrequests"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Pinniped().V1alpha1().CredentialRequests().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
