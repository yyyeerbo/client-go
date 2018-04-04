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

// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "k8s.io/api/storage/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"github.com/hyperhq/client-go/tools/cache"
)

// VolumeAttachmentLister helps list VolumeAttachments.
type VolumeAttachmentLister interface {
	// List lists all VolumeAttachments in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.VolumeAttachment, err error)
	// Get retrieves the VolumeAttachment from the index for a given name.
	Get(name string) (*v1alpha1.VolumeAttachment, error)
	VolumeAttachmentListerExpansion
}

// volumeAttachmentLister implements the VolumeAttachmentLister interface.
type volumeAttachmentLister struct {
	indexer cache.Indexer
}

// NewVolumeAttachmentLister returns a new VolumeAttachmentLister.
func NewVolumeAttachmentLister(indexer cache.Indexer) VolumeAttachmentLister {
	return &volumeAttachmentLister{indexer: indexer}
}

// List lists all VolumeAttachments in the indexer.
func (s *volumeAttachmentLister) List(selector labels.Selector) (ret []*v1alpha1.VolumeAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VolumeAttachment))
	})
	return ret, err
}

// Get retrieves the VolumeAttachment from the index for a given name.
func (s *volumeAttachmentLister) Get(name string) (*v1alpha1.VolumeAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("volumeattachment"), name)
	}
	return obj.(*v1alpha1.VolumeAttachment), nil
}
