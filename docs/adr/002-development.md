# ADR-002 - Development environment

## Context and Problem Statement

Go is a language that is compile. To run the application, we need to have to
build the application and run the binary. This process can be cumbersome while
you are developing the application, as it forces you to shutdown and start the
application in order to see your latest changes in the code.

## Solution

Implement a development environment that allows the developer to live reload the
application while you are developing.

[Air](https://github.com/air-verse/air) is a tool that allows you to do that. It watches for changes in the code and
rebuilds the application automatically. It support some configuration, which is
nice to have in different scenarios/applications.

## Other Solution Considered

- Fresh is another tool that allows you to do the same thing, but it has no
releases since 2022.

Creation Date: 23/08/2024
Status: Accepted
