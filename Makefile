PROJECTNAME := $(shell basename "$(PWD)")
OS := $(shell uname -s | awk '{print tolower($$0)}')
GOARCH := amd64

## run: execete main application in local machine
.PHONY: run
run:
	go run cmd/main.go

## run_local: execete main application using local env file
## windows user note: "set APPENV=local && cmd /c go run cmd/main.go"
.PHONY: run_local
run_local:
	export APPENV=local; go run cmd/main.go

## run_local2: execete main application using local env file
.PHONY: run_local2
run_local2:
	export APPENV=local2; go run cmd/main.go

## run_export_env: test execete main application without local env file but environment exported
.PHONY: run_export_env
run_export_env:
	export SECRET_DB_USERNAME=exportenv_user SECRET_DB_PASSWORD=exportenv_password; go run cmd/main.go

## test: run unit test
.PHONY: test
test:
	go test ./..

## tidy: special go mod tidy without golang database checksum(GOSUMDB) 
.PHONY: tidy
tidy:
	export GOSUMDB=off; go mod tidy

## updatelib: undate dependency libs
.PHONY: updatelibs
updatelibs:
	export GOSUMDB=off ;go get gitdev.devops.company.com/my_project/my-repo-something

## up: docker compose up
.PHONY: up
up:
	docker-compose up -d

## down: docker compose dowm
.PHONY: down
down:
	docker-compose down -v --remove-orphans

## help: helper
.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Project: ["$(PROJECTNAME)"]"
	@echo " Please choose a command"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
