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

This projects aims to show not only a process for CI/CD but also, to write code following clean code, TDD, branching strategy, versioning and so on.
On the other hand, the folder structure follows the definition specified [here](https://github.com/golang-standards/project-layout). But it's not a standardized folder structure. However, big projects such as Kubernetes, Prometheus, and others, follow that folder structure.


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
├── .dockerignore /
├── .gitignore /
├── .travis.yml /
├── .account.json.enc /
├── go.mod /
├── go.sum /
├── Makefile /
├── README.md /
│   ├── empty /
│   ├── nested /
│   │   ├── nested-a/
│   │   │   ├── a1/
│   │   │   │   ├── empty.js
│   │   │   ├── a2/
│   │   │   │   ├── a21/
│   │   │   │   │   ├── a211/
│   │   │   │   │   │   ├── empty.js
│   │   │   │   │   │   ├── example.js
│   │   ├── nested-b/
│   │   │   ├── b1/
│   │   │   │   ├── b11/
│   │   │   │   │   ├── empty.js
│   │   │   │   │   ├── empty.json
│   │   │   │   ├── empty.js
├── index.js
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