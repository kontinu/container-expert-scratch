
image:=kontinu/kontinu-operator:1.0
.PHONY: build push all crd demo deploy undo full-demo kontinu
build:
	operator-sdk build ${image}
	docker push ${image}

crd:
	kubectl apply -f deploy/crds/helm.operator-sdk_kontinus_crd.yaml


demo: crd
	operator-sdk run local --reconcile-period="1s"

deploy:
	kubectl apply -f deploy/crds/helm.operator-sdk_v1alpha1_kontinu_cr.yaml
kontinu: deploy

full-demo: crd
	kubectl apply -f deploy/
	${MAKE} deploy


undo:
	kubectl delete -f deploy/crds/helm.operator-sdk_v1alpha1_kontinu_cr.yaml || true
	kubectl delete -f deploy/crds/helm.operator-sdk_kontinus_crd.yaml || true
	kubectl delete -f deploy/ || true
