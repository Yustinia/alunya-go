ENTRY := ./cmd/

NAME := alunya
BUILDCON := $(shell fd -u -H -t file -e go)

$(NAME): $(BUILDCON)
	cd $(ENTRY) && go build -x .

.PHONY: go

go:
	go run $(ENTRY)
