package controllers

import (
	"context"

	aiv1 "github.com/example/llm-log-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Reconciler
type AILogSummaryConfigReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	LLM    LLMClient
}

func (r *AILogSummaryConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	cfg := &aiv1.AILogSummaryConfig{}
	if err := r.Get(ctx, req.NamespacedName, cfg); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	logs, err := r.streamPodLogsReq(ctx, cfg)
	if err != nil { return ctrl.Result{}, err }

	r.emitEvent(cfg, "Logs processed")

	if err := r.createInsight(ctx, logs); err != nil { return ctrl.Result{}, err }

	return ctrl.Result{}, nil
}

// ------------------------
// Helper functions
// ------------------------
func (r *AILogSummaryConfigReconciler) streamPodLogsReq(ctx context.Context, cfg *aiv1.AILogSummaryConfig) (string, error) {
	// placeholder: fetch logs from Pod
	return "dummy pod logs", nil
}

func (r *AILogSummaryConfigReconciler) emitEvent(cfg *aiv1.AILogSummaryConfig, msg string) {
	// placeholder: create Kubernetes event
}

func (r *AILogSummaryConfigReconciler) createInsight(ctx context.Context, summary string) error {
	response, err := r.LLM.Generate(summary)
	if err != nil { return err }

	// Create AIInsight object
	insight := &aiv1.AIInsight{
		Spec: aiv1.AIInsightSpec{
			Summary: response,
		},
	}
	return r.Client.Create(ctx, insight)
}

func (r *AILogSummaryConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&aiv1.AILogSummaryConfig{}).
		Complete(r)
}

