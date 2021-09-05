APP_NAME?=fizz
CORE_DEPLOYMENT?=fizz-deployment
YML_FILE?=deployments.yml
DOCKER_HUB?=tarek5
RANDOM?=$(shell bash -c 'echo $$RANDOM')

dockerize:
	docker build -t tarek5/fizz .
	docker push tarek5/fizz:latest

kube-deploy-app: dockerize
	kubectl apply -f $(YML_FILE)

kube-delete:
	kubectl delete -f $(YML_FILE)

kube-rollout:
	kubectl rollout restart deployment/$(CORE_DEPLOYMENT)

kube-patch:
	kubectl patch deployment $(CORE_DEPLOYMENT) -p "{\"spec\": {\"template\": {\"metadata\": {\"labels\": {\"build\":\"$(RANDOM)\"}}}}}"

kube-scale-down:
	kubectl scale deploy worker-deployment --replicas=1

docker-run:
	docker-compose up -d

docker-down:
	docker-compose down

prometheus:
	cp prometheus-fizz.yaml /tmp
	docker run -d -p 9090:9090 -v /tmp:/etc/prometheus prom/prometheus --config.file=/etc/prometheus/prometheus-fizz.yaml --storage.tsdb.path=/prometheus --web.console.libraries=/usr/share/prometheus/console_libraries --web.console.templates=/usr/share/prometheus/consoles

grafana:
	docker run -d -p 3000:3000 grafana/grafana

loki:
	cp loki-config.yaml /tmp
	cp promtail-config.yaml /tmp
	docker run -v /tmp:/mnt/config -p 3100:3100 grafana/loki -config.file=/mnt/config/loki-config.yaml
	docker run -v /tmp:/mnt/config -v /var/log:/var/log grafana/promtail -config.file=/mnt/config/promtail-config.yaml

