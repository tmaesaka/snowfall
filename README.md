# Snowfall

Snowfall is a reference Go implementation that attempts to illustrate how to concurrently send requests to various backend services and return the responses as one response. In theory, the request should only take as long as the slowest service.

## The Setup

This repository comes with two applications. A frontend and backend. Both applications have no external dependencies.

### Backend

An example microservice. It is just a minimal http server that responds with the current time in UTC.

```
$ go build backend.go
$ ./backend
Starting backend on port 8081...

$ curl localhost:8081
2016-07-05 15:50:13.580381848 +0000 UTC
```

### Frontend

Frontend is the interesting part that concurrently sends requests to the backend (see above) and groups the responses together. The frontend is currently hardcoded to spawn four worker [goroutines](https://golang.org/doc/effective_go.html#goroutines) on each request. For demo purpose, each worker is given a sequential ID and workers with an even ID number will sleep for two seconds.

```
$ go build frontend.go
$ ./frontend
Starting frontend on port 8080...

$ curl localhost:8080
worker_id: 3, 2016-07-05 17:17:54.697178619 +0000 UTC
worker_id: 1, 2016-07-05 17:17:54.697278168 +0000 UTC
worker_id: 2, 2016-07-05 17:17:56.701075292 +0000 UTC
worker_id: 0, 2016-07-05 17:17:56.701390656 +0000 UTC

==========================
Request served in: 2.00s
```

Make sure that the backend is running before sending requests to the frontend.

## Motivation

I originally started writing this code to learn Go from scratch. Snowfall is my first Go program that does something more than just "hello world".
