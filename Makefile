LDFLAGS = -s -w
GOFLAGS = -trimpath -ldflags="$(LDFLAGS)"

define build
	cd $(1) && \
		mkdir -p bin && \
		CGO_ENABLED=0 GOOS=linux   go build -o bin/$(1)-linux   $(GOFLAGS) && \
		CGO_ENABLED=0 GOOS=windows go build -o bin/$(1)-windows $(GOFLAGS) && \
		CGO_ENABLED=0 GOOS=darwin  go build -o bin/$(1)-darwin  $(GOFLAGS) 
endef

all: clean build

.PHONY: clean
clean:
	rm -rf **/bin

.PHONY: build
build:
	$(call build,hcl)
	$(call build,json)
	$(call build,toml)
	$(call build,yaml)