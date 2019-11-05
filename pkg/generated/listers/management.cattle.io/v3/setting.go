/*
Copyright 2019 Rancher Labs.

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

// Code generated by main. DO NOT EDIT.

package v3

import (
	v3 "github.com/rancher/rio/pkg/apis/management.cattle.io/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SettingLister helps list Settings.
type SettingLister interface {
	// List lists all Settings in the indexer.
	List(selector labels.Selector) (ret []*v3.Setting, err error)
	// Get retrieves the Setting from the index for a given name.
	Get(name string) (*v3.Setting, error)
	SettingListerExpansion
}

// settingLister implements the SettingLister interface.
type settingLister struct {
	indexer cache.Indexer
}

// NewSettingLister returns a new SettingLister.
func NewSettingLister(indexer cache.Indexer) SettingLister {
	return &settingLister{indexer: indexer}
}

// List lists all Settings in the indexer.
func (s *settingLister) List(selector labels.Selector) (ret []*v3.Setting, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.Setting))
	})
	return ret, err
}

// Get retrieves the Setting from the index for a given name.
func (s *settingLister) Get(name string) (*v3.Setting, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v3.Resource("setting"), name)
	}
	return obj.(*v3.Setting), nil
}
