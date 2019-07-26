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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	versioned "github.com/nimbess/stargazer/pkg/client/clientset/versioned"
	internalinterfaces "github.com/nimbess/stargazer/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/nimbess/stargazer/pkg/client/listers/unp/v1"
	unpv1 "github.com/nimbess/stargazer/pkg/crd/api/unp/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// UnpConfigInformer provides access to a shared informer and lister for
// UnpConfigs.
type UnpConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.UnpConfigLister
}

type unpConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewUnpConfigInformer constructs a new informer for UnpConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewUnpConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredUnpConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredUnpConfigInformer constructs a new informer for UnpConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredUnpConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NimbessV1().UnpConfigs(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NimbessV1().UnpConfigs(namespace).Watch(options)
			},
		},
		&unpv1.UnpConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *unpConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredUnpConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *unpConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&unpv1.UnpConfig{}, f.defaultInformer)
}

func (f *unpConfigInformer) Lister() v1.UnpConfigLister {
	return v1.NewUnpConfigLister(f.Informer().GetIndexer())
}