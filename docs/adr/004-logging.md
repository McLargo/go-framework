# ADR-004 - zap as the logging library

## Context and Problem Statement

The application needs a logging library to log the application's events and
errors. The logging library should be able to log in different levels and be
easy to use.

## Solution

The solution is to use [zap](https://pkg.go.dev/go.uber.org/zap) as the logging
library. Zap is a fast, structured, and leveled logging library for Go. It is
easy to use and has a good performance. It provides Development and Production
modes by default, and it is highly configurable.

Zap provides good documentation, and it is widely used in the Go community. It
provides a structured logging format, which is useful for debugging.

## Other Solution Considered

- logrus: logrus is a structured logger for Go, but it is slower than zap.
- log: the standard library log package is not structured and has limited
  features.

Creation Date: 11/02/2025
Status: accepted
