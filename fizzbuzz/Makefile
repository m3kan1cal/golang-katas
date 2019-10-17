# Collect all functions that should be built and deployed.
functions := $(shell find functions -name \*main.go | awk -F'/' '{print $$2}')

.PHONY: build clean deploy gomodgen

build: gomodgen # Build golang libraries.
	export GO111MODULE=on
	@for function in $(functions) ; do \
		env GOOS=darwin go build -ldflags="-s -w" -o bin/$$function functions/$$function/main.go ; \
	done

clean: # Remove golang binaries.
	rm -rf ./bin ./vendor Gopkg.lock

gomodgen: # Generate golang module.
	chmod u+x gomod.sh
	./gomod.sh
