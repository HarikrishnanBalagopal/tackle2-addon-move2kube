BUILD_DIR=bin
BIN_NAME=tackle-addon-move2kube

.PHONY: clean
clean:
	rm -rf ${BUILD_DIR}/${BIN_NAME}

.PHONY: build
build: ${BUILD_DIR}/${BIN_NAME}

${BUILD_DIR}/${BIN_NAME}:
	go build -o ${BUILD_DIR}/${BIN_NAME}

.PHONY: cbuild
cbuild:
	docker build -t quay.io/konveyor/tackle2-addon-move2kube:latest -f Dockerfile .

.PHONY: cpush
cpush:
	docker tag quay.io/konveyor/tackle2-addon-move2kube:latest quay.io/hari_balagopal/tackle2-addon-move2kube:latest
	docker push quay.io/hari_balagopal/tackle2-addon-move2kube:latest
