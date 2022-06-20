# Copyright 2022 Mike Messmore <mike@messmore.org>

# Just discover go files
GO_FILES=$(shell find . -name '*.go' | tr '\n' ' ')

# default to build locally
updatemgr: .pretty build-static $(GO_FILES)
	go build -o updatemgr

# (re)build static content
build-static:
	$(MAKE) -C srv/updatemgr-web/

# build all ready-to-go binaries for release
dist: release/updatemgr.linux.arm \
	release/updatemgr.linux.arm64 \
	release/updatemgr.linux.amd64

release:
	mkdir release

# This is rebuilt every time, which is weird
# So there's a filthy hack
# Don't bother compressing on intel
release/updatemgr.linux.amd64: export CGO_ENABLED = 0
release/updatemgr.linux.amd64: export GOARCH = amd64
release/updatemgr.linux.amd64: export GOOS = linux
release/updatemgr.linux.amd64: .pretty $(GO_FILES) build-static release
	@# I'm not sure why this isn't working right
	if ! [ -f $@ ]; then \
		go build -ldflags="-s -w" -o $@; \
	fi

# Same weird hack
# Compress on arm7 to save space on embedded
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

# Same weird hack
# Compress on arm64 to save space on embedded
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

# Friendly name to run by hand
pretty: .pretty

# Force formatting of go files, and tidy go mod stuff
.pretty: $(GO_FILES)
	find . -name "*.go" -print0 | xargs -0 goimports -w
	go mod tidy
	touch .pretty

# make our packages
# It's too complicated to figure out the resulting file
# name to bother having a proper dependency, so just
# phony this junk
.PHONY: package
package: release/updatemgr.linux.amd64 \
	release/updatemgr.linux.arm \
	release/updatemgr.linux.arm64

	sample/create_dpkg.sh

# clean local builds
.PHONY: clean
clean:
	rm -fr updatemgr

# clean release binaries and packages
.PHONY: dist-clean
dist-clean: clean
	rm -fr release
