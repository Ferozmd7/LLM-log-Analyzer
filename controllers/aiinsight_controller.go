package controllers

import (
    "context"

    ctrl "sigs.k8s.io/controller-runtime"
    "sigs.k8s.io/controller-runtime/pkg/client"

    aiv1 "github.com/example/llm-log-operator/api/v1alpha1"
)

type AIInsightReconciler struct {
    client.Client
    Scheme *runtime.Scheme
}

func (r *AIInsightReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    var insight aiv1.AIInsight
    if err := r.Get(ctx, req.NamespacedName, &insight); err != nil {
        return ctrl.Result{}, client.IgnoreNotFound(err)
    }

    return ctrl.Result{}, nil
}

func (r *AIInsightReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).
        For(&aiv1.AIInsight{}).
        Complete(r)
}
