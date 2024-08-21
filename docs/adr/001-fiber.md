# ADR-001 - Fiber as the web framework

## Context and Problem Statement

We need to choose a web framework to build the application. The web framework
should be fast, easy to use, and have a good community support.

## Solution

We will use Fiber v2 as the web framework to build the application. Fiber is a
web framework built on top of Fasthttp, the fastest HTTP engine for Go. Fiber is
fast, easy to use, and has a good community support. Also, Fiber has a lot of
features out of the box. It is a good choice for building web applications from
scratch.

## Other Solution Considered

- Fiber v3: Currently, Fiber v3 is in active development and it expected to have
  bugs and a lot of breaking changes.
- Go-chi: Go-chi is a lightweight, idiomatic and composable router for building
  Go HTTP services. It is a good alternative to Fiber, but it lacks some
  features that Fiber has.
- Gin gonic: another good alternative to Fiber, but it is slower than Fiber. It
  seems to be less popular and don't hace much support from the community.

Creation Date: 21/08/2024
Status: Accepted
