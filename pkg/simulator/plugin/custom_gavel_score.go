package plugin

import (
	"context"
	"fmt"

	simontype "github.com/hkust-adsl/kubernetes-scheduler-simulator/pkg/type"
	"github.com/hkust-adsl/kubernetes-scheduler-simulator/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

type CustomScorePlugin struct {
	handle framework.Handle
}

// Name implements framework.Plugin.
func (*CustomScorePlugin) Name() string {
	return simontype.CustomScorePluginName
}

// var _ framework.ScorePlugin = &CustomScorePlugin{}

func NewCustomScorePlugin(configuration runtime.Object, handle framework.Handle) (framework.Plugin, error) {
	return &CustomScorePlugin{
		handle: handle,
	}, nil
}

func (plugin *CustomScorePlugin) Score(ctx context.Context, state *framework.CycleState, pod *corev1.Pod, nodeName string) (int64, *framework.Status) {
	nodeResPtr := utils.GetNodeResourceViaHandleAndName(plugin.handle, nodeName)
	if nodeResPtr == nil {
		return framework.MinNodeScore, framework.NewStatus(framework.Error,
			fmt.Sprintf("failed to get nodeRes(%s)\n", nodeName))
	}

	nodeRes := *nodeResPtr
	podRes := utils.GetPodResource(pod)

	score, _ := calculateCustomGavelScore(nodeRes, podRes)
	return score, nil
}

func calculateCustomGavelScore(nodeRes simontype.NodeResource, podRes simontype.PodResource) (int64, string) {
	return 100, ""
}
