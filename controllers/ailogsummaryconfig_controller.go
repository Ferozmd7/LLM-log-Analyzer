package controllers

import (
	"context"
	"log"

	aiv1 "github.com/example/llm-log-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AILogSummaryConfigReconciler reconciles AILogSummaryConfig CRs
type AILogSummaryConfigReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	LLM    LLMClient
}

// Reconcile processes AILogSummaryConfig CRs
func (r *AILogSummaryConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	cfg := &aiv1.AILogSummaryConfig{}
	if err := r.Get(ctx, req.NamespacedName, cfg); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// For demo purposes, using PodName as input
	logs := "Dummy logs for pod: " + cfg.Spec.PodName

	// Generate insight
	insightText, err := r.LLM.Generate(logs)
	if err != nil {
		log.Println("LLM generation error:", err)
		return ctrl.Result{}, err
	}

	insight := &aiv1.AIInsight{
		Spec: aiv1.AIInsightSpec{
			PodName: cfg.Spec.PodName,
			Summary: insightText,
		},
	}

	if err := r.Create(ctx, insight); err != nil {
		log.Println("Failed to create AIInsight:", err)
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager registers reconciler with manager
func (r *AILogSummaryConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&aiv1.AILogSummaryConfig{}).
		Complete(r)
}

