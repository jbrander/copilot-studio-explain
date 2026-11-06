# cs-explain: a `kubectl explain` for Copilot Studio YAML (Go module in cli/).
BIN ?= bin/cse

.PHONY: cli schemas test golden install clean

## cli: sync the vendored schemas into the embed dir and build bin/cs-explain
cli: schemas
	cd cli && go build -trimpath -o ../$(BIN) .

## schemas: copy vendor/skills-for-copilot-studio/reference/* into cli/internal/schemas/data
schemas:
	cd cli && go generate .

## test: vet and test the CLI
test:
	cd cli && go vet ./... && go test ./...

## golden: re-record the golden outputs after an intended output change
golden:
	cd cli && go test ./cmd -run TestGolden -update

## install: build cs-explain into $(go env GOPATH)/bin
install: schemas
	cd cli && go build -trimpath -o "$$(go env GOPATH)/bin/cs-explain" .

clean:
	rm -rf bin
