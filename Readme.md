# A mini container

This is a minimal container simulation that helps to understand the internal how containers isolate themselves from other process. This is just a learning implementation inspired from [Container from scratch - Lizrice](https://youtu.be/8fi7uSYlOdc?si=b0ZaaGhip8PxgZNa)

## Documentation

For a simple implementation of the container we use first a command args i.e run which calls our run function and child for calling child function

### Run

This function is used to create a PID namespace and other namespaces.In here the the child function itself is called using another call.

The reason behind two different calls is by default the command executes in line `cmd.Run()` so the things that we do before that is in the parent process.
Hence if we mount the `/proc` directory for process file directory we will map it to the current namespace process which is parent process.

So to truly isolate we must have th namespace as which we created i.e run the command and inside it create a namespace and mount the filesystem. With this the kernel will take the PID namespace that we created as refernce and only list the process that runs inside. Such that acheiving isolation of process.
