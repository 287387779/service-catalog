/*
Copyright 2017 The Kubernetes Authors.

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

package tpr

import (
	"context"
	"testing"

	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	_ "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/install"
	"github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	namespace = "testns"
	name      = "testthing"
)

func TestCreate(t *testing.T) {
	broker := &v1alpha1.Broker{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1alpha1",
			Kind:       ServiceBrokerKind.String(),
		},
		ObjectMeta: metav1.ObjectMeta{Name: name},
	}
	outBroker := &v1alpha1.Broker{}
	keyer := Keyer{
		DefaultNamespace: namespace,
		ResourceName:     ServiceBrokerKind.String(),
		Separator:        "/",
	}
	iface := &store{
		decodeKey:    keyer.NamespaceAndNameFromKey,
		codec:        &fakeJSONCodec{},
		cl:           newFakeCoreRESTClient(),
		singularKind: ServiceBrokerKind,
	}
	if err := iface.Create(
		context.Background(),
		namespace+keyer.Separator+name,
		broker,
		outBroker,
		uint64(0),
	); err != nil {
		t.Fatalf("error on create (%s)", err)
	}
}

func TestRemoveNamespace(t *testing.T) {
	obj := &servicecatalog.ServiceClass{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "testns",
		},
	}
	if err := removeNamespace(obj); err != nil {
		t.Fatalf("couldn't remove namespace (%s", err)
	}
	if obj.Namespace != "" {
		t.Fatalf("couldn't remove namespace from object. it is still %s", obj.Namespace)
	}
}
