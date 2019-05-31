// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Compaction) DeepCopyInto(out *Compaction) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Compaction.
func (in *Compaction) DeepCopy() *Compaction {
	if in == nil {
		return nil
	}
	out := new(Compaction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DistributedLogProtocol) DeepCopyInto(out *DistributedLogProtocol) {
	*out = *in
	in.PersistentProtocol.DeepCopyInto(&out.PersistentProtocol)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DistributedLogProtocol.
func (in *DistributedLogProtocol) DeepCopy() *DistributedLogProtocol {
	if in == nil {
		return nil
	}
	out := new(DistributedLogProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Partition) DeepCopyInto(out *Partition) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Partition.
func (in *Partition) DeepCopy() *Partition {
	if in == nil {
		return nil
	}
	out := new(Partition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Partition) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionGroup) DeepCopyInto(out *PartitionGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionGroup.
func (in *PartitionGroup) DeepCopy() *PartitionGroup {
	if in == nil {
		return nil
	}
	out := new(PartitionGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PartitionGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionGroupList) DeepCopyInto(out *PartitionGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PartitionGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionGroupList.
func (in *PartitionGroupList) DeepCopy() *PartitionGroupList {
	if in == nil {
		return nil
	}
	out := new(PartitionGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PartitionGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionGroupSpec) DeepCopyInto(out *PartitionGroupSpec) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Raft != nil {
		in, out := &in.Raft, &out.Raft
		*out = new(RaftProtocol)
		(*in).DeepCopyInto(*out)
	}
	if in.PrimaryBackup != nil {
		in, out := &in.PrimaryBackup, &out.PrimaryBackup
		*out = new(PrimaryBackupProtocol)
		**out = **in
	}
	if in.Log != nil {
		in, out := &in.Log, &out.Log
		*out = new(DistributedLogProtocol)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionGroupSpec.
func (in *PartitionGroupSpec) DeepCopy() *PartitionGroupSpec {
	if in == nil {
		return nil
	}
	out := new(PartitionGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionGroupStatus) DeepCopyInto(out *PartitionGroupStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionGroupStatus.
func (in *PartitionGroupStatus) DeepCopy() *PartitionGroupStatus {
	if in == nil {
		return nil
	}
	out := new(PartitionGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionList) DeepCopyInto(out *PartitionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Partition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionList.
func (in *PartitionList) DeepCopy() *PartitionList {
	if in == nil {
		return nil
	}
	out := new(PartitionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PartitionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionSpec) DeepCopyInto(out *PartitionSpec) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Raft != nil {
		in, out := &in.Raft, &out.Raft
		*out = new(RaftProtocol)
		(*in).DeepCopyInto(*out)
	}
	if in.PrimaryBackup != nil {
		in, out := &in.PrimaryBackup, &out.PrimaryBackup
		*out = new(PrimaryBackupProtocol)
		**out = **in
	}
	if in.Log != nil {
		in, out := &in.Log, &out.Log
		*out = new(DistributedLogProtocol)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionSpec.
func (in *PartitionSpec) DeepCopy() *PartitionSpec {
	if in == nil {
		return nil
	}
	out := new(PartitionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PartitionStatus) DeepCopyInto(out *PartitionStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PartitionStatus.
func (in *PartitionStatus) DeepCopy() *PartitionStatus {
	if in == nil {
		return nil
	}
	out := new(PartitionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentProtocol) DeepCopyInto(out *PersistentProtocol) {
	*out = *in
	in.Storage.DeepCopyInto(&out.Storage)
	out.Compaction = in.Compaction
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentProtocol.
func (in *PersistentProtocol) DeepCopy() *PersistentProtocol {
	if in == nil {
		return nil
	}
	out := new(PersistentProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrimaryBackupProtocol) DeepCopyInto(out *PrimaryBackupProtocol) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrimaryBackupProtocol.
func (in *PrimaryBackupProtocol) DeepCopy() *PrimaryBackupProtocol {
	if in == nil {
		return nil
	}
	out := new(PrimaryBackupProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RaftProtocol) DeepCopyInto(out *RaftProtocol) {
	*out = *in
	in.PersistentProtocol.DeepCopyInto(&out.PersistentProtocol)
	if in.ElectionTimeout != nil {
		in, out := &in.ElectionTimeout, &out.ElectionTimeout
		*out = new(int64)
		**out = **in
	}
	if in.HeartbeatInterval != nil {
		in, out := &in.HeartbeatInterval, &out.HeartbeatInterval
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RaftProtocol.
func (in *RaftProtocol) DeepCopy() *RaftProtocol {
	if in == nil {
		return nil
	}
	out := new(RaftProtocol)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
	if in.ClassName != nil {
		in, out := &in.ClassName, &out.ClassName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}
