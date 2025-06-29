# controller

controller is a simple kubernetes controller that listening to kubernetes events and log it to the console.

this app is programmed to listen only to pods and secrets events in kubernetes at the moment, but will be extended later.

### install

use helm to install:

```
helm install controller helm/ -n controller --create-namespaces
```

### watch

check the pod logs with `kubectl`:

```
$kubectl logs -controller-9b7bf7cfc-8s7tz z 
INFO [2025-06-29 08:45:33] MESSAGE: connect to kubernetes api in cluster
INFO [2025-06-29 08:45:33] MESSAGE: strat running controller
INFO [2025-06-29 08:45:33] MESSAGE: pod controller-9b7bf7cfc-8s7tz added in namespace controller
INFO [2025-06-29 08:45:33] MESSAGE: pod coredns-674b8bbfcf-q8lrg added in namespace kube-system
INFO [2025-06-29 08:45:33] MESSAGE: pod etcd-minikube added in namespace kube-system
INFO [2025-06-29 08:45:33] MESSAGE: pod kube-apiserver-minikube added in namespace kube-system
INFO [2025-06-29 08:45:33] MESSAGE: pod kube-controller-manager-minikube added in namespace kube-system
INFO [2025-06-29 08:45:33] MESSAGE: pod kube-proxy-4wdqp added in namespace kube-system
INFO [2025-06-29 08:45:33] MESSAGE: pod kube-scheduler-minikube added in namespace kube-system
INFO [2025-06-29 08:45:33] MESSAGE: pod storage-provisioner added in namespace kube-system
INFO [2025-06-29 08:45:33] MESSAGE: pod secret-distributor-controller-manager-5c4448d5d4-x5pw5 added in namespace secret-distributor-system
INFO [2025-06-29 08:45:33] MESSAGE: pod sh.helm.release.v1.controller.v1 added in namespace controller
INFO [2025-06-29 08:45:33] MESSAGE: pod docker-config added in namespace etzba
INFO [2025-06-29 08:45:33] MESSAGE: pod docker-config added in namespace secret-distributor-system
INFO [2025-06-29 08:45:34] MESSAGE: pod controller-9b7bf7cfc-8s7tz updated, old pod controller-9b7bf7cfc-8s7tz in namespace controller
INFO [2025-06-29 08:46:02] MESSAGE: pod secret-distributor-controller-manager-5c4448d5d4-x5pw5 updated, old pod secret-distributor-controller-manager-5c4448d5d4-x5pw5 in namespace secret-distributor-system
INFO [2025-06-29 08:46:02] MESSAGE: pod secret-distributor-controller-manager-5c4448d5d4-62wns added in namespace secret-distributor-system
INFO [2025-06-29 08:46:02] MESSAGE: pod secret-distributor-controller-manager-5c4448d5d4-62wns updated, old pod secret-distributor-controller-manager-5c4448d5d4-62wns in namespace secret-distributor-system
INFO [2025-06-29 08:46:02] MESSAGE: pod secret-distributor-controller-manager-5c4448d5d4-62wns updated, old pod secret-distributor-controller-manager-5c4448d5d4-62wns in namespace secret-distributor-system
INFO [2025-06-29 08:46:02] MESSAGE: pod secret-distributor-controller-manager-5c4448d5d4-x5pw5 updated, old pod secret-distributor-controller-manager-5c4448d5d4-x5pw5 in namespace secret-distributor-system
INFO [2025-06-29 08:46:04] MESSAGE: pod secret-distributor-controller-manager-5c4448d5d4-x5pw5 updated, old pod secret-distributor-controller-manager-5c4448d5d4-x5pw5 in namespace secret-distributor-system
INFO [2025-06-29 08:46:04] MESSAGE: pod secret-distributor-controller-manager-5c4448d5d4-x5pw5 deleted in namespace secret-distributor-system
INFO [2025-06-29 08:46:05] MESSAGE: pod secret-distributor-controller-manager-5c4448d5d4-62wns updated, old pod secret-distributor-controller-manager-5c4448d5d4-62wns in namespace secret-distributor-system
INFO [2025-06-29 08:46:17] MESSAGE: pod secret-distributor-controller-manager-5c4448d5d4-62wns updated, old pod secret-distributor-controller-manager-5c4448d5d4-62wns in namespace secret-distributor-system
```
