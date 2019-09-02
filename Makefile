
VERSION ?= $(shell stty echo; read -p "Version: " pwd; stty echo; echo $$pwd)

test:
	@echo Running tests (with -race flag on) 
	@go test ./... -race

generate:
	@echo Generating commands code: runtime, doc, etc.
	@go generate gen/aws/generators/main.go

build: generate
	@echo Building application binary
	@go build -o acentera
build-only:
	@echo Building application binary
	@go build -o acentera

release:
	@echo Release version
	@echo export VERSION=$(VERSION) && sh ./release.sh
