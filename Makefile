GOCOMMAND=go
GOBUILD=$(GOCOMMAND) build
GOCLEAN=$(GOCOMMAND) clean
BINARYNAME=dessert

all:	build
build:
	$(GOBUILD)
clean:
	$(GOCLEAN)
