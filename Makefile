build:
	docker build -t junhuiwu/go-k8s-minimal:0.1.0 .

deploy:
	kubectl apply -f deploy/

restart:
	kubectl rollout restart deployment go-k8s-minimal