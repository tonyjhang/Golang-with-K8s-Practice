## Installation
This reop is a golang practice with k8s with local env.

## Install Kubectl
Reference below
[ref](https://kubernetes.io/docs/tasks/tools/#kubectl)

Here is example for MacOS.
```
brew install kubectl 
```
## Install Kind
Kind is useful tools can setup and remove k8s cluster easily.

Reference below
[ref](https://kind.sigs.k8s.io/docs/user/quick-start/#installation)
Here is example for MacOS.
```
brew install kind
kind create cluster
```
List cluster
```
kind get clusters
```
Delete cluster
```
kind delete cluster
```
## Run project in k8s
Build image then upload it, after that remember chagne image field in kube-manifests/GolangRestfulPractice-Deployment.yaml
```
cd kube-manifests
kubectl apply -f
```