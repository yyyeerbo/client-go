/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	rbac_v1alpha1 "k8s.io/api/rbac/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "github.com/hyperhq/client-go/informers/internalinterfaces"
	kubernetes "github.com/hyperhq/client-go/kubernetes"
	v1alpha1 "github.com/hyperhq/client-go/listers/rbac/v1alpha1"
	cache "github.com/hyperhq/client-go/tools/cache"
	time "time"
)

// RoleBindingInformer provides access to a shared informer and lister for
// RoleBindings.
type RoleBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.RoleBindingLister
}

type roleBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRoleBindingInformer constructs a new informer for RoleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRoleBindingInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRoleBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRoleBindingInformer constructs a new informer for RoleBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRoleBindingInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RbacV1alpha1().RoleBindings(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RbacV1alpha1().RoleBindings(namespace).Watch(options)
			},
		},
		&rbac_v1alpha1.RoleBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *roleBindingInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRoleBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *roleBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rbac_v1alpha1.RoleBinding{}, f.defaultInformer)
}

func (f *roleBindingInformer) Lister() v1alpha1.RoleBindingLister {
	return v1alpha1.NewRoleBindingLister(f.Informer().GetIndexer())
}
