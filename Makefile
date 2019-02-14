
ORG=ismacaulay
APP=chip8
DOCKER_MNT=/go/src/github.com/ismacaulay/chip8
MOCKGEN = $(shell go env GOPATH)/bin/mockgen

.PHONY: build test run coverage

image:
	docker build -t $(ORG)/$(APP) -f Dockerfile .

image-cli:
	docker build -t $(ORG)/$(APP)-cli  -f Dockerfile.cli .

image-tui:
	docker build -t $(ORG)/$(APP)-tui  -f Dockerfile.tui .

build:
	docker run --rm \
		-v $(shell pwd):$(DOCKER_MNT) \
		$(ORG)/$(APP) go build ./...

build-cli:
	docker run --rm \
		-v $(shell pwd):$(DOCKER_MNT) \
		-w $(DOCKER_MNT)/cmd/cli \
		$(ORG)/$(APP) go build

build-tui:
	docker run --rm \
		-v $(shell pwd):$(DOCKER_MNT) \
		-w $(DOCKER_MNT)/cmd/tui \
		$(ORG)/$(APP) go build

run-cli:
	docker run --rm -it \
		-v $(shell pwd)/cmd/cli/cli:/usr/local/bin/$(APP)-cli \
		--name $(APP)-cli \
		$(ORG)/$(APP)-cli

run-tui:
	docker run --rm -it \
		-v $(shell pwd)/cmd/tui/tui:/usr/local/bin/$(APP)-tui \
		--name $(APP)-tui \
		$(ORG)/$(APP)-tui
test:
	docker run --rm \
		-v $(shell pwd):$(DOCKER_MNT) \
		-w $(DOCKER_MNT) \
		$(ORG)/$(APP) go test ./... -cover

coverage:
	docker run --rm \
		-v $(shell pwd):$(DOCKER_MNT) \
		-w $(DOCKER_MNT) \
		$(ORG)/$(APP) go test ./... -coverprofile cover.out
	go tool cover -html=cover.out

mocks:
	$(MOCKGEN) -source=pkg/emu/registers/registers.go -destination=pkg/emu/registers/mock/registers_mock.go
	$(MOCKGEN) -source=pkg/emu/display/display.go -destination=pkg/emu/display/mock/display_mock.go
	$(MOCKGEN) -source=pkg/emu/keyboard/keyboard.go -destination=pkg/emu/keyboard/mock/keyboard_mock.go
	$(MOCKGEN) -source=pkg/emu/memory/memory.go -destination=pkg/emu/memory/mock/memory_mock.go
