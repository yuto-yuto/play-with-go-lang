ifneq ("$(wildcard .env)","")
	include .env
	export
	USER_NAME=username
	USER_ID=123
endif

.PHONY: show_env multi_lines lint format

show_env:
	@echo ADDED_IN_ENV: $${ADDED_IN_ENV}
	@echo USER_NAME: $${USER_NAME}
	@echo USER_ID: $${USER_ID}

define DEF1

multiline string can be written here.
The user name is $(USER_NAME)
The user ID is $(USER_ID)
endef

multi_lines:
	@echo $${DEF1}
	@echo "$${DEF1}"

lint:
	@golangci-lint run

format:
	@gofmt -w .