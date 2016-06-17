all: build push deploy_stage

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags "-w" -o bin/kube-monitor main.go
	docker build . | tail -n 1 | awk '{print $$3}' | xargs echo -n > .last_build

push:
	docker tag $$(cat .last_build) vungle/kube-monitor:$$(cat .last_build) 
	docker push vungle/kube-monitor:$$(cat .last_build)

deploy_prod:
	sed 's/DEPLOY_TAG/'"$$(cat .last_build)"'/g' deployment.yaml > /tmp/deployment.yaml
	kubectl apply -f /tmp/deployment.yaml
	kubectl get rs -o wide | grep kube-monitor
	kubectl get pods -o wide --all-namespaces | grep kube-monitor
	

deploy_stage:
	sed 's/DEPLOY_TAG/'"$$(cat .last_build)"'/g' deployment.yaml > /tmp/deployment.yaml
	kubectl apply -f /tmp/deployment.yaml
	kubectl get rs -o wide | grep kube-monitor
	kubectl get pods -o wide --all-namespaces | grep kube-monitor
