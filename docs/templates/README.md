# Project name

Small description of the project.

## User Experience

What is the output expected.

## Requirements

What are the requirements of the project.

## Description

Large description of the project, where user experience and requirements met
together.

## Project structure

Project structure is fixed, and it is based on the following structure:

- `cmd/`: folder to manage and gather files related to the commands
  - `bootstrap/`: folder to manage the bootstrap of the application
  - `config/`: folder to manage the configuration of the application
    - `main.go`: main file to start the application
- `docs/`: folder to manage and gather files related to the documentation
  - `adr`: folder to manage and gather files related to the architecture
    decision record
  - `assets`: folder to manage and gather assets related to the documentation
  - `flows`: folder to manage and gather files related to the flows
  - `templates`: folder to manage and gather templates
- `internal/`: folder to manage and gather files related to private code used by
  the application
  - `handlers`: folder to manage and gather files related to the handlers
- `pkg/`: folder to manage and gather files related to code that can be used by
  external applications
- `.air.toml`: file to manage air configuration
- `.dockerignore`: file to manage docker ignore
- `.env_template`: file to manage the environment variables
- `.gitignore`: file to manage git ignore
- `docker-compose.yml`: file to manage docker compose
- `Dockerfile`: file to manage docker
- `go.mod`: file to manage go modules
- `go.sum`: file to manage go modules
- `LICENSE`: file to manage the license
- `README.md`: main file to start the application
- `taskfile.yml`: file to manage tasks

As a reference, I follow the
[golang-standards/project-layout](https://github.com/golang-standards/project-layout/blob/master/README.md).

## Getting started

### Configuration

The application uses different configuration. The configuration is managed by
viper in two different ways:

- `.env`: to manage the environment variables. You can copy the `.env_template`
  file to `.env` in the root path and update the variables you want. This
  variables have precedence over the variables in the `cmd/config/config.yaml`
  file and it is used to manage the environment variables of the application.
- `cmd/config/config.yaml`: to manage the default configuration of the
  application. You can add/update the variables you want, but always push the
  changes to the repository.

### docker & docker compose utils

The project offers a dockerized environment to run the application. You will
need to have [docker](https://docs.docker.com/get-docker/) and
[docker compose](https://docs.docker.com/compose/install/) installed. Please
install in your machine using the instructions provided in the links.

If you are new to docker compose, and love to use the terminal like me, please
see below some of the most used commands:

- build images: `docker compose build`
- start containers: `docker compose up -d <name>`
- remove containers: `docker compose down`
- list all containers: `docker compose ps`
- stop containers: `docker compose stop <name>`
- enter inside the container: `docker exec -it <container_id> bash`
- to see the logs in real time: `docker logs -f <container_id>`
- list all images: `docker images`
- remove an image: `docker rmi <image_id>`

### Installation and usage

- step 1: `docker compose build`
- step 2: `docker compose up -d`
- step 3: go to your browser and open
  [http://localhost:3000/healthz](http://localhost:3000/healthz). If not opening,
  `docker logs -f <container_id>` to see the logs for errors.
- step 4: start coding

To simplify the usage and development, you can use the following commands inside
[taskfile](https://taskfile.dev/installation/):

- `task build-docker`: run docker
- `task build-docker-compose`: run docker compose

### Native installation

If you prefer to run the application natively, you can do it by running the
following commands:

- step 1: `go run cmd/main.go`
- step 2: go to your browser and open
  [http://localhost:3000/healthz](http://localhost:3000/healthz). If not opening,
  check the logs for errors.
- step 3: start coding

To simplify the usage and development, you can use the following commands inside
[taskfile](https://taskfile.dev/installation/):

- `task dev`: run the application in development mode
- `task build`: run the application in production mode

## Architecture Decision Record

A decision record is a document that captures an important architectural
decision made along with its context and consequences.

Below is a list of the ADRs for this project:

- [ADR-000](../../docs/templates/adr.md) - ADR template

## Flows

Flows of the project. Always keeping documentation close to code.

## Contribution

Use [conventional commit](https://www.conventionalcommits.org/en/v1.0.0/) to
your commit messages.

Trunk-based development vs Gitflow.

## Testing

I do believe that unit test are the most important part of a project, to keep
the code clean, and avoid adding breaking code to a piece of software that is
working. A good coverage of the application indicates that services are robust.
[TDD Test-Driven Development](
https://www.guru99.com/test-driven-development.html) is my ideal development
methodology.

## Security

If applicable.

## CI/CD

What is the CI/CD pipeline.

## Monitoring

What is the monitoring strategy and tools. Liveness endpoint, tracing and log
systems...

## Troubleshooting

Most common errors you can find when running the project.

## Performance

Usage of tools to measure the performance of the application.

## License

If applicable.
