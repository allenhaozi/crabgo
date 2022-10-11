package common

type KubeNativeAbility string

const (
	KubeDeployment          KubeNativeAbility = "Deployment"
	KubeStatefulSet         KubeNativeAbility = "StatefulSet"
	KubeDaemonSet           KubeNativeAbility = "DaemonSet"
	KubeJob                 KubeNativeAbility = "Job"
	KubeCronJob             KubeNativeAbility = "CronJob"
	KubeService             KubeNativeAbility = "Service"
	KubeServiceMonitor      KubeNativeAbility = "ServiceMonitor"
	KubeIngress             KubeNativeAbility = "Ingress"
	KubeConfigMap           KubeNativeAbility = "ConfigMap"
	KubePod                 KubeNativeAbility = "Pod"
	KubeServiceAPIVersion                     = "v1"
	KubeEndpoints           KubeNativeAbility = "Endpoints"
	KubeEndpointsAPIVersion                   = "v1"
	KubeNode                                  = "Node"
)

func (kna KubeNativeAbility) String() string {
	return string(kna)
}
