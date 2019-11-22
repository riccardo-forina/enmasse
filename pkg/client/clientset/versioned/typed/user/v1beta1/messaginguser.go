/*
 * Copyright 2018-2019, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"time"

	v1beta1 "github.com/enmasseproject/enmasse/pkg/apis/user/v1beta1"
	scheme "github.com/enmasseproject/enmasse/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MessagingUsersGetter has a method to return a MessagingUserInterface.
// A group's client should implement this interface.
type MessagingUsersGetter interface {
	MessagingUsers(namespace string) MessagingUserInterface
}

// MessagingUserInterface has methods to work with MessagingUser resources.
type MessagingUserInterface interface {
	Create(*v1beta1.MessagingUser) (*v1beta1.MessagingUser, error)
	Update(*v1beta1.MessagingUser) (*v1beta1.MessagingUser, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.MessagingUser, error)
	List(opts v1.ListOptions) (*v1beta1.MessagingUserList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.MessagingUser, err error)
	MessagingUserExpansion
}

// messagingUsers implements MessagingUserInterface
type messagingUsers struct {
	client rest.Interface
	ns     string
}

// newMessagingUsers returns a MessagingUsers
func newMessagingUsers(c *UserV1beta1Client, namespace string) *messagingUsers {
	return &messagingUsers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the messagingUser, and returns the corresponding messagingUser object, and an error if there is any.
func (c *messagingUsers) Get(name string, options v1.GetOptions) (result *v1beta1.MessagingUser, err error) {
	result = &v1beta1.MessagingUser{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("messagingusers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MessagingUsers that match those selectors.
func (c *messagingUsers) List(opts v1.ListOptions) (result *v1beta1.MessagingUserList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.MessagingUserList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("messagingusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested messagingUsers.
func (c *messagingUsers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("messagingusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a messagingUser and creates it.  Returns the server's representation of the messagingUser, and an error, if there is any.
func (c *messagingUsers) Create(messagingUser *v1beta1.MessagingUser) (result *v1beta1.MessagingUser, err error) {
	result = &v1beta1.MessagingUser{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("messagingusers").
		Body(messagingUser).
		Do().
		Into(result)
	return
}

// Update takes the representation of a messagingUser and updates it. Returns the server's representation of the messagingUser, and an error, if there is any.
func (c *messagingUsers) Update(messagingUser *v1beta1.MessagingUser) (result *v1beta1.MessagingUser, err error) {
	result = &v1beta1.MessagingUser{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("messagingusers").
		Name(messagingUser.Name).
		Body(messagingUser).
		Do().
		Into(result)
	return
}

// Delete takes name of the messagingUser and deletes it. Returns an error if one occurs.
func (c *messagingUsers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("messagingusers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *messagingUsers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("messagingusers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched messagingUser.
func (c *messagingUsers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.MessagingUser, err error) {
	result = &v1beta1.MessagingUser{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("messagingusers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
