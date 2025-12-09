package main

import (
	"os"

	"github.com/example/llm-log-operator/controllers"
	aiv1 "github.com/example/llm-log-operator/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	// Create controller manager
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme: nil, // We'll register CRDs below
	})
	if err != nil {
		panic(err)
	}

	// Register CRDs with the scheme
	if err := aiv1.AddToScheme(mgr.GetScheme()); err != nil {
		panic(err)
	}

	// Initialize LLM client (Ollama in this build)
	var llmClient controllers.LLMClient
	if os.Getenv("USE_OLLAMA") == "1" {
		model := os.Getenv("OLLAMA_MODEL")
		llmClient = controllers.NewOllamaClient(model)
	} else {
		panic("No supported LLM backend configured. Only Ollama is enabled in this build.")
	}

	// Setup the reconciler
	reconciler := &controllers.AILogSummaryConfigReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
		LLM:    llmClient,
	}

	if err := reconciler.SetupWithManager(mgr); err != nil {
		panic(err)
	}

	// Start the manager
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		panic(err)
	}
}

