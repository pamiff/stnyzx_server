#Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build

OUTPUT=./bin

all: build
build:
	$(GOBUILD) -o $(OUTPUT)/auth ./auth/cmd