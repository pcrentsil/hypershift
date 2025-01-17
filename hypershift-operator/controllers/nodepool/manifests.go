package nodepool

import (
	"fmt"

	hyperv1 "github.com/openshift/hypershift/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	capiv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

const (
	EC2VolumeDefaultSize int64  = 16
	EC2VolumeDefaultType string = "gp3"
)

func machineDeployment(nodePool *hyperv1.NodePool, clusterName string, controlPlaneNamespace string) *capiv1.MachineDeployment {
	resourcesName := generateName(clusterName, nodePool.Spec.ClusterName, nodePool.GetName())
	return &capiv1.MachineDeployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resourcesName,
			Namespace: controlPlaneNamespace,
		},
	}
}

func machineSet(nodePool *hyperv1.NodePool, clusterName string, controlPlaneNamespace string) *capiv1.MachineSet {
	resourcesName := generateName(clusterName, nodePool.Spec.ClusterName, nodePool.GetName())
	return &capiv1.MachineSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      resourcesName,
			Namespace: controlPlaneNamespace,
		},
	}
}

func machineHealthCheck(nodePool *hyperv1.NodePool, controlPlaneNamespace string) *capiv1.MachineHealthCheck {
	return &capiv1.MachineHealthCheck{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      nodePool.GetName(),
			Namespace: controlPlaneNamespace,
		},
	}
}

func inPlaceUpgradePod(namespace, nodeName string) *corev1.Pod {
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      fmt.Sprintf("machine-config-daemon-%s", nodeName),
		},
	}
}

func inPlaceUpgradeNamespace(nodePool *hyperv1.NodePool) *corev1.Namespace {
	return &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%s-upgrade", nodePool.GetName()),
		},
	}
}

func inPlaceUpgradeConfigMap(nodePool *hyperv1.NodePool, namespace string) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      fmt.Sprintf("%s-upgrade", nodePool.GetName()),
		},
	}
}

func IgnitionUserDataSecret(namespace, name, payloadInputHash string) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      fmt.Sprintf("user-data-%s-%s", name, payloadInputHash),
		},
	}
}

func TokenSecret(namespace, name, payloadInputHash string) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: namespace,
			Name:      fmt.Sprintf("token-%s-%s", name, payloadInputHash),
		},
	}
}
