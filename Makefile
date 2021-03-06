export COLOR_NC='\033[0m' # No Color
export COLOR_WHITE='\033[1;37m'
export COLOR_BLACK='\033[0;30m'
export COLOR_BLUE='\033[0;34m'
export COLOR_LIGHT_BLUE='\033[1;34m'
export COLOR_GREEN='\033[0;32m'
export COLOR_LIGHT_GREEN='\033[1;32m'
export COLOR_CYAN='\033[0;36m'
export COLOR_LIGHT_CYAN='\033[1;36m'
export COLOR_RED='\033[91m'
export COLOR_LIGHT_REenv='\033[1;31m'
export COLOR_PURPLE='\033[0;35m'
export COLOR_LIGHT_PURPLE='\033[1;35m'
export COLOR_BROWN='\033[0;33m'
export COLOR_YELLOW='\033[1;33m'
export COLOR_GRAY='\033[0;30m'
export COLOR_LIGHT_GRAY='\033[0;37m'
export PORT=8080
export ENVIRONMENT=develop

MAKEFLAGS += --no-print-directory

# Commands for app.
app:
ifndef action
	@echo ${COLOR_RED}"Error: You must define the action (a) as follows: make app action=[action]"
	@echo ${COLOR_BLUE}"[action]: run, test, unit-test, integration-test, e2e-test"
	@echo ${COLOR_GREEN}"	     run: start the gocicd server without Docker" ${COLOR_YELLOW}"(If you want to use docker-compose to start the app, run: make docker-compose action=up)"
	@echo ${COLOR_GREEN}"	     test: run all the kind of tests"
	@echo ${COLOR_GREEN}"	     unit-test: run only the unit tests"
	@echo ${COLOR_GREEN}"	     integration-test: run only the integration tests"
	@echo ${COLOR_GREEN}"	     e2e-test: run only the end to end tests"${COLOR_NC}
else
	./scripts/app/app.sh
endif

# Commands for Docker
docker:
ifndef action
	@echo ${COLOR_RED}"Error: You must define the action (a) as follows: make docker action=[action]"
	@echo ${COLOR_BLUE}"[action]: build, run"
	@echo ${COLOR_GREEN}"	     build: build a Docker container"
	@echo ${COLOR_GREEN}"	     run: start the gocicd server on Docker" ${COLOR_YELLOW}"(If you want to use docker-compose to start the app, run: make docker-compose action=up)"${COLOR_NC}
else
	./scripts/docker/docker.sh
endif

# Commands for docker-compose
docker-compose:
ifndef action
	@echo ${COLOR_RED}"Error: You must define the action (a) as follows: make docker-compose action=[action]"
	@echo ${COLOR_BLUE}"[action]: up, down"
	@echo ${COLOR_GREEN}"	     up: start the app using docker-compose"
	@echo ${COLOR_GREEN}"	     down: stop the app using docker-compose"${COLOR_NC}
else 
	ACTION={a} ./scripts/docker-compose/docker-compose.sh
endif

gcr:
	./deploy/gcr/deploy.sh

infra:
	@terraform init ./iaac/terraform/gke
	@terraform validate dir ./iaac/terraform/gke
	@terraform apply -var-file=./iaac/terraform/gke/cluster.tfvars ./iaac/terraform/gke

gke:
ifndef service
	@echo ${COLOR_RED}"Error: You must define the service to deploy (a) as follows: make gke service=[SERVICE_FOLDER_NAME]"
	@echo "For instance: make gke service=servicea"
	@echo ${COLOR_NC}""
else 
	SERVICE=${service} ./scripts/gke/deploy.sh
endif