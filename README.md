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

Frontend is the interesting part that concurrently sends requests to the backend (see above) and groups the responses together. The frontend is currently hardcoded to submit four requests to the backend, where each request is given a sequential integer ID. Requests with an even request_id will sleep for two seconds.

```
$ go build frontend.go
$ ./frontend
Starting frontend on port 8080...

$ curl localhost:8080
request_id: 1, 2016-07-05 15:50:13.580381848 +0000 UTC
request_id: 3, 2016-07-05 15:50:13.5804684 +0000 UTC
request_id: 2, 2016-07-05 15:50:15.583446668 +0000 UTC
request_id: 0, 2016-07-05 15:50:15.583657375 +0000 UTC

==========================
Request served in: 2.00s
```

Make sure that the backend is running before sending requests to the frontend.

## Motivation

I originally started writing this code to learn Go from scratch. So, this is my first Go program that does something more than just "hello world". I might develop this project into a generic microservices broker if I ever become good enough at writing Go.
