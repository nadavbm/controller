package controller

import (
	"fmt"

	"github.com/nadavbm/controller/pkg/logger"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
)

// PodsInformer
type PodsInformer struct {
	logger *logger.Log
}

// newPodsInformer add new pods informer event handlers
func (c *Controller) newPodsInformer() cache.SharedIndexInformer {
	c.PodsInformer = &PodsInformer{c.logger}
	podsInformer := c.SharedFactory.Core().V1().Pods().Informer()
	podsInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.PodsInformer.OnAdd,
		UpdateFunc: c.PodsInformer.OnUpdate,
		DeleteFunc: c.PodsInformer.OnDelete,
	})
	return podsInformer
}

// OnAdd watches resource addition events and log it
func (c *PodsInformer) OnAdd(obj interface{}) {
	pod := obj.(*corev1.Pod)
	c.logger.Info(fmt.Sprintf("pod %s added in namespace %s", pod.GetName(), pod.GetNamespace()))
}

// OnUpdate watches resource update events and log it
func (c *PodsInformer) OnUpdate(oldObj, obj interface{}) {
	oldPod := oldObj.(*corev1.Pod)
	pod := obj.(*corev1.Pod)
	c.logger.Info(fmt.Sprintf("pod %s updated, old pod %s in namespace %s", pod.GetName(), oldPod.GetName(), pod.GetNamespace()))
}

// OnDelete watches resource deletion events and log it
func (c *PodsInformer) OnDelete(obj interface{}) {
	pod := obj.(*corev1.Pod)
	c.logger.Info(fmt.Sprintf("pod %s deleted in namespace %s", pod.GetName(), pod.GetNamespace()))
}

// SecretInformer
type SecretsInformer struct {
	logger *logger.Log
}

// newSecretsInformer add new secrets informer event handlers
func (c *Controller) newSecretsInformer() cache.SharedIndexInformer {
	c.SecretsInformer = &SecretsInformer{c.logger}
	secretsInformer := c.SharedFactory.Core().V1().Secrets().Informer()
	secretsInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    c.SecretsInformer.OnAdd,
		UpdateFunc: c.SecretsInformer.OnUpdate,
		DeleteFunc: c.SecretsInformer.OnDelete,
	})
	return secretsInformer
}

// OnAdd watches resource addition events and log it
func (c *SecretsInformer) OnAdd(obj interface{}) {
	pod := obj.(*corev1.Secret)
	c.logger.Info(fmt.Sprintf("pod %s added in namespace %s", pod.GetName(), pod.GetNamespace()))
}

// OnUpdate watches resource update events and log it
func (c *SecretsInformer) OnUpdate(oldObj, obj interface{}) {
	oldPod := oldObj.(*corev1.Secret)
	pod := obj.(*corev1.Secret)
	c.logger.Info(fmt.Sprintf("pod %s updated, old pod %s in namespace %s", pod.GetName(), oldPod.GetName(), pod.GetNamespace()))
}

// OnDelete watches resource deletion events and log it
func (c *SecretsInformer) OnDelete(obj interface{}) {
	pod := obj.(*corev1.Secret)
	c.logger.Info(fmt.Sprintf("pod %s deleted in namespace %s", pod.GetName(), pod.GetNamespace()))
}
