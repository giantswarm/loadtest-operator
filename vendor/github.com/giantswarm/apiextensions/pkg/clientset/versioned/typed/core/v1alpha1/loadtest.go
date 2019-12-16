/*
Copyright 2019 Giant Swarm GmbH.

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

	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
	scheme "github.com/giantswarm/apiextensions/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LoadtestsGetter has a method to return a LoadtestInterface.
// A group's client should implement this interface.
type LoadtestsGetter interface {
	Loadtests(namespace string) LoadtestInterface
}

// LoadtestInterface has methods to work with Loadtest resources.
type LoadtestInterface interface {
	Create(*v1alpha1.Loadtest) (*v1alpha1.Loadtest, error)
	Update(*v1alpha1.Loadtest) (*v1alpha1.Loadtest, error)
	UpdateStatus(*v1alpha1.Loadtest) (*v1alpha1.Loadtest, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Loadtest, error)
	List(opts v1.ListOptions) (*v1alpha1.LoadtestList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Loadtest, err error)
	LoadtestExpansion
}

// loadtests implements LoadtestInterface
type loadtests struct {
	client rest.Interface
	ns     string
}

// newLoadtests returns a Loadtests
func newLoadtests(c *CoreV1alpha1Client, namespace string) *loadtests {
	return &loadtests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the loadtest, and returns the corresponding loadtest object, and an error if there is any.
func (c *loadtests) Get(name string, options v1.GetOptions) (result *v1alpha1.Loadtest, err error) {
	result = &v1alpha1.Loadtest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loadtests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Loadtests that match those selectors.
func (c *loadtests) List(opts v1.ListOptions) (result *v1alpha1.LoadtestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LoadtestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("loadtests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested loadtests.
func (c *loadtests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("loadtests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a loadtest and creates it.  Returns the server's representation of the loadtest, and an error, if there is any.
func (c *loadtests) Create(loadtest *v1alpha1.Loadtest) (result *v1alpha1.Loadtest, err error) {
	result = &v1alpha1.Loadtest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("loadtests").
		Body(loadtest).
		Do().
		Into(result)
	return
}

// Update takes the representation of a loadtest and updates it. Returns the server's representation of the loadtest, and an error, if there is any.
func (c *loadtests) Update(loadtest *v1alpha1.Loadtest) (result *v1alpha1.Loadtest, err error) {
	result = &v1alpha1.Loadtest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("loadtests").
		Name(loadtest.Name).
		Body(loadtest).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *loadtests) UpdateStatus(loadtest *v1alpha1.Loadtest) (result *v1alpha1.Loadtest, err error) {
	result = &v1alpha1.Loadtest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("loadtests").
		Name(loadtest.Name).
		SubResource("status").
		Body(loadtest).
		Do().
		Into(result)
	return
}

// Delete takes name of the loadtest and deletes it. Returns an error if one occurs.
func (c *loadtests) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loadtests").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *loadtests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("loadtests").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched loadtest.
func (c *loadtests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Loadtest, err error) {
	result = &v1alpha1.Loadtest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("loadtests").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
