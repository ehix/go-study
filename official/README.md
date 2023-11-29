# Offical Tutorial

[Site here](https://go.dev/doc/tutorial/getting-started)

## Important notes

### [`go mod edit -replace`](https://go.dev/doc/tutorial/call-module-code): 

For production use, you’d publish the `example.com/greetings` module from its repository (with a module path that reflected its published location), where Go tools could find it to download it. For now, because you haven't published the module yet, you need to adapt the example.com/hello module so it can find the `example.com/greetings` code on your local file system.

To do that, use the go mod edit command to edit the example.com/hello module to redirect Go tools from its module path (where the module isn't) to the local directory (where it is).

1. From the command prompt in the hello directory, run the following command:

```
$ go mod edit -replace example.com/greetings=../greetings
```

The command specifies that `example.com/greetings` should be replaced with `../greetings` for the purpose of locating the dependency. After you run the command, the go.mod file in the hello directory should include a replace directive:

```go
module example.com/hello

go 1.16

replace `example.com/greetings` => ../greetings
```

2. From the command prompt in the hello directory, run the go mod tidy command to synchronize the `example.com/hello` module's dependencies, adding those required by the code, but not yet tracked in the module.

```
$ go mod tidy
go: found example.com/greetings in example.com/greetings v0.0.0-00010101000000-000000000000
```

After the command completes, the example.com/hello module's go.mod file should look like this:

```go
module example.com/hello

go 1.16

replace example.com/greetings => ../greetings

require example.com/greetings v0.0.0-00010101000000-000000000000
```