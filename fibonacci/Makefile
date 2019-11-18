# Collect all functions that should be built and deployed.
cmds := $(shell find cmd -name \*main.go | awk -F'/' '{print $$2}')

.PHONY: build clean deploy gomodgen

build: gomodgen # Build golang libraries.
	export GO111MODULE=on
	@for cmd in $(cmds) ; do \
		env GOOS=linux go build -ldflags="-s -w" -o bin/$$cmd cmd/$$cmd/main.go ; \
	done

clean: # Remove golang binaries.
	rm -rf ./bin ./vendor Gopkg.lock

gomodgen: # Generate golang module.
	chmod u+x gomod.sh
	./gomod.sh
