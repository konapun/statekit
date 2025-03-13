# statekit

Statekit is a minimal state management library for Go providing proctored access to state items and observability.

The API is partially a concession to the barely-adequate nature of Go generics, but should mostly work without needing to do any runtime type assertions in the client code. If Go ever gets generic type defs on methods, I'll improve the API surface.

I'm currently developing this for use as part of a game I'm working on and will provide more complete documentation as it proves itself in practice. For now, check the example in `examples/`.

## Installation
```sh
go get github.com/konapun/statekit
```
