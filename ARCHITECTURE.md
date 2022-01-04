# Project overview
```
.
├── cmd
│   └── main.go
├── .env
├── go.mod
├── go.sum
├── pkg
│   ├── config
│   │   └── config.go
│   ├── handlers
│   │   ├── handlers.go
│   │   └── index.go
│   ├── models
│   │   ├── base.go
│   │   └── user.go
│   ├── server
│   │   ├── routes.go
│   │   ├── server.go
│   │   └── server_test.go
│   └── storage
│       ├── sql
│       │   ├── sql.go
│       │   ├── user.go
│       │   └── user_test.go
│       └── storage.go
└── README.md
```

## cmd
This is the main applicaction starting point, the one in charge of instantiating a new db connection, setting the http server, its routes, etc... Also implements a basic graceful shutdown system.

## .env
Environment variables can be loaded using a `.env` file, a simple `.env.example` is provided.

## pkg
This is where the main logic of the application resides.

### pkg.config
The `config` package loads variables from the system environment and hooks them up in the application.

### pkg.handlers
Handlers are directly connected to routes (can be viewed as "controllers") and handle the route's action.
Handlers are usually tied to the model they refer to.

### pkg.models
The application models, usually used in conjuction to an ORM to provide sensible mappings.

### pkg.server
Abstracts the `net/http` package and provides a way to easily hook up middlewares to routes.

### pkg.storage
The `Storage` interface provides an easy way to implement different kinds of databases to the app.
