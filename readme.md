# Go Redis Rate Limiter

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/viper-18/go-redis-ratelimitter)
![GitHub](https://img.shields.io/github/license/viper-18/go-redis-ratelimitter)

A simple rate limiter implementation in Go using Redis as the backend storage.

## Introduction

Go Redis Rate Limiter is a project that demonstrates how to implement rate limiting in a Go application using Redis. It provides a middleware that can be integrated with web frameworks like Fiber to limit the number of requests a client can make within a specified time period.

## Features

- **Redis Integration**: Utilizes Redis as a backend for storing request timestamps.
- **Flexible Configuration**: Adjustable rate limits and time windows.
- **Middleware Integration**: Compatible with Go web frameworks like Fiber.

## Installation

To use this project, ensure you have Go installed. Then, install the project and its dependencies:

```bash
go get github.com/viper-18/go-redis-ratelimitter
