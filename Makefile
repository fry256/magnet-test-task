PWD = $(shell pwd)
UID = $(shell id -u)
GID = $(shell id -g)
SWAGGER := docker run --rm -it  --user $(UID):$(GID) -e GOPATH=$(HOME)/go:/go -v $(HOME):$(HOME) -w $(PWD) quay.io/goswagger/swagger

gen:
	@rm -rf ./cmd
	@rm -rf ./internal
	@mkdir -p cmd
	@mkdir -p internal/generated
	@mkdir -p internal/app
	@mkdir -p internal/config
	@$(SWAGGER) generate server \
		-f ./swagger-api/swagger.yaml \
		-t ./internal/generated -C ./swagger-templates/default-server.yaml \
		--template-dir ./swagger-templates/templates \
		--name magnet-test-task

swagger:
	docker run --rm -p 8088:8080 -e BASE_URL=/swagger -e URLS='[\
		{url:"/docs/swagger.yaml", name:"API v1"},\
	]' -v $(PWD)/docs:/etc/nginx/html/docs swaggerapi/swagger-ui

prune:
	@docker volume prune -f

build:
	docker-compose -f build/dev/docker-compose.yml build -q --parallel --progress plain

up:
	docker-compose -f build/dev/docker-compose.yml up

run: prune build up

stop:
	docker-compose -f build/dev/docker-compose.yml down
