LOCAL_OS=$(shell go tool dist banner | grep 'Go' | cut -d' ' -f 4 | cut -d/ -f 1)
LOCAL_ARCH=$(shell go tool dist banner | grep 'Go' | cut -d' ' -f 4 | cut -d/ -f 2)
GO_FILES=$(shell find . -name '*.go' | tr '\n' ' ')

build-static:
	$(MAKE) -C srv/updatemgr-web/

updatemgr: .pretty build-static $(GO_FILES)
	go build -o updatemgr


.PHONY: docker
docker: release/updatemgr.linux.amd64
	docker build -t updatemgr:latest .

.PHONY: docker-clean
docker-clean:
	docker image prune -f

.PHONY: release-local
release-local: release/updatemgr.$(LOCAL_OS).$(LOCAL_ARCH)

.PHONY: dist
dist: release/updatemgr.linux.arm \
	release/updatemgr.linux.arm64 \
	release/updatemgr.linux.amd64

release:
	mkdir release

release/updatemgr.linux.amd64: export CGO_ENABLED = 0
release/updatemgr.linux.amd64: export GOARCH = amd64
release/updatemgr.linux.amd64: export GOOS = linux
release/updatemgr.linux.amd64: .pretty $(GO_FILES) build-static release
	go build -ldflags="-s -w" -o $@
	# upx --brute $@

release/updatemgr.linux.arm: export CGO_ENABLED = 0
release/updatemgr.linux.arm: export GOARCH = arm
release/updatemgr.linux.arm: export GOOS = linux
release/updatemgr.linux.arm: .pretty $(GO_FILES) build-static release
	go build -ldflags="-s -w" -o $@
	upx --brute $@

release/updatemgr.linux.arm64: export CGO_ENABLED = 0
release/updatemgr.linux.arm64: export GOARCH = arm64
release/updatemgr.linux.arm64: export GOOS = linux
release/updatemgr.linux.arm64: .pretty $(GO_FILES) build-static release
	go build -ldflags="-s -w" -o $@
	upx --brute $@

pretty: .pretty

.pretty: $(GO_FILES)
	find . -name "*.go" -print0 | xargs -0 goimports -w
	touch .pretty

.PHONY: rek8s
rek8s: docker
	$(MAKE) -C k8s clean deploy

.PHONY: clean
clean:
	rm -fr updatemgr

.PHONY: dist-clean
dist-clean: clean
	rm -fr release
