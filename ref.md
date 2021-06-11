# K8s-101



This is a small practice to demostrate the common knowledge of k8s.



## Requirements



Before you start this practice, please make sure you have `Docker` on your laptop/workstation.

If you are using `Docker Desktop`, please increase the RAM of your VM to at least 8GB.





## Kubectl



To access the kubernetes cluster, you will need to use the tool called `kubectl`.

About how to install thehttps://kubernetes.io/docs/tasks/tools/#kubectl

 `kubectl`, please check the official doc here:





## Helm



To deploy services to k8s, we will use `Helm` to help us. Please following the doc to setup the helm CLI first.

https://helm.sh/docs/intro/install/#from-script





## Kind



The `Kind`, means **Kubernetes IN Docker**. Is a tool for running local Kubernetes clusters using Docker container. This is a very convenience tool when you develop application on k8s.



To install the `Kind` CLI, please follow the instruction from its official doc:

https://kind.sigs.k8s.io/docs/user/quick-start/#installation



After you finish the installation, please use following command to create the k8s cluster at your local env:

```shell

kubeclt 

```



After the kind node is ready, you can use the `kubectl` to access your local cluster.

```shell

kubectl get nodes

```



## Nginx Ingress



In the kubernetes, we will expose the service to the external network by using `NodePort` or `Ingress`. Usually, we only use the Ingress if your service is a HTTP based service.

There are many solution to use Ingress with your kubernetes. In this practice, we will use `Nginx Ingress`.



To install the Nginx Ingrss, we will use its helm chart. Here is the command to install the Nginx ingress with helm chart:

```shell

helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx

helm repo update

helm install ingress-nginx ingress-nginx/ingress-nginx \

    --namespace ingress-nginx \

    --create-namespace \

    --set controller.hostPort.enabled=true

kubectl -n ingress-nginx rollout status deployment ingress-nginx-controller









list 

helm ls -n ingress-nginx


delete ingress-nginx in default namspace

helm delete ingress-nginx


🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰🥰

helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx && \
helm repo update && \
helm install ingress-nginx ingress-nginx/ingress-nginx \
	--namespace ingress-nginx \
    --create-namespace \
    --set controller.hostPort.enabled=true && \
kubectl -n ingress-nginx rollout status deployment ingress-nginx-controller


Rancher
cert-manager
helm repo add jetstack https://charts.jetstack.io

helm repo update

kubectl apply --validate=false -f https://github.com/jetstack/cert-manager/releases/download/v0.15.0/cert-manager.crds.yaml

kubectl create namespace cert-manager

helm install cert-manager jetstack/cert-manager \
    --namespace cert-manager \
    --version v0.15.0

kubectl -n cert-manager rollout status deploy/cert-manager
install rancher
helm repo add rancher-latest https://releases.rancher.com/server-charts/latest
helm repo update


kubectl create namespace cattle-system
helm install rancher rancher-latest/rancher \
    --namespace cattle-system \
    --set hostname=rancher.127.0.0.1.nip.io \
    --set replicas=1


kubectl -n cattle-system rollout status deploy/rancher




kubectl get svc -n ingress-nginx

#show
NAME                                 TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)                      AGE
ingress-nginx-controller             LoadBalancer   10.96.27.132   <pending>     80:32212/TCP,443:31798/TCP   6m11s
ingress-nginx-controller-admission   ClusterIP      10.96.92.107   <none>        443/TCP                      6m11s




小幫手 設定在.zshrc

alias k="/usr/local/bin/kubectl"


就可以使用 k

k get no(po/deploy/svc)

