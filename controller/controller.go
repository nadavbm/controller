package controller

import (
	"fmt"
	"os"

	"github.com/nadavbm/controller/pkg/logger"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

// Controller
type Controller struct {
	logger               *logger.Log
	Client               *kubernetes.Clientset
	SharedFactory        informers.SharedInformerFactory
	PodsInformer         *PodsInformer
	PodsInformerIndex    cache.SharedIndexInformer
	SecretsInformer      *SecretsInformer
	SecretsInformerIndex cache.SharedIndexInformer
}

// New creates a new instance of controller
func New(logger *logger.Log) (*Controller, error) {
	_, inCluster := os.LookupEnv("KUBERNETES_PORT")
	controller := &Controller{
		logger: logger,
	}
	if inCluster {
		logger.Info("connect to kubernetes api in cluster")
		k8sClient, err := connectInCluster()
		if err != nil {
			return nil, err
		}
		controller.Client = k8sClient
	} else {
		logger.Info("connect to kubernetes api from kube config")
		k8sClient, err := connectFromKubeConfig()
		if err != nil {
			return nil, err
		}
		controller.Client = k8sClient
	}
	controller.SharedFactory = informers.NewSharedInformerFactory(controller.Client, 0)
	controller.PodsInformerIndex = controller.newPodsInformer()
	controller.SecretsInformerIndex = controller.newSecretsInformer()
	return controller, nil
}

func (c *Controller) Run(stopper chan struct{}) {
	c.logger.Info("strat running controller")
	go c.PodsInformerIndex.Run(stopper)
	if !cache.WaitForCacheSync(stopper, c.PodsInformerIndex.HasSynced) {
		runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
		return
	}

	go c.SecretsInformerIndex.Run(stopper)
	if !cache.WaitForCacheSync(stopper, c.SecretsInformerIndex.HasSynced) {
		runtime.HandleError(fmt.Errorf("timed out waiting for caches to sync"))
		return
	}
}
