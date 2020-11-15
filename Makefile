all: build

test:
	cd ${CI_PROJECT_DIR} ;\
	go test ./... -coverprofile cover.out

build:
	cd ${CI_PROJECT_DIR} ;\
    go mod download ;\
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o ${module} cmd/main.go ;\

package:
	cd ${CI_PROJECT_DIR}; \
    cp ${CI_PROJECT_DIR}/install/Dockerfile ${RELEASE_PATH} ;\
    cp ${CI_PROJECT_DIR}/install/entrypoint.sh ${RELEASE_PATH} ;\
    docker build -t ${imagename} ${RELEASE_PATH} ;\
    docker push ${imagename}

push-meta:
	cd ${CI_PROJECT_DIR} ;\
    tar czvf install/${module}.meta.tar.gz -C install/ . ;\
	cd install/ && push_meta

deploy:
