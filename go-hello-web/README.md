# Example go webserver

This folder contains a skeleton example code for a simple Go webserver.  
It has multiple endpoints, uses a (fake) external service, dependency injection, unit tests etc.

The server itself provides two endpoints:

- `/ping` returns "Pong!"
- `/roll` rolls a die and returns a random number between 1-6

## Code layout

Each directory represents a module/package of the program.  
The main entrypoint for this sample is `cmd/webserver/main.go`.
The `main()` function here is responsible for instantiating and configuring the web server object and its dependencies.
It also start the webserver listening and will run until you manually terminate the program.

The `/dice` package is an example package acting as an 'external' dependency to show how we use it in our webserver.
In a real program our web service might depend on similar packages, such as the [AWS Go SDK](https://github.com/aws/aws-sdk-go), that enable interaction with an external service.

The `/internal/webservice` package has the core logic of our program which defines how to respond to HTTP requests.  
The web server utilises our external dice service through an abstracted interface. The interface only declares functions the server will need and allows us to easily mock out our external dependency during testing.  
The tests for this package verify the expected handling of failure cases, and a variety of success cases.

_Note: `internal` is a special directory name for controlling package imports in Go, but it's not important to this particular project. You can read more about it [here](https://golang.org/doc/go1.4#internalpackages) if you like._

## Run the example

### The web server

First, to build the webserver execute (while in the `go-hello-web` folder):

```sh
go build ./cmd/webserver
```

Run the server:

```sh
./webserver
```

Then, if you open up a separate terminal, you can call the endpoints:

```sh
$ curl localhost:8080/ping
Pong!
$ curl localhost:8080/roll
6
$ curl localhost:8080/roll
6
$ curl localhost:8080/roll
2
```

### The tests

Running the tests is simple:

```sh
go test ./...
```

This means we will run all tests that we can find in all subdirectories below the current folder.
