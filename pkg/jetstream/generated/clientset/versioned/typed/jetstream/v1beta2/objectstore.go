// Copyright 2025 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1beta2

import (
	context "context"

	jetstreamv1beta2 "github.com/nats-io/nack/pkg/jetstream/apis/jetstream/v1beta2"
	applyconfigurationjetstreamv1beta2 "github.com/nats-io/nack/pkg/jetstream/generated/applyconfiguration/jetstream/v1beta2"
	scheme "github.com/nats-io/nack/pkg/jetstream/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ObjectStoresGetter has a method to return a ObjectStoreInterface.
// A group's client should implement this interface.
type ObjectStoresGetter interface {
	ObjectStores(namespace string) ObjectStoreInterface
}

// ObjectStoreInterface has methods to work with ObjectStore resources.
type ObjectStoreInterface interface {
	Create(ctx context.Context, objectStore *jetstreamv1beta2.ObjectStore, opts v1.CreateOptions) (*jetstreamv1beta2.ObjectStore, error)
	Update(ctx context.Context, objectStore *jetstreamv1beta2.ObjectStore, opts v1.UpdateOptions) (*jetstreamv1beta2.ObjectStore, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, objectStore *jetstreamv1beta2.ObjectStore, opts v1.UpdateOptions) (*jetstreamv1beta2.ObjectStore, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*jetstreamv1beta2.ObjectStore, error)
	List(ctx context.Context, opts v1.ListOptions) (*jetstreamv1beta2.ObjectStoreList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *jetstreamv1beta2.ObjectStore, err error)
	Apply(ctx context.Context, objectStore *applyconfigurationjetstreamv1beta2.ObjectStoreApplyConfiguration, opts v1.ApplyOptions) (result *jetstreamv1beta2.ObjectStore, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, objectStore *applyconfigurationjetstreamv1beta2.ObjectStoreApplyConfiguration, opts v1.ApplyOptions) (result *jetstreamv1beta2.ObjectStore, err error)
	ObjectStoreExpansion
}

// objectStores implements ObjectStoreInterface
type objectStores struct {
	*gentype.ClientWithListAndApply[*jetstreamv1beta2.ObjectStore, *jetstreamv1beta2.ObjectStoreList, *applyconfigurationjetstreamv1beta2.ObjectStoreApplyConfiguration]
}

// newObjectStores returns a ObjectStores
func newObjectStores(c *JetstreamV1beta2Client, namespace string) *objectStores {
	return &objectStores{
		gentype.NewClientWithListAndApply[*jetstreamv1beta2.ObjectStore, *jetstreamv1beta2.ObjectStoreList, *applyconfigurationjetstreamv1beta2.ObjectStoreApplyConfiguration](
			"objectstores",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *jetstreamv1beta2.ObjectStore { return &jetstreamv1beta2.ObjectStore{} },
			func() *jetstreamv1beta2.ObjectStoreList { return &jetstreamv1beta2.ObjectStoreList{} },
		),
	}
}
