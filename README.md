# LLM Log Analyzer Operator

A production-grade Kubernetes operator that:
- Streams pod logs
- Sends batches to an LLM (OpenAI, Anthropic, etc.)
- Detects anomalies
- Creates Kubernetes Events
- Optionally stores persistent insights in an AIInsight CRD

## Features
- Go + Kubebuilder
- CRDs with validation
- AIInsight storage
- LLM provider abstraction
- Helm chart
- Kind demo environment

## Quickstart
```bash
make docker-build
make docker-push
helm install llm-log-operator ./helm/llm-log-operator
kubectl apply -f config/samples/ailogsummaryconfig.yaml
