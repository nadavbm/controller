package controller

import (
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// connectInCluster creates kubernetes client for pods running inside kubernetes cluster
func connectInCluster() (*kubernetes.Clientset, error) {
	restConfig, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}

	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	return clientSet, nil
}

// connectFromKubeConfig creates kubernetes client from local kube config (in homedir)
func connectFromKubeConfig() (*kubernetes.Clientset, error) {
	kubeConfig := fmt.Sprintf("%s/.kube/config", os.Getenv("HOME"))
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return nil, err
	}

	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		return nil, err
	}

	return clientSet, nil
}
