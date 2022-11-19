SHELL := /bin/bash

README_TEMPLATE_FILE=./templates/README.md.gotmpl

# List of targets the `readme` target should call before generating the readme
export README_DEPS ?= docs/targets.md docs/terraform.md

-include $(shell curl -sSL -o .build-harness "https://cloudposse.tools/build-harness"; echo .build-harness)

## Lint terraform code
lint:
	$(SELF) terraform/install terraform/get-modules terraform/lint terraform/validate

## Install dependencies
install: terraform/install
	@exit 0

## Run all tests against terraform module
test: test/bats test/unit
	@exit 0

## Run bats tests against terraform module
test/bats:
	(cd ./test && make)

## Run unit tests against terraform module
test/unit:
	(cd ./test/src && make test)