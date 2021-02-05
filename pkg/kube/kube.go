package kube

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientapi "k8s.io/client-go/tools/clientcmd/api"
)

func NewClient(kubeconfig string) (*kubernetes.Clientset, error) {
	c, err := buildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(c)
}

func buildConfigFromFlags(masterURL, kubeconfigPath string) (*rest.Config, error){
	if kubeconfigPath == "" && masterURL == "" {
		kubeconfig, err := rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
		return kubeconfig, nil
	}

	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath:kubeconfigPath},
		&clientcmd.ConfigOverrides{ClusterInfo:clientapi.Cluster{Server: masterURL}}).ClientConfig()
}