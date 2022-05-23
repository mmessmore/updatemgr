LOCAL_OS=$(shell go tool dist banner | grep 'Go' | cut -d' ' -f 4 | cut -d/ -f 1)
LOCAL_ARCH=$(shell go tool dist banner | grep 'Go' | cut -d' ' -f 4 | cut -d/ -f 2)
GO_FILES=$(shell find . -name '*.go' | tr '\n' ' ')

updatemgr: .pretty build-static $(GO_FILES)
	go build -o updatemgr

build-static:
	$(MAKE) -C srv/updatemgr-web/

.PHONY: run
run: updatemgr
	./updatemgr serve

dist: release/updatemgr.linux.arm \
	release/updatemgr.linux.arm64 \
	release/updatemgr.linux.amd64

release:
	mkdir release

release/updatemgr.linux.amd64: export CGO_ENABLED = 0
release/updatemgr.linux.amd64: export GOARCH = amd64
release/updatemgr.linux.amd64: export GOOS = linux
release/updatemgr.linux.amd64: .pretty $(GO_FILES) build-static release
	@# I'm not sure why this isn't working right
	if ! [ -f $@ ]; then \
		go build -ldflags="-s -w" -o $@; \
	else \
		echo skipping build; \
	fi

release/updatemgr.linux.arm: export CGO_ENABLED = 0
release/updatemgr.linux.arm: export GOARCH = arm
release/updatemgr.linux.arm: export GOOS = linux
release/updatemgr.linux.arm: .pretty $(GO_FILES) build-static release
	@# Compress for tiny computers
	@# I'm not sure why this isn't working right
	if ! [ -f $@ ]; then\
		go build -ldflags="-s -w" -o $@; \
		upx --brute $@; \
	fi

release/updatemgr.linux.arm64: export CGO_ENABLED = 0
release/updatemgr.linux.arm64: export GOARCH = arm64
release/updatemgr.linux.arm64: export GOOS = linux
release/updatemgr.linux.arm64: .pretty $(GO_FILES) build-static release
	@# Compress for tiny computers
	@# I'm not sure why this isn't working right
	if ! [ -f $@ ]; then\
		go build -ldflags="-s -w" -o $@; \
		upx --brute $@; \
	fi

pretty: .pretty

.pretty: $(GO_FILES)
	find . -name "*.go" -print0 | xargs -0 goimports -w
	touch .pretty

.PHONY: package
package: release/updatemgr.linux.amd64 \
	release/updatemgr.linux.arm \
	release/updatemgr.linux.arm64

	sample/create_dpkg.sh


.PHONY: clean
clean:
	rm -fr updatemgr

.PHONY: dist-clean
dist-clean: clean
	rm -fr release
