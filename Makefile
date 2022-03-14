LOCAL_OS=$(shell go tool dist banner | grep 'Go' | cut -d' ' -f 4 | cut -d/ -f 1)
LOCAL_ARCH=$(shell go tool dist banner | grep 'Go' | cut -d' ' -f 4 | cut -d/ -f 2)
GO_FILES=$(shell find . -name '*.go' | tr '\n' ' ')

build: build-static updatemgr

build-static:
	$(MAKE) -C srv/updatemgr-web/

updatemgr: .pretty $(GO_FILES)
	go build -o updatemgr


.PHONY: docker
docker: release/updatemgr.linux.amd64
	docker build -t updatemgr:latest .

.PHONY: docker-clean
docker-clean:
	docker image prune -f

.PHONY: run
run: updatemgr
	./updatemgr

.PHONY: release-local
release-local: release/updatemgr.$(LOCAL_OS).$(LOCAL_ARCH)


.PHONY: release
# release: release/updatemgr.linux.amd64 release/updatemgr.linux.arm release/updatemgr.darwin.amd64
release: release/updatemgr.linux.arm release/updatemgr.darwin.amd64

release/updatemgr.linux.amd64: .pretty $(GO_FILES) build-static
	mkdir -p release
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o $@
	# upx --brute $@

release/updatemgr.linux.arm: .pretty $(GO_FILES) build-static
	mkdir -p release
	CGO_ENABLED=0 GOARCH=arm GOOS=linux go build -ldflags="-s -w" -o $@
	# upx --brute $@

release/updatemgr.darwin.amd64: .pretty $(GO_FILES) build-static
	mkdir -p release
	CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -ldflags="-s -w" -o $@
	# upx --brute $@

pretty: .pretty

.pretty: $(GO_FILES)
	find . -name "*.go" -print0 | xargs -0 goimports -w
	touch .pretty

.PHONY: rek8s
rek8s: docker
	$(MAKE) -C k8s clean deploy

.PHONY: clean
clean:
	rm -fr updatemgr release
