/*
Copyright 2024 Anton Shadrin.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RestoreRequestSpec defines the desired state of RestoreRequest.
type RestoreRequestSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// +kubebuilder:validation:MinLength=0

	// ResourceSelector is a key=value label throug which operator determines which resource to restore.
	ResourceSelector string `json:"resourceSelector"`

	// Version is a version to restore. Could be retrived from BackupRequest spec.
	// +optional
	Version *int32 `json:"version,omitempty"`
}

// RestoreRequestStatus defines the observed state of RestoreRequest.
type RestoreRequestStatus struct {

	// RestoreStatus contains information about current restore process.
	RestoreStatus string `json:"restoreStatus"`

	// Active is a pointer to restore job
	// +optional
	Active corev1.ObjectReference `json:"active,omitempty"`
}

// +kubebuilder:validation:Enum=Processing;Failed;Succeded

// RestoreStatus is a valid status string
type RestoreStatus string

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// RestoreRequest is the Schema for the restorerequests API.
type RestoreRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RestoreRequestSpec   `json:"spec,omitempty"`
	Status RestoreRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RestoreRequestList contains a list of RestoreRequest.
type RestoreRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RestoreRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RestoreRequest{}, &RestoreRequestList{})
}
