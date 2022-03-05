/*
Copyright 2022.

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

// ImagePullSecretInjectorSpec defines the desired state of ImagePullSecretInjector
type ImagePullSecretInjectorSpec struct {
	// Secrets list to inject into service accounts
	Secrets []SecretConfig `json:"secrets"`
	// NamespaceAnnotationSelector filter the namespace where to update service account
	// +optional
	Namespaces []string `json:"namespaces,omitempty"`
	// AnnotationSelector filter the annotation where to update service account
	// at least on of selector should be present
	// +optional
	AnnotationSelector []ServiceAccountSelector `json:"annotationSelector,omitempty"`
	// LabelSelector filter the annotation where to update service account
	// at least on of selector should be present
	// +optional
	LabelSelector []ServiceAccountSelector `json:"labelSelector,omitempty"`
	// NamespaceAnnotationSelector filter the namespace where the service account is to update
	// +optional
	NamespaceAnnotationSelector []ServiceAccountSelector `json:"namespaceAnnotationSelector,omitempty"`
	// NamespaceLabelSelector filter the namespace where the service account is to update
	// +optional
	NamespaceLabelSelector []ServiceAccountSelector `json:"namespaceLabelSelector,omitempty"`
}

// ImagePullSecretInjectorStatus defines the observed state of ImagePullSecretInjector
type ImagePullSecretInjectorStatus struct {
	// Count number of service accounts handled
	Count *int32 `json:"count"`

	// Conditions represents the latest available observations of play
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:resource:scope=Cluster,shortName=impsi
//+kubebuilder:printcolumn:name="Count",type="number",priority=1,JSONPath=".status.count"
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description="CreationTimestamp is a timestamp representing the server time when this object was created. It is not guaranteed to be set in happens-before order across separate operations. Clients may not set this value. It is represented in RFC3339 form and is in UTC."

// ImagePullSecretInjector is the Schema for the imagepullsecretinjectors API
type ImagePullSecretInjector struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImagePullSecretInjectorSpec   `json:"spec,omitempty"`
	Status ImagePullSecretInjectorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ImagePullSecretInjectorList contains a list of ImagePullSecretInjector
type ImagePullSecretInjectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImagePullSecretInjector `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ImagePullSecretInjector{}, &ImagePullSecretInjectorList{})
}
