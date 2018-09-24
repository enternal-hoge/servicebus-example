# servicebus-example
Azure ServiceBus example golang

I try to execute Microsoft Azure Service Bus Client for Golang Quick Start, but does not working.

so I created a example that reads the test code and works correctly.

# Requiment
+ go 1.11
+ [Microsoft Azure Service Bus Client for Golang](https://github.com/Azure/azure-service-bus-go)

# Note
+ This application use go 1.11.
+ don't use glide or dep, use go modules.

set environment variables.
```
GO111MODULE=on
```

on first build, down load dependency modules automatically.
```
$ cd servicebus-example
$ go build
```


# Running
```
$ cd servicebus-example
$ go run main.go
```

# Lisence
MIT

# Author
Kenichi Murakata
