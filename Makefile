GOCMD=go
GOGENRATE=$(GOCMD) generate

all: generate

generate: deps
		$(GOGENRATE)

deps:
		dep ensure
