package main

import (
	"log"
	"os"

	"github.com/example/llm-log-operator/controllers"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{})
	if err != nil { log.Fatal(err) }

	var llmClient controllers.LLMClient
	if os.Getenv("USE_LOCAL_LLM") == "1" {
		llmClient = controllers.NewLocalLLMClient()
	} else {
		apiKey := os.Getenv("OPENAI_API_KEY")
		llmClient = controllers.NewOpenAIClient(apiKey)
	}

	reconciler := &controllers.AILogSummaryConfigReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
		LLM:    llmClient,
	}

	if err := reconciler.SetupWithManager(mgr); err != nil { log.Fatal(err) }

	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil { log.Fatal(err) }
}

