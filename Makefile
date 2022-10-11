all: build

test:
	cd ${CI_PROJECT_DIR} ;\
	go test ./... ;\
    go test -v -test.short -coverprofile=coverage.out -timeout=10s `go list ./... |grep -v cmd |grep -v docs` -json > report.json ;\

build:
	cd ${CI_PROJECT_DIR} ;\
	docker build -t ${imagename} .;\
    docker push ${imagename};\

gen-openapi-docs:
	cd ${CI_PROJECT_DIR}; \
	swag init \
		--parseDependency \
		--parseInternal \
		-d pkg/sage/restfulapi/ \
		-g app_config.go \
