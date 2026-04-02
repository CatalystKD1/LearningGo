# Explination

If you have multiple go files in the same directory and they all have the same package you can run them together.

## How
First you need to create a Go module using this command

```bash
go mod init {name}
```

Then you can run the command

```bash
go run .
```

Which allows you to write to run the main file with all of the files in the module/package.