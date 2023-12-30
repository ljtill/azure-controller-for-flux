/*
MIT License

Copyright (c) 2023 Lyon Till

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeploymentStackSourceRef defines the desired source reference of DeploymentStack
type DeploymentStackSourceRef struct {
	// Kind of the Deployment Stack Source Reference
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=GitRepository
	Kind string `json:"kind,omitempty"`

	// Name of the Deployment Stack Source Reference
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=64
	Name string `json:"name,omitempty"`

	// Namespace of the Deployment Stack Source Reference
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=64
	Namespace string `json:"namespace,omitempty"`
}

// DeploymentStackSpec defines the desired state of DeploymentStack
type DeploymentStackSpec struct {
	// Name of the Deployment Stack
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=64
	Name string `json:"stackName"`

	// Scope of the Deployment Stack
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=ManagementGroup;Subscription;ResourceGroup
	Scope string `json:"deploymentScope"`

	// Location of the Deployment Stack
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=64
	Location string `json:"location"`

	// Template of the Deployment Stack// +kubebuilder:validation:Required
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=64
	Template string `json:"templatePath"`

	// Template Spec of the Deployment Stack
	// TODO: Not implemented yet
	TemplateSpec string `json:"templateSpec,omitempty"`

	// Management Group ID of the Deployment Stack
	ManagementGroupId string `json:"managementGroupId,omitempty"`

	// Subscription ID of the Deployment Stack
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=32
	SubscriptionId string `json:"subscriptionId,omitempty"`

	// Resource Group ID of the Deployment Stack
	ResourceGroupName string `json:"resourceGroupId,omitempty"`

	// Deny Settings Mode of the Deployment Stack
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=None;DenyDelete;DenyWriteAndDelete
	DenySettings string `json:"denySettingsMode"`

	// Delete Resources of the Deployment Stack
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=true;false
	DeleteResources bool `json:"deleteResources,omitempty"`

	// Delete Resource Groups of the Deployment Stack
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=true;false
	DeleteResourceGroups bool `json:"deleteResourceGroups,omitempty"`

	// Source Reference of the Deployment Stack
	SourceRef DeploymentStackSourceRef `json:"sourceRef"`

	// Interval of the Deployment Stack
	Interval string `json:"interval,omitempty"`

	// Retry Interval of the Deployment Stack
	RetryInterval string `json:"retryInterval,omitempty"`
}

// DeploymentStackStatus defines the observed state of DeploymentStack
type DeploymentStackStatus struct{}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DeploymentStack is the Schema for the deploymentstacks API
type DeploymentStack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeploymentStackSpec   `json:"spec,omitempty"`
	Status DeploymentStackStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DeploymentStackList contains a list of DeploymentStack
type DeploymentStackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeploymentStack `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DeploymentStack{}, &DeploymentStackList{})
}
