# available images

- kontinu/nginx-menu:v1.0
- kontinu/nodejs-expert:v1.0
- kontinu/python-expert:v1.0
- kontinu/go-expert:v1.0

# manifests


# Helm

## Local Dev


```bash
helm create chart


#check templates


helm dep up chart

helm upgrade --install kontinu kontinu-chart

```

## Package and Share.


```bash
$ helm package $YOUR_CHART_PATH/ # build the tgz file and copy it here
$ helm repo index . # create or update the index.yaml for repo
$ git add .
$ git commit -m 'New chart version'

```


```bash

$ helm repo add sample 'https://raw.githubusercontent.com/your-repo'
$ helm repo update
$ helm search kontinu
```
