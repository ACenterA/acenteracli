
test:
	@echo Running tests (with -race flag on) 
	@go test ./... -race

generate:
	@echo Generating commands code: runtime, doc, etc.
	@go mod vendor
	@go generate gen/aws/generators/main.go

build: generate
	@echo Building application binary
	@go build -mod=mod -o acentera || (go get -u golang.org/x/tools/cmd/goimports && go build -o acentera)

build-only:
	@echo Building application binary
	@go build -mod=mod -o acentera

build-watch:
	@echo Watch/Build application binary
	@exec watchexec -e go -r -- go build -o .acentera

release:
	@echo Release version
	@sh ./release.sh
