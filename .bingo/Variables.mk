# Auto generated binary variables helper managed by https://github.com/bwplotka/bingo v0.6. DO NOT EDIT.
# All tools are designed to be build inside $GOBIN.
BINGO_DIR := $(dir $(lastword $(MAKEFILE_LIST)))
GOPATH ?= $(shell go env GOPATH)
GOBIN  ?= $(firstword $(subst :, ,${GOPATH}))/bin
GO     ?= $(shell which go)

# Below generated variables ensure that every time a tool under each variable is invoked, the correct version
# will be used; reinstalling only if needed.
# For example for calens variable:
#
# In your main Makefile (for non array binaries):
#
#include .bingo/Variables.mk # Assuming -dir was set to .bingo .
#
#command: $(CALENS)
#	@echo "Running calens"
#	@$(CALENS) <flags/args..>
#
CALENS := $(GOBIN)/calens-v0.2.0
$(CALENS): $(BINGO_DIR)/calens.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/calens-v0.2.0"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=calens.mod -o=$(GOBIN)/calens-v0.2.0 "github.com/restic/calens"

GOLINT := $(GOBIN)/golint-v0.0.0-20210508222113-6edffad5e616
$(GOLINT): $(BINGO_DIR)/golint.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/golint-v0.0.0-20210508222113-6edffad5e616"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=golint.mod -o=$(GOBIN)/golint-v0.0.0-20210508222113-6edffad5e616 "golang.org/x/lint/golint"

STATICCHECK := $(GOBIN)/staticcheck-v0.3.1
$(STATICCHECK): $(BINGO_DIR)/staticcheck.mod
	@# Install binary/ries using Go 1.14+ build command. This is using bwplotka/bingo-controlled, separate go module with pinned dependencies.
	@echo "(re)installing $(GOBIN)/staticcheck-v0.3.1"
	@cd $(BINGO_DIR) && $(GO) build -mod=mod -modfile=staticcheck.mod -o=$(GOBIN)/staticcheck-v0.3.1 "honnef.co/go/tools/cmd/staticcheck"

