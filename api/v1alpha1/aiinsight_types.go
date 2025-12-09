package v1alpha1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT: define your spec and status as before

type AIInsightSpec struct {
    PodName         string   `json:"podName"`
    Summary         string   `json:"summary"`
    Anomalies       []string `json:"anomalies"`
    Recommendations []string `json:"recommendations"`
}

type AIInsightStatus struct {
    CreatedAt metav1.Time `json:"createdAt"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type AIInsight struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec   AIInsightSpec   `json:"spec"`
    Status AIInsightStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type AIInsightList struct {
    metav1.TypeMeta `json:",inline"`
    metav1.ListMeta `json:"metadata,omitempty"`
    Items           []AIInsight `json:"items"`
}

