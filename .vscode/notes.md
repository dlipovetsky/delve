# How the `goroutine` command works

1. Client handles the command [here](../pkg/terminal/command.go#L905)
1. Client sends the command to the service [here](../service/rpc2/client.go#208)
1. Service  handles the command [here](../service/debugger/debugger.go#1264) 
1. Debugger switches the goroutine [here](../pkg/proc/target.go#315)

# Example of support for GDB "non-stop" mode in a Debug Adapter

https://github.com/eclipse-cdt-cloud/cdt-gdb-adapter/pull/194