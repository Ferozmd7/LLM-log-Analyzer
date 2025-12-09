package main

import (
    "os"
    "github.com/example/llm-log-operator/controllers"
    "github.com/example/llm-log-operator/api/v1alpha1"
    ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
    mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
        Scheme: controllers.Scheme,
    })
    if err != nil {
        panic(err)
    }

    llm := controllers.NewLLMClient()

    if err := (&controllers.AILogSummaryConfigReconciler{
        Client: mgr.GetClient(),
        Scheme: mgr.GetScheme(),
        LLM:    llm,
    }).SetupWithManager(mgr); err != nil {
        panic(err)
    }

    if err := (&controllers.AIInsightReconciler{
        Client: mgr.GetClient(),
        Scheme: mgr.GetScheme(),
    }).SetupWithManager(mgr); err != nil {
        panic(err)
    }

    if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
        panic(err)
    }
}
