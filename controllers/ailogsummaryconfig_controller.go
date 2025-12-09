package controllers

import (
    "bufio"
    "bytes"
    "context"
    "time"

    corev1 "k8s.io/api/core/v1"
    "k8s.io/apimachinery/pkg/labels"
    "k8s.io/apimachinery/pkg/runtime"
    ctrl "sigs.k8s.io/controller-runtime"
    "sigs.k8s.io/controller-runtime/pkg/client"
    "sigs.k8s.io/controller-runtime/pkg/log"

    aiv1 "github.com/example/llm-log-operator/api/v1alpha1"
)

type AILogSummaryConfigReconciler struct {
    client.Client
    Scheme *runtime.Scheme
    LLM    *LLMClient
}

func (r *AILogSummaryConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    logger := log.FromContext(ctx)

    var cfg aiv1.AILogSummaryConfig
    if err := r.Get(ctx, req.NamespacedName, &cfg); err != nil {
        return ctrl.Result{}, client.IgnoreNotFound(err)
    }

    selector := labels.SelectorFromSet(cfg.Spec.PodSelector)
    var pods corev1.PodList

    if err := r.List(ctx, &pods, &client.ListOptions{
        Namespace:     req.Namespace,
        LabelSelector: selector,
    }); err != nil {
        return ctrl.Result{}, err
    }

    buf := new(bytes.Buffer)

    for _, pod := range pods.Items {
        req := r.streamPodLogsReq(&pod)
        stream, err := req.Stream(ctx)
        if err != nil {
            logger.Error(err, "Failed to stream logs", "pod", pod.Name)
            continue
        }
        defer stream.Close()

        s := bufio.NewScanner(stream)
        s.Buffer(make([]byte, 4096), cfg.Spec.MaxLogBytes)

        for s.Scan() {
            line := s.Bytes()
            if buf.Len()+len(line) > cfg.Spec.MaxLogBytes {
                buf.Next(len(line))
            }
            buf.Write(line)
            buf.WriteByte('\n')
        }
    }

    logs := buf.String()
    if len(logs) == 0 {
        return ctrl.Result{RequeueAfter: d(cfg.Spec.IntervalSeconds)}, nil
    }

    analysis, err := r.LLM.AnalyzeLogs(ctx, cfg.Spec.LLMModel, logs)
    if err != nil {
        logger.Error(err, "LLM call failed")
        return ctrl.Result{RequeueAfter: d(cfg.Spec.IntervalSeconds)}, nil
    }

    r.emitEvent(ctx, &cfg, analysis)

    if cfg.Spec.GenerateInsightsCR {
        _ = r.createInsight(ctx, &cfg, analysis)
    }

    return ctrl.Result{RequeueAfter: d(cfg.Spec.IntervalSeconds)}, nil
}

func (r *AILogSummaryConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).
        For(&aiv1.AILogSummaryConfig{}).
        Complete(r)
}

func d(s int) time.Duration { return time.Duration(s) * time.Second }
