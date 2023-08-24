#===============================================================================
# Valiables
#===============================================================================

#-------------------------------------------------------------------------------
# Application
#-------------------------------------------------------------------------------
APP_NAME := traffic-monitor
APP_REPO := github.com/ikmski/traffic-monitor

APP_VERSION := 0.1.0
APP_SOURCES := $(shell find . -name "*.go" -type f)

LDFLAGS :=
LDFLAGS += -s
LDFLAGS += -w

#===============================================================================
# Tasks
#===============================================================================

#-------------------------------------------------------------------------------
# Application
#-------------------------------------------------------------------------------

.PHONY: run build clean init test

run:
	go run ./...

build: bin/$(APP_NAME)

bin/$(APP_NAME): $(APP_SOURCES)
	go build \
		-a -v \
		-ldflags "$(LDFLAGS)" \
		-o $@ \

build-arm64: bin/$(APP_NAME)_arm64
bin/$(APP_NAME)_arm64: $(APP_SOURCES)
	GOOS=linux \
	GOARCH=arm64 \
	go build \
		-a -v \
		-ldflags "$(LDFLAGS)" \
		-o $@ \

clean:
	go clean -i ./...
	rm -rf bin/*

init:
	if [ ! -f go.mod ]; then \
		go mod init $(APP_REPO); \
	fi
	go get -u github.com/Songmu/make2help/cmd/make2help

test:
	go test -v $(APP_REPO)/...

download-deps:
	go mod download

update-deps:
	go get -u ./...
	go mod tidy

