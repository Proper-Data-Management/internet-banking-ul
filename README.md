#### added new blank module
```sh
# bash add_blank_module.sh :name_module
bash add_blank_module.sh customer
```

#### help
```sh
➜  al_hilal_core git:(main) make help

Usage:
  make <target>

Targets:
  Build:
    run                 Build and run your project
    build               Build your project and put the output binary in out/bin/
    clean               Remove build related file
    vendor              Copy of all packages needed to support builds and tests in the vendor directory
    watch               Run the code with cosmtrek/air to have automatic reload on changes
  Test:
    test                Run the tests of the project
    coverage            Run the tests of the project and export the coverage
  Lint:
    lint                Run all available linters
    lint-dockerfile     Lint your Dockerfile
    lint-go             Use golintci-lint on your project
    lint-yaml           Use yamllint on the yaml file of your projects
  Docker:
    docker-build        Use the dockerfile to build the container
    docker-release      Release the container with tag latest and version
  Dependencies:
    global-deps         Install Oracle and Vasco dependencies
    docker-deps         Install Oracle and Vasco dependencies for the docker container withouth the need of sudo
  Help:
    help                Show this help.

```

#### Install Oracle dependencies
```sh
make global-deps
```

#### Build project
```sh
make build
```

#### Run project
```sh
make run
```