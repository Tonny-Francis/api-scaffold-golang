# API Scaffold Golang

<center>
  <img src="https://camo.githubusercontent.com/0d7633d0e4571f4b9800a6aa104d03386a86df5c33b4650d7dce88a6de1c8904/68747470733a2f2f7261772e6769746875622e636f6d2f676f6c616e672d73616d706c65732f676f706865722d766563746f722f6d61737465722f676f706865722e706e67" width="150">
</center>

## Description

The project is an API developed in Golang, using the Gin-Gonic framework to facilitate the creation of routes and handling HTTP requests. The project structure is organized and modular, following the best practices for Golang development.

## Technologies Used

- **Gin-Gonic:** Efficient web framework for building Golang web services.
- **Docker:** Containers to facilitate deployment and execution in different environments.
- **Makefile:** Task automation for project compilation, execution, and maintenance.
- **Devcontainer:** Support for a consistent development environment.
- **Air:** Facilitates development with automatic reloading during the development process.
- **Testes Unitários:** Implemented to ensure the robustness and correct execution of routes.

## Project Structure

```plaintext
.
├── Dockerfile
├── Dockerfile.dev
├── Makefile
├── README.md
├── bin
│   └── api
├── cmd
│   └── api
│       ├── main.go
│       └── main_test.go
├── env
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   ├── router.go
│   │   └── v1
│   │       └── routes.go
│   ├── config
│   └── container
│       └── container.go
├── pkg
│   ├── core
│   │   ├── env
│   │   │   └── env.go
│   │   ├── logger
│   │   │   └── logger.go
│   │   ├── router
│   │   │   └── mode.go
│   │   └── server
│   │       └── server.go
│   ├── domain
│   │   └── example
│   │       ├── handler.go
│   │       ├── schemas.go
│   │       └── service.go
│   └── services
│       ├── http
│       │   ├── error.go
│       │   ├── requestHandler.go
│       │   ├── response.go
│       │   └── service.go
│       └── validator
│           ├── schemaValidator.go
│           └── service.go
└── tmp
    └── build-errors.log
```

## How to Use

### Prerequisites

- Make sure you have Docker installed.
- Install Air for automatic reloading during development.

### Environment Setup

1. Clone the repository.
2. Run `make build` to compile the project.

### Run Locally

- Run `make run` to start the API locally.

### Run Using Devcontainer in VSCode

1. Open the project in VSCode.
2. Install the "Remote - Containers" extension.
3. Open the command palette (Ctrl + Shift + P) and select "Remote-Containers: Reopen in Container".
4. VSCode will generate and open the Devcontainer automatically.

### Development

- Use the Devcontainer for a consistent development environment.
- Run `make dev` to start the project with Air for automatic reloading.

### Tests

- Run `make test` to execute unit tests.

## Contribution

Contributions are welcome! Feel free to open issues or send pull requests.

## References

This repository has been structured with insights from the Golang community. Take a look at the following projects that inspired and provided guidelines for the organization of this repository:

- [facily-tech/go-scaffold](https://github.com/facily-tech/go-scaffold)
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)

## Image Reference

The Gopher image used in this project was obtained from the [golang-samples/gopher-vector](https://github.com/golang-samples/gopher-vector) repository. All credits for creating this image go to the maintainers of the original project.