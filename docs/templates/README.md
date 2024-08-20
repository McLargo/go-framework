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

TODO

## Getting started

The project offers a dockerized environment to run the application. You will
need to have [docker](https://docs.docker.com/get-docker/) and
[docker compose](https://docs.docker.com/compose/install/) installed. Please
install in your machine using the instructions provided in the links.

### docker & docker compose utils

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

### poetry

If you are new to poetry, see below some of the most used commands:

- `poetry --help`: to see all the commands available
- `poetry init`: to create a new project
- `poetry install`: to install all the dependencies
- `poetry add <package_name>`: to add a new package to the project
- `poetry remove <package_name>`: to remove a package from the project
- `poetry run <command>`: to run a command inside the virtual environment
- `poetry shell`: to enter inside the virtual environment

### Installation and usage

- step 1: `docker compose build`
- step 2: `docker compose up -d`
- step 3: go to your browser and open
  [http://localhost:8000/healthz](http://localhost:8000/healthz). If not opening,
  `docker logs -f <container_id>` to see the logs for errors.
- step 4: start coding

### Native installation

If you prefer to run the application natively, you can do it by running the
following commands:

TODO: add instructions to run the application natively + requirements

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

TODO

## License

If applicable.
