// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package drivecrd

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Drive.
func (in *Drive) DeepCopy() *Drive {
	if in == nil {
		return nil
	}
	out := new(Drive)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Drive) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriveList) DeepCopyInto(out *DriveList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Drive, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriveList.
func (in *DriveList) DeepCopy() *DriveList {
	if in == nil {
		return nil
	}
	out := new(DriveList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DriveList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}