# ADR-005 - Performance

## Context and Problem Statement

We need to ensure that the system is able to handle the expected load and
provide a good user experience.

## Solution

[Vegeta](https://github.com/tsenart/vegeta) is a versatile HTTP load testing
tool built out of a need to drill HTTP services with a constant request rate.

It can be used both as a library and as a command line utility, making it super
easy to use. Plus, it has a rich set of features, such as:

- Fast and scalable
- Supports HTTP/1.1, HTTP/2, and WebSockets
- Supports distributed load testing
- Supports custom headers and body
- Supports rate limiting
- Supports multiple output formats

## Other Solution Considered

[locust](https://github.com/locustio/locust) is another load testing tool, but
it is developed in Python.

Creation Date: 25/03/2025
Status: accepted
