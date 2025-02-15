// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/gocrane/api/prediction/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodePredictions implements NodePredictionInterface
type FakeNodePredictions struct {
	Fake *FakePredictionV1alpha1
}

var nodepredictionsResource = schema.GroupVersionResource{Group: "prediction.crane.io", Version: "v1alpha1", Resource: "nodepredictions"}

var nodepredictionsKind = schema.GroupVersionKind{Group: "prediction.crane.io", Version: "v1alpha1", Kind: "NodePrediction"}

// Get takes name of the nodePrediction, and returns the corresponding nodePrediction object, and an error if there is any.
func (c *FakeNodePredictions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NodePrediction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(nodepredictionsResource, name), &v1alpha1.NodePrediction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodePrediction), err
}

// List takes label and field selectors, and returns the list of NodePredictions that match those selectors.
func (c *FakeNodePredictions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NodePredictionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(nodepredictionsResource, nodepredictionsKind, opts), &v1alpha1.NodePredictionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodePredictionList{ListMeta: obj.(*v1alpha1.NodePredictionList).ListMeta}
	for _, item := range obj.(*v1alpha1.NodePredictionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodePredictions.
func (c *FakeNodePredictions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(nodepredictionsResource, opts))
}

// Create takes the representation of a nodePrediction and creates it.  Returns the server's representation of the nodePrediction, and an error, if there is any.
func (c *FakeNodePredictions) Create(ctx context.Context, nodePrediction *v1alpha1.NodePrediction, opts v1.CreateOptions) (result *v1alpha1.NodePrediction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(nodepredictionsResource, nodePrediction), &v1alpha1.NodePrediction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodePrediction), err
}

// Update takes the representation of a nodePrediction and updates it. Returns the server's representation of the nodePrediction, and an error, if there is any.
func (c *FakeNodePredictions) Update(ctx context.Context, nodePrediction *v1alpha1.NodePrediction, opts v1.UpdateOptions) (result *v1alpha1.NodePrediction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(nodepredictionsResource, nodePrediction), &v1alpha1.NodePrediction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodePrediction), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodePredictions) UpdateStatus(ctx context.Context, nodePrediction *v1alpha1.NodePrediction, opts v1.UpdateOptions) (*v1alpha1.NodePrediction, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(nodepredictionsResource, "status", nodePrediction), &v1alpha1.NodePrediction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodePrediction), err
}

// Delete takes name of the nodePrediction and deletes it. Returns an error if one occurs.
func (c *FakeNodePredictions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(nodepredictionsResource, name), &v1alpha1.NodePrediction{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodePredictions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(nodepredictionsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodePredictionList{})
	return err
}

// Patch applies the patch and returns the patched nodePrediction.
func (c *FakeNodePredictions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NodePrediction, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(nodepredictionsResource, name, pt, data, subresources...), &v1alpha1.NodePrediction{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodePrediction), err
}
