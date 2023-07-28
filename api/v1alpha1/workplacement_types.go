/*
Copyright 2021 Syntasso.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// WorkPlacementSpec defines the desired state of WorkPlacement
type WorkPlacementSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	TargetClusterName string `json:"targetClusterName,omitempty"`

	// Workload represents the manifest workload to be deployed on worker
	Workload WorkloadTemplate `json:"workload,omitempty"`
}

// WorkPlacementStatus defines the observed state of WorkPlacement
type WorkPlacementStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// WorkPlacement is the Schema for the workplacements API
type WorkPlacement struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkPlacementSpec   `json:"spec,omitempty"`
	Status WorkPlacementStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// WorkPlacementList contains a list of WorkPlacement
type WorkPlacementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkPlacement `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WorkPlacement{}, &WorkPlacementList{})
}
