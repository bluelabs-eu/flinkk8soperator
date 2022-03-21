// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/lyft/flinkk8soperator/pkg/apis/app/v1beta1"
	scheme "github.com/lyft/flinkk8soperator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FlinkApplicationsGetter has a method to return a FlinkApplicationInterface.
// A group's client should implement this interface.
type FlinkApplicationsGetter interface {
	FlinkApplications(namespace string) FlinkApplicationInterface
}

// FlinkApplicationInterface has methods to work with FlinkApplication resources.
type FlinkApplicationInterface interface {
	Create(ctx context.Context, flinkApplication *v1beta1.FlinkApplication, opts v1.CreateOptions) (*v1beta1.FlinkApplication, error)
	Update(ctx context.Context, flinkApplication *v1beta1.FlinkApplication, opts v1.UpdateOptions) (*v1beta1.FlinkApplication, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.FlinkApplication, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.FlinkApplicationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FlinkApplication, err error)
	FlinkApplicationExpansion
}

// flinkApplications implements FlinkApplicationInterface
type flinkApplications struct {
	client rest.Interface
	ns     string
}

// newFlinkApplications returns a FlinkApplications
func newFlinkApplications(c *FlinkV1beta1Client, namespace string) *flinkApplications {
	return &flinkApplications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the flinkApplication, and returns the corresponding flinkApplication object, and an error if there is any.
func (c *flinkApplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.FlinkApplication, err error) {
	result = &v1beta1.FlinkApplication{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("flinkapplications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FlinkApplications that match those selectors.
func (c *flinkApplications) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.FlinkApplicationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.FlinkApplicationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("flinkapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested flinkApplications.
func (c *flinkApplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("flinkapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a flinkApplication and creates it.  Returns the server's representation of the flinkApplication, and an error, if there is any.
func (c *flinkApplications) Create(ctx context.Context, flinkApplication *v1beta1.FlinkApplication, opts v1.CreateOptions) (result *v1beta1.FlinkApplication, err error) {
	result = &v1beta1.FlinkApplication{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("flinkapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(flinkApplication).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a flinkApplication and updates it. Returns the server's representation of the flinkApplication, and an error, if there is any.
func (c *flinkApplications) Update(ctx context.Context, flinkApplication *v1beta1.FlinkApplication, opts v1.UpdateOptions) (result *v1beta1.FlinkApplication, err error) {
	result = &v1beta1.FlinkApplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("flinkapplications").
		Name(flinkApplication.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(flinkApplication).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the flinkApplication and deletes it. Returns an error if one occurs.
func (c *flinkApplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("flinkapplications").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *flinkApplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("flinkapplications").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched flinkApplication.
func (c *flinkApplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FlinkApplication, err error) {
	result = &v1beta1.FlinkApplication{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("flinkapplications").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
