package main

import (
	"fmt"
	"os"

	"github.com/nadavbm/controller/pkg/logger"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/runtime"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	logger := logger.New()

	client, err := getClient(logger)
	if err != nil {
		panic(err)
	}
	factory := informers.NewSharedInformerFactory(client.client, 0)
	podsInformer := factory.Core().V1().Pods().Informer()
	stopper := make(chan struct{})
	defer close(stopper)
	defer runtime.HandleCrash()
	podsInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    client.onAdd,
		UpdateFunc: client.onUpdate,
		DeleteFunc: client.onDelete,
	})
	go podsInformer.Run(stopper)
	if !cache.WaitForCacheSync(stopper, podsInformer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Timed out waiting for caches to sync"))
		return
	}
	<-stopper
}

// onAdd
func (c *conn) onAdd(obj interface{}) {
	pod := obj.(*corev1.Pod)
	c.logger.Info(fmt.Sprintf("pod %s added in namespace %s", pod.GetName(), pod.GetNamespace()))
}

func (c *conn) onUpdate(oldObj, obj interface{}) {
	oldPod := oldObj.(*corev1.Pod)
	pod := obj.(*corev1.Pod)
	c.logger.Info(fmt.Sprintf("pod %s updated, old pod %s in namespace %s", pod.GetName(), oldPod.GetName(), pod.GetNamespace()))
}

// onDelete
func (c *conn) onDelete(obj interface{}) {
	pod := obj.(*corev1.Pod)
	c.logger.Info(fmt.Sprintf("pod %s deleted in namespace %s", pod.GetName(), pod.GetNamespace()))
}

// Client is a kubernetes api client
type Client interface {
	Connect() (*conn, error)
}

type conn struct {
	logger *logger.Log
	client *kubernetes.Clientset
}

// Connect creates a kubernetes api client that can communicate with master api
func getClient(logger *logger.Log) (*conn, error) {
	_, inCluster := os.LookupEnv("KUBERNETES_PORT")
	if inCluster {
		logger.Info("connect to kubernetes api in cluster")
		k8sClient, err := connectInCluster()
		if err != nil {
			return nil, err
		}

		return &conn{
			logger: logger,
			client: k8sClient,
		}, nil
	}

	logger.Info("connect to kubernetes api from kube config")
	k8sClient, err := connectFromKubeConfig()
	if err != nil {
		return nil, err
	}

	return &conn{
		logger: logger,
		client: k8sClient,
	}, nil
}

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
