package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type AILogSummaryConfigSpec struct {
	PodName string `json:"podName"`
}

type AILogSummaryConfigStatus struct {
	LastUpdated metav1.Time `json:"lastUpdated"`
}

type AILogSummaryConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AILogSummaryConfigSpec   `json:"spec"`
	Status AILogSummaryConfigStatus `json:"status,omitempty"`
}

type AILogSummaryConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []AILogSummaryConfig `json:"items"`
}

// DeepCopy methods
func (in *AILogSummaryConfigSpec) DeepCopyInto(out *AILogSummaryConfigSpec) { *out = *in }
func (in *AILogSummaryConfigSpec) DeepCopy() *AILogSummaryConfigSpec {
	if in == nil { return nil }
	out := new(AILogSummaryConfigSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *AILogSummaryConfigStatus) DeepCopyInto(out *AILogSummaryConfigStatus) { *out = *in }
func (in *AILogSummaryConfigStatus) DeepCopy() *AILogSummaryConfigStatus {
	if in == nil { return nil }
	out := new(AILogSummaryConfigStatus)
	in.DeepCopyInto(out)
	return out
}

func (in *AILogSummaryConfig) DeepCopyInto(out *AILogSummaryConfig) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

func (in *AILogSummaryConfig) DeepCopy() *AILogSummaryConfig {
	if in == nil { return nil }
	out := new(AILogSummaryConfig)
	in.DeepCopyInto(out)
	return out
}

func (in *AILogSummaryConfig) DeepCopyObject() runtime.Object {
	if in == nil { return nil }
	out := new(AILogSummaryConfig)
	in.DeepCopyInto(out)
	return out
}

func (in *AILogSummaryConfigList) DeepCopyInto(out *AILogSummaryConfigList) {
	*out = *in
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		out.Items = make([]AILogSummaryConfig, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}
}

func (in *AILogSummaryConfigList) DeepCopy() *AILogSummaryConfigList {
	if in == nil { return nil }
	out := new(AILogSummaryConfigList)
	in.DeepCopyInto(out)
	return out
}

func (in *AILogSummaryConfigList) DeepCopyObject() runtime.Object {
	if in == nil { return nil }
	out := new(AILogSummaryConfigList)
	in.DeepCopyInto(out)
	return out
}

