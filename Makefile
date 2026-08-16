ENTRY := ./cmd/
NAME := alunya
BUILDCON := $(shell fd -u -H -t file -e go)

$(NAME): $(BUILDCON)
	go build -o $(NAME) -x $(ENTRY)

.PHONY: go

go:
	go run $(ENTRY)
