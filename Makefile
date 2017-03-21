SHORT_NAME ?= minio-broker
IMAGE_PREFIX ?= deis
VERSION ?= git-$(shell git rev-parse --short HEAD)
IMAGE := quay.io/${IMAGE_PREFIX}/${SHORT_NAME}:${VERSION}
REPO_PATH := github.com/deis/${SHORT_NAME}

MUTABLE_VERSION ?= canary
MUTABLE_IMAGE := quay.io/${IMAGE_PREFIX}/${SHORT_NAME}:${MUTABLE_VERSION}

DEV_ENV_IMAGE := quay.io/deis/go-dev:0.20.0
DEV_ENV_WORK_DIR := /go/src/${REPO_PATH}
DEV_ENV_OPTS := --rm -v ${CURDIR}:${DEV_ENV_WORK_DIR} -w ${DEV_ENV_WORK_DIR}
DEV_ENV_CMD := docker run ${DEV_ENV_OPTS} ${DEV_ENV_IMAGE}
LDFLAGS := "-s -X main.version=${VERSION}"
BINARY_DEST_DIR = rootfs/opt/minio-broker/sbin
PACKAGES = $(go list $(glide novendor))

info:
	@echo "Build tag:       ${VERSION}"
	@echo "Immutable tag:   ${IMAGE}"
	@echo "Mutable tag:     ${MUTABLE_IMAGE}"

install-deps:
	curl -SsL https://github.com/Masterminds/glide/releases/download/v0.12.3/glide-v0.12.3-linux-amd64.tar.gz | tar xz && mv linux-amd64/glide ../bin && rm -rf linux-amd64
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

bootstrap:
	glide install

build: build-binary-in-container build-image
push: docker-push


run-interactive:
	docker run -it ${IMAGE} bash

build-binary:
	GOOS=linux GOARCH=amd64 go build -ldflags ${LDFLAGS} -o $(BINARY_DEST_DIR)/$(SHORT_NAME) .

build-binary-in-container:
	${DEV_ENV_CMD} make build-binary

build-image:
	docker build \
		--pull \
	 	--build-arg VERSION=${VERSION} \
	 	--build-arg BUILD_DATE=`date -u +'%Y-%m-%dT%H:%M:%SZ'` \
		-t ${IMAGE} rootfs
	docker tag ${IMAGE} ${MUTABLE_IMAGE}

docker-push: docker-mutable-push docker-immutable-push

docker-immutable-push:
	docker push ${IMAGE}

docker-mutable-push:
	docker push ${MUTABLE_IMAGE}


ci: test

test: test-unit

test-unit:
	${DEV_ENV_CMD} ginkgo -r .
