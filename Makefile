
.PHONY: service

gokit:
	pushd cmd/go_kit_sample && GOOS=linux go build && popd
	docker build -f build/go_kit_sample/Dockerfile -t yangand/go_kit_sample .

service:
	pushd cmd/service && GOOS=linux go build && popd ; \
	docker build -f build/service/Dockerfile -t prometheus_sample_service .
	
up:
	docker-compose -f deployments/docker-compose.yaml up -d
	
down:
	docker-compose -f deployments/docker-compose.yaml down