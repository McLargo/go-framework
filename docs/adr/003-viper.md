# ADR-002 - Viper as the configuration manager

## Context and Problem Statement

The application needs to load configuration from multiple sources (files, flags,
environment variables...) It needs to be flexible and easy to use.

## Solution

Viper is a complete configuration solution for Go applications including
12-Factor apps. It is designed to work inside an application, supporting many
formats and sources.

It supports:

- environment variables
- default values
- explicit configuration files
- key/value stores such as Consul

Viper is also able to watch for changes in the configuration file, and re-read
the configuration.

Viper is a popular choice for configuration management in Go applications.

## Other Solution Considered

godotenv is not flexible enough, as it does not support load of multiple files.

koanf is a good alternative, but it is not as popular as Viper.

Creation Date: 10/02/2025
Status: accepted
