/*
Copyright 2021.

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
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ModelSpec defines the desired state of Model
type ModelSpec struct {
	InferenceArtifactSpec `json:",inline"`
	// List of extra requirements for this model to be loaded on a compatible server
	Requirements []string `json:"requirements,omitempty"`
	// Memory needed for model
	// +optional
	Memory *resource.Quantity `json:"memory,omitempty"`
	// Name of the Server to deploy this artifact
	// +optional
	Server *string `json:"server,omitempty"`
	// Number of replicas - default 1
	Replicas *int32 `json:"replicas,omitempty"`
	// Min number of replicas - default equal to replicas
	MinReplicas *int32 `json:"minReplicas,omitempty"`
	// Max number of replicas - default equal to replicas
	MaxReplicas *int32 `json:"maxReplicas,omitempty"`
	// Model already loaded on a server. Don't schedule.
	// Default false
	PreLoaded bool `json:"preloaded,omitempty"`
	// Dedicated server exclusive to this model
	// Default false
	Dedicated bool `json:"dedicated,omitempty"`
	// Payload logging
	Logger *LoggingSpec `json:"logger,omitempty"`
}

type LoggingSpec struct {
	// URI to logging endpoint.
	// +optional
	Uri *string `json:"uri,omitempty"`
	// Percentage of payloads to log
	// Defaults to 100%
	// +optional
	Percent *int32 `json:"percent,omitempty"`
}

type InferenceArtifactSpec struct {
	//Artifact type
	// +optional
	ModelType *string `json:"modelType,omitempty"`
	// Storage URI for the model repository
	// +optional
	StorageURI *string `json:"storageUri,omitempty"`
	// Schema URI
	// +optional
	SchemaURI *string `json:"schemaUri,omitempty"`
	// Service Account Name for auth
	// +optional
	ServiceAccountName *string `json:"serviceAccountName,omitempty"`
	// Secret name
	// +optional
	SecretName *string `json:"secretName,omitempty"`
}

// ModelStatus defines the observed state of Model
type ModelStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Model is the Schema for the models API
type Model struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ModelSpec   `json:"spec,omitempty"`
	Status ModelStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ModelList contains a list of Model
type ModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Model `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Model{}, &ModelList{})
}