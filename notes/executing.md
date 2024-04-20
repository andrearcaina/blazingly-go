# Executing and Running a Go File

Since Go is a compiled language, the best practice is to compile the file into an executable and run it.

An example would be:

```bash
$ go build main.go
$ ./main
```

An alternative is to compile and run it.

```bash
$ go run main.go
```

An alternative is to compile without a main entry point.

```bash
$ go build
$ ./package_name # (default is the name of the directory)
```

An alternative is to compile and run without a main entry point.

```bash
$ go run
```

