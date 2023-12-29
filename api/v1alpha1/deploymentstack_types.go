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

// DeploymentStackScope defines the desired scope of DeploymentStack
type DeploymentStackScope string

const (
	Tenant          DeploymentStackScope = "Tenant"
	ManagementGroup DeploymentStackScope = "ManagementGroup"
	Subscription    DeploymentStackScope = "Subscription"
	ResourceGroup   DeploymentStackScope = "ResourceGroup"
)

// DeploymentStackDenySettingMode defines the desired deny setting mode of DeploymentStack
type DeploymentStackDenySettings string

const (
	None               DeploymentStackDenySettings = "None"
	DenyDelete         DeploymentStackDenySettings = "DenyDelete"
	DenyWriteAndDelete DeploymentStackDenySettings = "DenyWriteAndDelete"
)

// DeploymentStackSourceRefKind defines the desired source reference kind of DeploymentStack
type DeploymentStackSourceRefKind string

const (
	GitRepository DeploymentStackSourceRefKind = "GitRepository"
)

// DeploymentStackSourceRef defines the desired source reference of DeploymentStack
type DeploymentStackSourceRef struct {
	Kind      DeploymentStackSourceRefKind `json:"kind,omitempty"`
	Name      string                       `json:"name,omitempty"`
	Namespace string                       `json:"namespace,omitempty"`
}

// DeploymentStackSpec defines the desired state of DeploymentStack
type DeploymentStackSpec struct {
	Name          string                      `json:"stackName,omitempty"`
	Scope         DeploymentStackScope        `json:"deploymentScope,omitempty"`
	Location      string                      `json:"location,omitempty"`
	Template      string                      `json:"templatePath,omitempty"`
	DenySettings  DeploymentStackDenySettings `json:"denySettingsMode,omitempty"`
	SourceRef     DeploymentStackSourceRef    `json:"sourceRef,omitempty"`
	Interval      string                      `json:"interval,omitempty"`
	RetryInterval string                      `json:"retryInterval,omitempty"`
}

// DeploymentStackStatus defines the observed state of DeploymentStack
type DeploymentStackStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

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
