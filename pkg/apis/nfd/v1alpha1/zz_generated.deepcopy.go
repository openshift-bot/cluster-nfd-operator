// +build !ignore_autogenerated

// Code generated by operator-sdk-v0.12.0-x86_64-linux-gnu. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFeatureDiscovery) DeepCopyInto(out *NodeFeatureDiscovery) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFeatureDiscovery.
func (in *NodeFeatureDiscovery) DeepCopy() *NodeFeatureDiscovery {
	if in == nil {
		return nil
	}
	out := new(NodeFeatureDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeFeatureDiscovery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFeatureDiscoveryList) DeepCopyInto(out *NodeFeatureDiscoveryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeFeatureDiscovery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFeatureDiscoveryList.
func (in *NodeFeatureDiscoveryList) DeepCopy() *NodeFeatureDiscoveryList {
	if in == nil {
		return nil
	}
	out := new(NodeFeatureDiscoveryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeFeatureDiscoveryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFeatureDiscoverySpec) DeepCopyInto(out *NodeFeatureDiscoverySpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFeatureDiscoverySpec.
func (in *NodeFeatureDiscoverySpec) DeepCopy() *NodeFeatureDiscoverySpec {
	if in == nil {
		return nil
	}
	out := new(NodeFeatureDiscoverySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeFeatureDiscoveryStatus) DeepCopyInto(out *NodeFeatureDiscoveryStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeFeatureDiscoveryStatus.
func (in *NodeFeatureDiscoveryStatus) DeepCopy() *NodeFeatureDiscoveryStatus {
	if in == nil {
		return nil
	}
	out := new(NodeFeatureDiscoveryStatus)
	in.DeepCopyInto(out)
	return out
}
