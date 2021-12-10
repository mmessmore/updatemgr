LOCAL_OS=$(shell go tool dist banner | grep 'Go' | cut -d' ' -f 4 | cut -d/ -f 1)
LOCAL_ARCH=$(shell go tool dist banner | grep 'Go' | cut -d' ' -f 4 | cut -d/ -f 2)
GO_FILES=$(shell find . -maxdepth 1 -name '*.go' | tr '\n' ' ')

updatemgr: .pretty $(GO_FILES)
	go build -o updatemgr

.PHONY: run
run: updatemgr
	./updatemgr

.PHONY: release-local
release-local: release/updatemgr.$(LOCAL_OS).$(LOCAL_ARCH)


.PHONY: release
# release: release/updatemgr.linux.amd64 release/updatemgr.linux.arm release/updatemgr.darwin.amd64
release: release/updatemgr.linux.arm release/updatemgr.darwin.amd64

release/updatemgr.linux.amd64: .pretty $(GO_FILES)
	mkdir -p release
	GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o $@
	upx --brute $@

release/updatemgr.linux.arm: .pretty $(GO_FILES)
	mkdir -p release
	GOARCH=arm GOOS=linux go build -ldflags="-s -w" -o $@
	upx --brute $@

release/updatemgr.darwin.amd64: .pretty $(GO_FILES)
	mkdir -p release
	GOARCH=amd64 GOOS=darwin go build -ldflags="-s -w" -o $@
	upx --brute $@

pretty: .pretty

.pretty: $(GO_FILES)
	find . -name "*.go" -print0 | xargs -0 goimports -w
	touch .pretty

.PHONY: clean
clean:
	rm -fr updatemgr release
