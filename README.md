# Go CI/CD

[![Build Status](https://travis-ci.org/davidenq/go-cicd.svg?branch=master)](https://travis-ci.org/davidenq/go-cicd)
[![Coverage Status](https://coveralls.io/repos/github/davidenq/go-cicd/badge.svg?branch=master)](https://coveralls.io/github/davidenq/go-cicd?branch=master)


# Checklist
- [x] unit testing with TDD
- [x] integration testing
- [x] end to end testing
- [x] pipeline ci/cd using Travis-ci service (automatic deployment in different stages specified at .travis.yml file)
- [x] deploy on local and cloud (google cloud run) (support docker-compose, google cloud run and google kubernetes engine)
- [x] branching strategy based on Github Flow (Master for production, other branchs for development/staging)
- [x] semantic versioning with git tags
- [x] Infraestructure as a code with terraform and bash scripting for createing cluster on Kubernetes
- [x] deploy docker containers on Google Kubernetes Engine through yaml files, kubectl and bash scripting
- [x] Handle scripts as a centralized way through Makefile
- [x] created a load balancer on Google Kubernetes Engine with 2 microservices
- [x] Published Docker containers on Google Registry
- [x] Code coverage using Coveralls service

This projects aims to show not only a process for CI/CD but also, to write code following clean code, TDD, branching strategy, versioning and so on.
On the other hand, the folder structure follows the definition specified [here](https://github.com/golang-standards/project-layout). But it's not a standardized folder structure. However, big projects such as Kubernetes, Prometheus, and others, follow that folder structure.

# List of Commands
The current implementation has a Makefile as a centralized handler of bash scripting. So for that reason, it's important to know what commands you can run

## app command
The app command are used to work without Docker
### Testing 
Running unit and integration tests
```bash
$make app action=test
```
Running unit test
```bash
$make app action=unit-test
```

Running unit integration-test
```bash
$make app action=integration-test
```

To perform end-to-end testing, you must first run the application
```bash
$make app action=run
```
```bash
$make app action=e2e-test
```
###  Running the app

```bash
$make app action=run
```

## docker command

docker command allows you to work based-on Docker.

### Running the app

In order to work with Docker, you must first build de Docker image.
```bash
$make docker action=build
```
```bash
$make docker action=app
```

## docker-compose command

docker-compose command just has two actions: up and down. The first one will build the Docker image and after that, will  run the application. The second one, will turn off the Docker container.
```bash
$make docker-compose action=up
```
```bash
$make docker-compose action=down
```

## travis command
Travis commands through make, are used for logging and for creating encrypt variables. Although those commands are not difficult to use, at first instance, those were added having in mind a ease way to use travis commands, but is a better option running Travis directly with its CLI. `make travis commands` will be deprecated for the next release.

## gcr command
gcr command is used to deploy the Docker container on Google Cloud Run. It command is called by the Travis-CI service. You can not do anything in local with this command because there are environmental variables gotten from travis environment variables.

## infra command

infra command is used to create a cluster on Google Kubernetes Engine. Terraform is used for providing this funcionality. This command is used only in local. For that reason, it's important that you can install gcloud CLI in your local machine, set the project of Google and provide credentials for authentication and authorization. This command doesn't provide you that funcionality (authentication and authorization).

For create a cluster:
```bash
$make infra
```
It's important to set all environment variables  on `./iaac/terraform/gke/cluster.tfvars`. You can get and example at `./iaac/terraform/gke/cluster.tfvars.example`.
For security, `cluster.tfvars` is ignoreg for git.

## gke command
gke command is used to create the workloads on the cluster previously created. Bash scripting, kubectl and yaml files are used to automate this tasks.

```bash
$make gke service=[FOLDER_SERVICE_NAME]
#FOLDER_SERVICE_NAME is the name of folder localed at ./deploy/gke, you have two options servicea or serviceb
#example
$ make gke service=servicea
$ make gke service=serviceb
```
To expose those service, you need to create a ingress. You can execut the below command for that:

```bash
kubectly apply -f ./deploy/gke/ingress/service.yaml

# Guidelines

1. Write unit testing (with TDD), integration testing and end to end testing
  a. For each unit test wrote, it must be committed to Github before writing the implementation.
  b. After that, must write the implementation and it will be tested in order to check whether it passed or not. When it has passed, the implementation must be committed to Github.
  c. Each implementation must be checked in order to improve the code following clean code and best development practices. After that, the new refactor must be committed to Github.
2. GitHubFlow as the strategy Branching
  - Github Flow could be used in the first instance and only for test purposes in order to understand in a better way this strategy branching.
3. Semver for versioning
  - It's important to generate a release only when you consider that the implementation could be completed to be tested.

# Folder structure
```
├── .vscode /
├── build /
│   ├── ci /
│   ├── package /
│   │   ├── docker/
│   │   │   ├── Dockerfile
├── cmd /
│   ├── domain /
│   ├── infra /
│   ├── interface /
│   ├── middleware /
│   ├── types /
│   ├── .env.example
│   ├── index.go
├── deploy /
│   ├── docker-compose /
│   ├── gcr /
│   ├── gke /
├── iaac /
│   ├── terraform /
│   │   ├── gke /
│   │   │   ├── cluster.tf
│   │   │   ├── cluster.tfvars.example
│   │   │   ├── variables.tf
├── scripts /
├── tests /
├── .dockerignore
├── .gitignore
├── .travis.yml
├── .account.json.enc
├── go.mod
├── go.sum
├── Makefile
├── README.md
```


# To start running testing

For running unit testing:

`$ make unit-test`

# To start to develop

## Requirements
- `Docker`
- `docker-compose`
- `ruby for travis cli`
- `go > 1.12` if you want to develop without Docker
- `brew install goreleaser`
- `terraform`

### install travis
```bash
gem install travis
```

#### issues installing travis
- if you get  'You don't have write permissions for the /Library/Ruby/Gems/2.6.0 directory' running the below commands:
  ```bash
  $ echo 'export PATH="/usr/local/opt/ruby/bin:$PATH"' >> ~/.bash_profile
  ```

  ```bash
  $ source ~/.bash_profile
  ```
[source](https://stackoverflow.com/questions/51126403/you-dont-have-write-permissions-for-the-library-ruby-gems-2-3-0-directory-ma?rq=1)

- if you get command not found, try:
  ```bash
  $ brew install travis
  ```


  ```bash
  travis encrypt-file credenciales.json --add
  ```