# Start minikube


```
minikube start --vm=true --driver=hyperkit
```

```
minikube addons enable dashboard
minikube addons enable ingress
```

```
marcoscano @ blackpearl ~
 └─  🐳  minikube profile list
|----------|-----------|---------|--------------|------|---------|---------|
| Profile  | VM Driver | Runtime |      IP      | Port | Version | Status  |
|----------|-----------|---------|--------------|------|---------|---------|
| minikube | hyperkit  | docker  | 192.168.64.3 | 8443 | v1.19.2 | Running |
|----------|-----------|---------|--------------|------|---------|---------|

```
