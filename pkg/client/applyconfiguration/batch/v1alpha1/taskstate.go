/*
Copyright The Volcano Authors.

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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// TaskStateApplyConfiguration represents an declarative configuration of the TaskState type for use
// with apply.
type TaskStateApplyConfiguration struct {
	Phase map[v1.PodPhase]int32 `json:"phase,omitempty"`
}

// TaskStateApplyConfiguration constructs an declarative configuration of the TaskState type for use with
// apply.
func TaskState() *TaskStateApplyConfiguration {
	return &TaskStateApplyConfiguration{}
}

// WithPhase puts the entries into the Phase field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Phase field,
// overwriting an existing map entries in Phase field with the same key.
func (b *TaskStateApplyConfiguration) WithPhase(entries map[v1.PodPhase]int32) *TaskStateApplyConfiguration {
	if b.Phase == nil && len(entries) > 0 {
		b.Phase = make(map[v1.PodPhase]int32, len(entries))
	}
	for k, v := range entries {
		b.Phase[k] = v
	}
	return b
}
