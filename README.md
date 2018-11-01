# Golang Goban

A Goban webapp written in Golang

## Getting Started

This is a Docker app all you need to do is install the prereqs below then start the app as instructed

### Prerequisites

* [Docker](https://docs.docker.com/install/)
* [Docker Compose](https://docs.docker.com/compose/install/)

### Starting

Right now the app starts before the db is healthy, I need to add a health check to the docker-compose. In the meantime do:

```
docker-compose up -d db
```

```
// checking for "server is ready for connections"
docker-compose logs -f db
```

```
docker-compose up app
```