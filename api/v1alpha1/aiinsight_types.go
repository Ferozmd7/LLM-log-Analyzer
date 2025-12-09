package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type AIInsightSpec struct {
	PodName         string   `json:"podName"`
	Summary         string   `json:"summary"`
	Anomalies       []string `json:"anomalies"`
	Recommendations []string `json:"recommendations"`
}

type AIInsightStatus struct {
	CreatedAt metav1.Time `json:"createdAt"`
}

type AIInsight struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AIInsightSpec   `json:"spec"`
	Status AIInsightStatus `json:"status,omitempty"`
}

type AIInsightList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []AIInsight `json:"items"`
}

// DeepCopy methods
func (in *AIInsightSpec) DeepCopyInto(out *AIInsightSpec) {
	*out = *in
	if in.Anomalies != nil {
		out.Anomalies = make([]string, len(in.Anomalies))
		copy(out.Anomalies, in.Anomalies)
	}
	if in.Recommendations != nil {
		out.Recommendations = make([]string, len(in.Recommendations))
		copy(out.Recommendations, in.Recommendations)
	}
}

func (in *AIInsightSpec) DeepCopy() *AIInsightSpec {
	if in == nil { return nil }
	out := new(AIInsightSpec)
	in.DeepCopyInto(out)
	return out
}

func (in *AIInsightStatus) DeepCopyInto(out *AIInsightStatus) { *out = *in }
func (in *AIInsightStatus) DeepCopy() *AIInsightStatus {
	if in == nil { return nil }
	out := new(AIInsightStatus)
	in.DeepCopyInto(out)
	return out
}

func (in *AIInsight) DeepCopyInto(out *AIInsight) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

func (in *AIInsight) DeepCopy() *AIInsight {
	if in == nil { return nil }
	out := new(AIInsight)
	in.DeepCopyInto(out)
	return out
}

func (in *AIInsight) DeepCopyObject() runtime.Object {
	if in == nil { return nil }
	out := new(AIInsight)
	in.DeepCopyInto(out)
	return out
}

func (in *AIInsightList) DeepCopyInto(out *AIInsightList) {
	*out = *in
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		out.Items = make([]AIInsight, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}
}

func (in *AIInsightList) DeepCopy() *AIInsightList {
	if in == nil { return nil }
	out := new(AIInsightList)
	in.DeepCopyInto(out)
	return out
}

func (in *AIInsightList) DeepCopyObject() runtime.Object {
	if in == nil { return nil }
	out := new(AIInsightList)
	in.DeepCopyInto(out)
	return out
}

