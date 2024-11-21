# CLI helpers.

# help
help:
 @echo "Command line helpers for this project.\n"
 @just -l

# Run go linting
linting:
 goimports -w .
 go fmt .
 - go vet .
 - staticcheck .

# Run pre-commit
all-checks:
  pre-commit run --all-files

# Compile for all platforms
compile-all:
  ./distil.sh

# Setup linting
setup:
  go install golang.org/x/tools/cmd/goimports@latest
  go install honnef.co/go/tools/cmd/staticcheck@latest

# Fix imports
fix-imports:
  goimports -w .
