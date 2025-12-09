package v1alpha1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type AILogSummaryConfigSpec struct {
    PodSelector        map[string]string `json:"podSelector"`
    IntervalSeconds    int               `json:"intervalSeconds"`
    MaxLogBytes        int               `json:"maxLogBytes"`
    LLMModel           string            `json:"llmModel"`
    GenerateInsightsCR bool              `json:"generateInsightsCR"`
}

type AILogSummaryConfigStatus struct {
    LastAnalysisTime metav1.Time       `json:"lastAnalysisTime"`
    Conditions       []metav1.Condition `json:"conditions,omitempty"`
}

type AILogSummaryConfig struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec   AILogSummaryConfigSpec   `json:"spec,omitempty"`
    Status AILogSummaryConfigStatus `json:"status,omitempty"`
}
