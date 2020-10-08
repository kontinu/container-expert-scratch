# CI/CD

- [BUILD] Docker (app,test suite) => image:v1.1.100
- [TEST] docker-compose, kubernetes (namespace: testing)
- [DEPLOY]:
  -  edit image in YAML manifests/: image:v1.1.100
  -  kubectl apply -f manifests/
