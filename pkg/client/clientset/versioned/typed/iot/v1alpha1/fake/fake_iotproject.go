/*
 * Copyright 2018-2019, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/enmasseproject/enmasse/pkg/apis/iot/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIoTProjects implements IoTProjectInterface
type FakeIoTProjects struct {
	Fake *FakeIotV1alpha1
	ns   string
}

var iotprojectsResource = schema.GroupVersionResource{Group: "iot.enmasse.io", Version: "v1alpha1", Resource: "iotprojects"}

var iotprojectsKind = schema.GroupVersionKind{Group: "iot.enmasse.io", Version: "v1alpha1", Kind: "IoTProject"}

// Get takes name of the ioTProject, and returns the corresponding ioTProject object, and an error if there is any.
func (c *FakeIoTProjects) Get(name string, options v1.GetOptions) (result *v1alpha1.IoTProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(iotprojectsResource, c.ns, name), &v1alpha1.IoTProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IoTProject), err
}

// List takes label and field selectors, and returns the list of IoTProjects that match those selectors.
func (c *FakeIoTProjects) List(opts v1.ListOptions) (result *v1alpha1.IoTProjectList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(iotprojectsResource, iotprojectsKind, c.ns, opts), &v1alpha1.IoTProjectList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IoTProjectList{ListMeta: obj.(*v1alpha1.IoTProjectList).ListMeta}
	for _, item := range obj.(*v1alpha1.IoTProjectList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ioTProjects.
func (c *FakeIoTProjects) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(iotprojectsResource, c.ns, opts))

}

// Create takes the representation of a ioTProject and creates it.  Returns the server's representation of the ioTProject, and an error, if there is any.
func (c *FakeIoTProjects) Create(ioTProject *v1alpha1.IoTProject) (result *v1alpha1.IoTProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(iotprojectsResource, c.ns, ioTProject), &v1alpha1.IoTProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IoTProject), err
}

// Update takes the representation of a ioTProject and updates it. Returns the server's representation of the ioTProject, and an error, if there is any.
func (c *FakeIoTProjects) Update(ioTProject *v1alpha1.IoTProject) (result *v1alpha1.IoTProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(iotprojectsResource, c.ns, ioTProject), &v1alpha1.IoTProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IoTProject), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIoTProjects) UpdateStatus(ioTProject *v1alpha1.IoTProject) (*v1alpha1.IoTProject, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(iotprojectsResource, "status", c.ns, ioTProject), &v1alpha1.IoTProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IoTProject), err
}

// Delete takes name of the ioTProject and deletes it. Returns an error if one occurs.
func (c *FakeIoTProjects) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(iotprojectsResource, c.ns, name), &v1alpha1.IoTProject{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIoTProjects) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(iotprojectsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IoTProjectList{})
	return err
}

// Patch applies the patch and returns the patched ioTProject.
func (c *FakeIoTProjects) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.IoTProject, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(iotprojectsResource, c.ns, name, pt, data, subresources...), &v1alpha1.IoTProject{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IoTProject), err
}
