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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/baetyl/baetyl-cloud/plugin/kube/apis/cloud/v1alpha1"
	scheme "github.com/baetyl/baetyl-cloud/plugin/kube/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NodeDesiresGetter has a method to return a NodeDesireInterface.
// A group's client should implement this interface.
type NodeDesiresGetter interface {
	NodeDesires(namespace string) NodeDesireInterface
}

// NodeDesireInterface has methods to work with NodeDesire resources.
type NodeDesireInterface interface {
	Create(*v1alpha1.NodeDesire) (*v1alpha1.NodeDesire, error)
	Update(*v1alpha1.NodeDesire) (*v1alpha1.NodeDesire, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NodeDesire, error)
	List(opts v1.ListOptions) (*v1alpha1.NodeDesireList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodeDesire, err error)
	NodeDesireExpansion
}

// nodeDesires implements NodeDesireInterface
type nodeDesires struct {
	client rest.Interface
	ns     string
}

// newNodeDesires returns a NodeDesires
func newNodeDesires(c *CloudV1alpha1Client, namespace string) *nodeDesires {
	return &nodeDesires{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the nodeDesire, and returns the corresponding nodeDesire object, and an error if there is any.
func (c *nodeDesires) Get(name string, options v1.GetOptions) (result *v1alpha1.NodeDesire, err error) {
	result = &v1alpha1.NodeDesire{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodedesires").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodeDesires that match those selectors.
func (c *nodeDesires) List(opts v1.ListOptions) (result *v1alpha1.NodeDesireList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NodeDesireList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodedesires").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodeDesires.
func (c *nodeDesires) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("nodedesires").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a nodeDesire and creates it.  Returns the server's representation of the nodeDesire, and an error, if there is any.
func (c *nodeDesires) Create(nodeDesire *v1alpha1.NodeDesire) (result *v1alpha1.NodeDesire, err error) {
	result = &v1alpha1.NodeDesire{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("nodedesires").
		Body(nodeDesire).
		Do().
		Into(result)
	return
}

// Update takes the representation of a nodeDesire and updates it. Returns the server's representation of the nodeDesire, and an error, if there is any.
func (c *nodeDesires) Update(nodeDesire *v1alpha1.NodeDesire) (result *v1alpha1.NodeDesire, err error) {
	result = &v1alpha1.NodeDesire{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nodedesires").
		Name(nodeDesire.Name).
		Body(nodeDesire).
		Do().
		Into(result)
	return
}

// Delete takes name of the nodeDesire and deletes it. Returns an error if one occurs.
func (c *nodeDesires) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodedesires").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodeDesires) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodedesires").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched nodeDesire.
func (c *nodeDesires) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodeDesire, err error) {
	result = &v1alpha1.NodeDesire{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("nodedesires").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}