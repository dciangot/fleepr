VERSION?=`git describe --tags`
DOCBIN?=mkdocs
BUILD_DATE := `date +%Y-%m-%d\ %H:%M`
VERSIONFILE := version.go

GOCMD=go
GOBUILD=$(GOCMD) build -mod=vendor -installsuffix cgo -a -x -ldflags "-w -v"
GOBUILD_DBG=$(GOCMD) build -x -mod=vendor
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=fleepr
REPO=github.com/dciangot/fleepr

export GO111MODULE=on
# Force 64 bit architecture
export GOARCH=amd64

all: build test

build:
	$(GOBUILD) -o $(BINARY_NAME)

build-debug:
	env CGO_ENABLED=0 $(GOBUILD_DBG) -o $(BINARY_NAME) -v

doc:
	cp README.md docs/README.md
	BUILD_DOC=true ./$(BINARY_NAME)

publish-doc:
	cp README.md docs/README.md
	$(DOCBIN) gh-deploy

test: build
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

run: build
	./$(BINARY_NAME)

install:
	$(GOCMD) install $(REPO)

tidy:
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor

docker-bin-build:
	docker run --rm -it -v ${PWD}:/go -w /go/ golang:1.13.6 go build -mod vendor -o "$(BINARY_NAME)" -v

docker-img-build:
	docker build . -t dodas

windows-build:
	env GOOS=windows CGO_ENABLED=0 $(GOBUILD) -mod vendor -o $(BINARY_NAME).exe -v

macos-build:
	env GOOS=darwin CGO_ENABLED=0 $(GOBUILD) -mod vendor -o $(BINARY_NAME)_osx -v

gensrc:
	rm -f $(VERSIONFILE)
	@echo "package main" > $(VERSIONFILE)
	@echo "const (" >> $(VERSIONFILE)
	@echo "  VERSION = \"$(VERSION)\"" >> $(VERSIONFILE)
	@echo "  BUILD_DATE = \"$(BUILD_DATE)\"" >> $(VERSIONFILE)
	@echo ")" >> $(VERSIONFILE)

build-release: tidy gensrc build doc publish-doc test windows-build macos-build
	zip dodas.zip dodas
	zip dodas.exe.zip dodas.exe
	zip dodas_osx.zip dodas_osx
