.PHONY: proto kubernetes-redeploy prometheus

APP_NAME?=saas
CORE_DEPLOYMENT?=saas-deployment
YML_FILE?=deployments.yml
DOCKER_HUB?=tarek5
RANDOM?=$(shell bash -c 'echo $$RANDOM')

generate-mock:
	mockgen -source=internal/services/authservice/service.go -destination=internal/mocks/auth_service_mock.go -package=mocks

test:
	go test -v ./...

build-worker:
	GOOS=windows GOARCH=amd64 go build saas-boilerplate/cmd/worker

test-cover:
	go test -coverprofile=cover.out -coverpkg=github.com/cyantarek/go-couchbase-url-shortener/saas-boilerplate/server

deploy-manual:
	ssh ubuntu@3.17.133.93 'cd /home/ubuntu/lowdly-app && git pull origin master && make docker-compose'

dockerize:
	docker build -t tarek5/fizz .
	docker push tarek5/fizz:latest

run-docker-db:
	docker-compose -f docker-compose.dev.yml up -d db-redis db-postgres pgadmin redisadmin

docker-compose:
	docker-compose -f docker-compose.yml up -d --build

docker-compose-local:
	docker-compose -f docker-compose.yml up -d --build db-redis db-postgres adminer nsqlookupd nsqd nsqadmin

sentry:
	docker-compose -f docker-compose.staging.yml exec sentry sentry upgrade
	docker-compose -f docker-compose.staging.yml exec sentry sentry createuser

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

SERVICES := auth order product
# SERVICES := auth payment content upload contact

proto:
	for service in $(SERVICES); do \
  		cd api/proto && \
  		protoc --go_out=../../../ \
		--go-grpc_out=../../../ \
		--grpc-gateway_out=logtostderr=true:../../../ \
  		"$$service"_*.proto && cd ../..; \
    done

gql:
	gqlgen generate

docker-run:
	docker-compose up -d

docker-down:
	docker-compose down

scaffold:
	mkdir config
	mkdir assets
	mkdir -p cmd/app
	mkdir -p pkg/{gql,proto}
	mkdir -p api/{gql,proto}
	mkdir -p internal/{core,outside,pkg}
	mkdir -p internal/core/{application,domain,port}
	mkdir -p internal/core/application/dto
	mkdir -p internal/core/port/{imcoming,outgoing}
	mkdir -p internal/outside/adapter/{driven,driving}
	mkdir -p internal/outside/adapter/driving/{gqlhandler,grpchandler,httphandler}
	mkdir -p internal/pkg/{transports,jwtpacker,helpers,logger}
	mkdir -p internal/pkg/transports/{graphql,grpc,grpcgateway,grpcweb,http,middlewares}

	cp $$HOME/boilerplate/config.txt config/config.go
	cp $$HOME/boilerplate/config_instance.txt config/config_instance.go

deploy-helm:
	helm install --generate-name fizzcharts/

delete-helm:
	helm ls --all --short | xargs -L1 helm delete

prometheus:
	cp prometheus-fizz.yaml /tmp
	docker run -d -p 9090:9090 -v /tmp:/etc/prometheus prom/prometheus --config.file=/etc/prometheus/prometheus-fizz.yaml --storage.tsdb.path=/prometheus --web.console.libraries=/usr/share/prometheus/console_libraries --web.console.templates=/usr/share/prometheus/consoles

grafana:
	docker run -d -p 3000:3000 grafana/grafana
