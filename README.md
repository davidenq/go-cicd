# Go CI/CD

[![Build Status](https://travis-ci.org/davidenq/go-cicd.svg?branch=master)](https://travis-ci.org/davidenq/go-cicd)
[![Coverage Status](https://coveralls.io/repos/github/davidenq/go-cicd/badge.svg?branch=master)](https://coveralls.io/github/davidenq/go-cicd?branch=master)

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

# To start running testing

For running unit testing:

`$ make unit-test`

# To start to develop

## Requirements
- `Docker`
- `docker-compose`
- `ruby for travis cli`
- `go > 1.12` if you want to develop without Docker

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