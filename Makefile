all: tidy test testcov cuefmtcheck lintcue cuegen dirtycheck lintinsights

tidy:
	@echo "  >  Tidying go.mod ..."
	@go mod tidy

test:
	@echo "  >  Running tests ..."
	@go vet ./...
	@go test ./...

testcov:
	@echo "Running tests and generating coverage output ..."
	@go test ./... -coverprofile coverage.out -covermode count
	@sleep 2 # Sleeping to allow for coverage.out file to get generated
	@echo "Current test coverage : $(shell go tool cover -func=coverage.out | grep total | grep -Eo '[0-9]+\.[0-9]+') %"

# Verify CUE formatting in ./schemas
cuefmtcheck:
	@echo "  >  Verifying CUE formatting in ./schemas ..."
	@cue fmt --check --files ./schemas

lintcue:
	@echo "  >  Linting CUE files ..."
	@cue eval ./schemas/layer-1.cue --all-errors --verbose
	@cue eval ./schemas/layer-2.cue --all-errors --verbose
	@cue eval ./schemas/layer-3.cue --all-errors --verbose
	@cue eval ./schemas/layer-4.cue --all-errors --verbose

cuegen:
	@echo "  >  Generating types from cue schema ..."
	@cue exp gengotypes ./schemas/layer-1.cue
	@mv cue_types_gen.go layer1/generated_types.go
	@cue exp gengotypes ./schemas/layer-2.cue
	@mv cue_types_gen.go layer2/generated_types.go
	@cue exp gengotypes ./schemas/layer-3.cue
	@mv cue_types_gen.go layer3/generated_types.go
	@cue exp gengotypes ./schemas/layer-4.cue
	@mv cue_types_gen.go layer4/generated_types.go
	@go build -o utils/types_tagger utils/types_tagger.go
	@utils/types_tagger layer1/generated_types.go
	@utils/types_tagger layer2/generated_types.go
	@utils/types_tagger layer3/generated_types.go
	@utils/types_tagger layer4/generated_types.go
	@rm utils/types_tagger

dirtycheck:
	@echo "  >  Checking for uncommitted changes ..."
	@if [ -n "$$(git status --porcelain)" ]; then \
		echo "  >  Uncommitted changes to generated files found!"; \
		echo "  >  Run make cuegen and commit the results."; \
		exit 1; \
	else \
		echo "  >  No uncommitted changes to generated files found."; \
	fi

lintinsights:
	@echo "  >  Linting security-insights.yml ..."
	@curl -O --silent https://raw.githubusercontent.com/ossf/security-insights-spec/refs/tags/v2.1.0/schema.cue
	@cue vet -d '#SecurityInsights' security-insights.yml schema.cue
	@rm schema.cue
	@echo "  >  Linting security-insights.yml complete."

PHONY: tidy test testcov lintcue cuegen dirtycheck lintinsights
