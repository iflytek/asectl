GO_VERSION=$(shell go env GOVERSION)
GIT_COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
BUILD_TIME=$(shell date)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
GOBINPATH=$(shell which go)
GO_PATH=$(abspath ../../../..)

BUILD_COMMAND="$(GOBINPATH) build -v -ldflags \"-X 'github.com/xfyun/asectl/cmd/version.goVersion=$(GO_VERSION)' -X 'github.com/xfyun/asectl/cmd/version.gitCommit=$(GIT_COMMIT)' -X 'github.com/xfyun/asectl/cmd/version.branch=$(BRANCH)' -X 'github.com/xfyun/asectl/cmd/version.buildTime=$(BUILD_TIME)' -X 'github.com/xfyun/asectl/cmd/version.oSArch=$(GOOS)/$(GOARCH)'\" ."


all:	export GOPATH=$(GO_PATH)
all:
	@echo "begin build asectl ................"
	@echo ""
	@echo "Go version: " $(GO_VERSION)
	@echo "Git commit: " $(GIT_COMMIT)
	@echo "Branch:     " $(BRANCH)
	@echo "Build time: " $(BUILD_TIME)
	@echo "OS/Arch:    " $(GOOS)/$(GOARCH)
	@echo "GOPATH:     " $(GO_PATH)

	@echo $(BUILD_COMMAND)
	@bash -c $(BUILD_COMMAND)

clean:
	rm -f asectl
