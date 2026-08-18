ENTRY := ./cmd/
NAME := alunya
BUILDCON := $(shell fd -u -H -t file -e go)
DOCKERMP := /app

$(NAME): $(BUILDCON)
	go build -o $(NAME) -x $(ENTRY)

.PHONY: go up

go:
	go run $(ENTRY)

up:
	docker compose run -it --rm -v $$(pwd):$(DOCKERMP) -w $(DOCKERMP) --name $(NAME) $(NAME) bash
	docker compose down
