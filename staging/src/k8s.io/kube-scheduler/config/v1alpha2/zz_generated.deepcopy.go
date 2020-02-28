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

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSchedulerConfiguration) DeepCopyInto(out *KubeSchedulerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.LeaderElection.DeepCopyInto(&out.LeaderElection)
	out.ClientConnection = in.ClientConnection
	if in.HealthzBindAddress != nil {
		in, out := &in.HealthzBindAddress, &out.HealthzBindAddress
		*out = new(string)
		**out = **in
	}
	if in.MetricsBindAddress != nil {
		in, out := &in.MetricsBindAddress, &out.MetricsBindAddress
		*out = new(string)
		**out = **in
	}
	in.DebuggingConfiguration.DeepCopyInto(&out.DebuggingConfiguration)
	if in.DisablePreemption != nil {
		in, out := &in.DisablePreemption, &out.DisablePreemption
		*out = new(bool)
		**out = **in
	}
	if in.PercentageOfNodesToScore != nil {
		in, out := &in.PercentageOfNodesToScore, &out.PercentageOfNodesToScore
		*out = new(int32)
		**out = **in
	}
	if in.BindTimeoutSeconds != nil {
		in, out := &in.BindTimeoutSeconds, &out.BindTimeoutSeconds
		*out = new(int64)
		**out = **in
	}
	if in.PodInitialBackoffSeconds != nil {
		in, out := &in.PodInitialBackoffSeconds, &out.PodInitialBackoffSeconds
		*out = new(int64)
		**out = **in
	}
	if in.PodMaxBackoffSeconds != nil {
		in, out := &in.PodMaxBackoffSeconds, &out.PodMaxBackoffSeconds
		*out = new(int64)
		**out = **in
	}
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make([]KubeSchedulerProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSchedulerConfiguration.
func (in *KubeSchedulerConfiguration) DeepCopy() *KubeSchedulerConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeSchedulerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeSchedulerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSchedulerLeaderElectionConfiguration) DeepCopyInto(out *KubeSchedulerLeaderElectionConfiguration) {
	*out = *in
	in.LeaderElectionConfiguration.DeepCopyInto(&out.LeaderElectionConfiguration)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSchedulerLeaderElectionConfiguration.
func (in *KubeSchedulerLeaderElectionConfiguration) DeepCopy() *KubeSchedulerLeaderElectionConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeSchedulerLeaderElectionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSchedulerProfile) DeepCopyInto(out *KubeSchedulerProfile) {
	*out = *in
	if in.SchedulerName != nil {
		in, out := &in.SchedulerName, &out.SchedulerName
		*out = new(string)
		**out = **in
	}
	if in.Plugins != nil {
		in, out := &in.Plugins, &out.Plugins
		*out = new(Plugins)
		(*in).DeepCopyInto(*out)
	}
	if in.PluginConfig != nil {
		in, out := &in.PluginConfig, &out.PluginConfig
		*out = make([]PluginConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSchedulerProfile.
func (in *KubeSchedulerProfile) DeepCopy() *KubeSchedulerProfile {
	if in == nil {
		return nil
	}
	out := new(KubeSchedulerProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugin) DeepCopyInto(out *Plugin) {
	*out = *in
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugin.
func (in *Plugin) DeepCopy() *Plugin {
	if in == nil {
		return nil
	}
	out := new(Plugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginConfig) DeepCopyInto(out *PluginConfig) {
	*out = *in
	in.Args.DeepCopyInto(&out.Args)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginConfig.
func (in *PluginConfig) DeepCopy() *PluginConfig {
	if in == nil {
		return nil
	}
	out := new(PluginConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginSet) DeepCopyInto(out *PluginSet) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = make([]Plugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = make([]Plugin, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginSet.
func (in *PluginSet) DeepCopy() *PluginSet {
	if in == nil {
		return nil
	}
	out := new(PluginSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugins) DeepCopyInto(out *Plugins) {
	*out = *in
	if in.QueueSort != nil {
		in, out := &in.QueueSort, &out.QueueSort
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.PreFilter != nil {
		in, out := &in.PreFilter, &out.PreFilter
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.Filter != nil {
		in, out := &in.Filter, &out.Filter
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.PreScore != nil {
		in, out := &in.PreScore, &out.PreScore
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.Score != nil {
		in, out := &in.Score, &out.Score
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.Reserve != nil {
		in, out := &in.Reserve, &out.Reserve
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.Permit != nil {
		in, out := &in.Permit, &out.Permit
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.PreBind != nil {
		in, out := &in.PreBind, &out.PreBind
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.Bind != nil {
		in, out := &in.Bind, &out.Bind
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.PostBind != nil {
		in, out := &in.PostBind, &out.PostBind
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	if in.Unreserve != nil {
		in, out := &in.Unreserve, &out.Unreserve
		*out = new(PluginSet)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugins.
func (in *Plugins) DeepCopy() *Plugins {
	if in == nil {
		return nil
	}
	out := new(Plugins)
	in.DeepCopyInto(out)
	return out
}